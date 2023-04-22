package challenge

import "competition-backend/app/http/models"

type Challenge struct {
	models.BaseModel

	Title  string
	Desc   string
	Answer string
	Score  int64

	models.CommonTimestampsField
}
