package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ManagerLoginRequest struct {
	Username string `json:"username,omitempty" valid:"username"`
	Password string `json:"password,omitempty" valid:"password"`
}

func ManagerLogin(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"username": []string{"required", "min:5"},
		"password": []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"username": []string{
			"required:用户名为必填项",
			"min:用户名长度需大于 5",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
	}
	errs := validate(data, rules, messages)
	return errs
}
