package question

import (
	"competition-backend/app/http/models"
	"competition-backend/pkg/database"
)

type Question struct {
	models.BaseModel

	Title string `json:"title,omitempty"`
	Desc  string `json:"desc,omitempty"`
	Score int64  `json:"score,omitempty"`

	Answer string `json:"-"`
}

func (q *Question) Create() {
	database.DB.Table("questions").Create(&q)
}

func (q *Question) Delete() (rowsAffected int64) {
	res := database.DB.Delete(&q)
	return res.RowsAffected
}

func (q *Question) Save() (rowsAffected int64) {
	res := database.DB.Table("questions").Save(&q)
	return res.RowsAffected
}

func Get(id string) (q Question) {
	database.DB.Table("questions").Where("id", id).First(&q)
	return
}
