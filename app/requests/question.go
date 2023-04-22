package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type QuestionRequest struct {
	Title  string `json:"title,omitempty" valid:"title"`
	Desc   string `json:"desc,omitempty" valid:"desc"`
	Answer string `json:"answer,omitempty" valid:"answer"`
	Score  int64  `json:"score,omitempty" valid:"score"`
}

func CreateQuestion(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"title":  []string{"required"},
		"answer": []string{"required"},
		"desc":   []string{"required"},
		"score":  []string{"required", "min:0"},
	}
	messages := govalidator.MapData{
		"title":  []string{"required:题目名称为必填项"},
		"desc":   []string{"required:题目描述为必填项"},
		"answer": []string{"required:题目答案为必填项"},
		"score":  []string{"required:题目分值为必填项", "min:分值需大于 0"},
	}

	return validate(data, rules, messages)
}
