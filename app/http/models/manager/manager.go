package manager

import (
	"competition-backend/app/http/models"
	"competition-backend/pkg/database"
	"competition-backend/pkg/hash"
	"gorm.io/gorm"
)

type Manager struct {
	models.BaseModel
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	models.CommonTimestampsField
}

func Get(id string) (data Manager) {
	database.DB.Table("manager").Where("id", id).First(&data)
	return
}

func (m *Manager) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, m.Password)
}

func (m *Manager) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.IsHashed(m.Password) {
		m.Password = hash.BcryptHash(m.Password)
	}
	return
}

func GetByUserName(username string) (m Manager) {
	database.DB.Table("manager").Where("username = ?", username).First(&m)
	return
}
