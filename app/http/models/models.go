package models

import (
	"competition-backend/pkg/database"
	"github.com/spf13/cast"
	"time"
)

type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `database:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt time.Time `database:"column:updated_at;index;" json:"updated_at,omitempty"`
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return cast.ToString(a.ID)
}

// GetUserByID 通过 ID 查找用户，适用 Admin 和 User
func GetUserByID[T interface{}](id string, table string) (user T) {
	database.DB.Table(table).Where("id", id).First(&user)
	return
}
