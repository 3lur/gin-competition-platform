package game

import "competition-backend/app/http/models"

type Game struct {
	models.BaseModel

	Title  string `json:"title,omitempty"`
	Desc   string `json:"desc,omitempty"`
	Status uint8  `json:"status,omitempty"`

	models.CommonTimestampsField
}
