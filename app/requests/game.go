package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type GameRequest struct {
	Title string `json:"title,omitempty" valid:"title"`
	Desc  string `json:"desc,omitempty" valid:"desc"`
}

func CreateGame(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"title": []string{"required"},
		"desc":  []string{"required"},
	}
	messages := govalidator.MapData{
		"title": []string{"required:比赛名称为必填项"},
		"desc":  []string{"required:比赛描述为必填项"},
	}
	return validate(data, rules, messages)
}
