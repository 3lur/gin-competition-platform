package logger

import (
	"competition-backend/pkg/helpers"
	"context"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type GormLogger struct {
	ZapLogger     *zap.Logger
	SlowThreshold time.Duration
}

// NewGormLogger 实例化一个 GormLogger 对象
func NewGormLogger() GormLogger {
	return GormLogger{
		ZapLogger:     Logger,                 // 使用全局的 logger.ZapLogger 对象
		SlowThreshold: 200 * time.Millisecond, // 慢查询阈值，单位为千分之一秒
	}
}

func (l GormLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	return GormLogger{
		ZapLogger:     l.ZapLogger,
		SlowThreshold: l.SlowThreshold,
	}
}

func (l GormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	l.logger().Sugar().Debugf(s, i...)
}

func (l GormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	l.logger().Sugar().Warnf(s, i...)
}

func (l GormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	l.logger().Sugar().Errorf(s, i...)
}

func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	// 获取运行时间
	elapsed := time.Since(begin)
	// 获取 SQL 请求和返回条数
	sql, rows := fc()

	// 通用字段
	logFields := []zap.Field{
		zap.String("sql", sql),
		zap.String("time", helpers.MicrosecondsStr(elapsed)),
		zap.Int64("rows", rows),
	}

	// Gorm 错误
	if err != nil {
		// 记录未找到的错误使用 warning 等级
		if errors.Is(err, gorm.ErrRecordNotFound) {
			l.logger().Warn("Database ErrRecordNotFound", logFields...)
		} else {
			// 其他错误使用 error 等级
			logFields = append(logFields, zap.Error(err))
			l.logger().Error("Database Error", logFields...)
		}
	}

	// 慢查询日志
	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		l.logger().Warn("Database Slow Log", logFields...)
	}

	// 记录所有 SQL 请求
	l.logger().Debug("Database Query", logFields...)
}

func (l GormLogger) logger() *zap.Logger {
	// 跳过 database 内置的调用
	var (
		gormPkg    = filepath.Join("database.io", "database")
		zapGormPkg = filepath.Join("moul.io", "zapgorm2")
	)

	// 减去一次封装，以及一次在 logger 初始化里添加 zap.AddCallerSkip(1)
	clone := l.ZapLogger.WithOptions(zap.AddCallerSkip(-2))

	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		case strings.Contains(file, gormPkg):
		case strings.Contains(file, zapGormPkg):
		default:
			// 返回一个附带跳过行号的新的 zap logger
			return clone.WithOptions(zap.AddCallerSkip(i))
		}
	}
	return l.ZapLogger
}
