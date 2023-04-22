package middlewares

import (
	"bytes"
	"competition-backend/pkg/helpers"
	"competition-backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"io"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := &responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = w

		// 获取请求数据
		var requestBody []byte
		if c.Request.Body != nil {
			// buffer 对象，只能读取一次
			requestBody, _ = io.ReadAll(c.Request.Body)
			// 读取后重新赋值 c.Request.Body，以供后续的其他操作
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 设置开始时间
		start := time.Now()

		c.Next()

		// 记录日志的逻辑
		cost := time.Since(start)
		responseStatus := c.Writer.Status()

		logFields := []zap.Field{
			zap.Int("status", responseStatus),
			zap.String("request", c.Request.Method+" "+c.Request.URL.String()),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", helpers.MicrosecondsStr(cost)),
		}

		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
			// 请求的内容
			logFields = append(logFields, zap.String("Request Body", string(requestBody)))

			// 响应的内容
			logFields = append(logFields, zap.String("Response Body", w.body.String()))
		}

		if responseStatus > 400 && responseStatus <= 499 {
			logger.Warn("HTTP Warning "+cast.ToString(responseStatus), logFields...)
		} else if responseStatus >= 500 && responseStatus <= 599 {
			logger.Error("HTTP Error "+cast.ToString(responseStatus), logFields...)
		} else {
			logger.Debug("HTTP Access Log", logFields...)
		}
	}
}
