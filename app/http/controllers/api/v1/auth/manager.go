package auth

import (
	v1 "competition-backend/app/http/controllers/api/v1"
	"competition-backend/app/http/models/manager"
	"competition-backend/app/requests"
	"competition-backend/pkg/jwt"
	"competition-backend/pkg/redis"
	"competition-backend/pkg/response"
	"errors"
	"github.com/gin-gonic/gin"
)

type ManagerAuthController struct {
	v1.BaseAPIController
}

func (mc *ManagerAuthController) Login(c *gin.Context) {
	request := requests.ManagerLoginRequest{}
	if ok := requests.Validate(c, &request, requests.ManagerLogin); !ok {
		return
	}
	m, err := login(request.Username, request.Password)
	if err != nil {
		response.Unauthorized(c, "账号或密码错误")
	} else {
		token := jwt.NewJWT().IssueToken(m.GetStringID(), m.Username)
		result := redis.Redis.Set("manager-jwt", token)
		if !result {
			return
		}
		response.Success(c)
	}
}

func (mc *ManagerAuthController) Logout(c *gin.Context) {
	result := redis.Redis.Del("manager-jwt")
	if !result {
		response.Abort500(c, "操作失败，请稍后重试")
		return
	}
	response.Success(c)
}

func login(username, password string) (manager.Manager, error) {
	m := manager.GetByUserName(username)
	if m.ID == 0 {
		return manager.Manager{}, errors.New("账号不存在")
	}
	if !m.ComparePassword(password) {
		return manager.Manager{}, errors.New("密码错误")
	}
	return m, nil
}
