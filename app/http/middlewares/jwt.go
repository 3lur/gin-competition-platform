package middlewares

import (
	"competition-backend/app/http/models/manager"
	"competition-backend/pkg/config"
	"competition-backend/pkg/jwt"
	"competition-backend/pkg/redis"
	"competition-backend/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ManagerJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 从 redis 中获取 token
		hasKey := redis.Redis.Has("manager-jwt")
		token := redis.Redis.Get("manager-jwt")
		// 获取不到 token 直接返回错误
		if !(hasKey && token != "") {
			response.Unauthorized(c, fmt.Sprintf("身份认证失败，请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}

		// 2. 将 token 添加到请求头，注意格式是：Bearer ...
		c.Request.Header.Set("Authorization", "Bearer "+token)
		// 解析出错
		claims, err := jwt.NewJWT().ParseToken(c)
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("身份认证失败，请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}
		data := manager.Get(claims.UserID)
		if data.ID == 0 {
			response.Unauthorized(c, "用户不存在，可能已被删除")
			return
		}

		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		c.Set("current_user_id", data.ID)
		c.Set("current_user_name", data.Username)
		c.Set("current_user", data)

		c.Next()
	}
}

//func JWT() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		claims, err := jwt.NewJWT().ParseToken(c)
//		// JWT 解析失败
//		if err != nil {
//			response.Unauthorized(c, fmt.Sprintf("身份认证失败，请查看 %v 相关的接口认证文档", config.GetString("app.name")))
//			return
//		}
//
//		// 通过 claims 中的 UserID 查找出对应的用户
//		// 如果用户已经删除，即使 Token 是正确的，也无法进行请求
//		user := models.GetUserByID[admin.Admin](claims.UserID, "admin")
//		if user.ID == 0 {
//			response.Unauthorized(c, "找不到对应用户，用户可能已删除")
//			return
//		}
//
//		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
//		c.Set("current_user_id", user.GetStringID())
//		c.Set("current_user_name", user.Username)
//		c.Set("current_user", user)
//
//		c.Next()
//	}
//}
