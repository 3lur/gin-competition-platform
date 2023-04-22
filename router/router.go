package router

import (
	v1API "competition-backend/app/http/controllers/api/v1"
	"competition-backend/app/http/controllers/api/v1/auth"
	"competition-backend/app/http/middlewares"
	"competition-backend/pkg/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAPIRoutes(r *gin.Engine) {
	var v1 *gin.RouterGroup
	if len(config.Get("app.api_domain")) == 0 {
		v1 = r.Group("/api/v1")
	} else {
		v1 = r.Group("/v1")
	}

	{

		// 管理员路由 manager-api
		managerRouterGroup := v1.Group("/manager-api")
		{
			mc := new(auth.ManagerAuthController)
			managerRouterGroup.POST("/login", mc.Login)
			managerRouterGroup.POST("/logout", middlewares.ManagerJWT(), mc.Logout)
			managerRouterGroup.GET("/", middlewares.ManagerJWT(), func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "manager",
				})
			})
		}

		// 用户路由
		userRouterGroup := v1.Group("/user-api")
		{
			userRouterGroup.GET("/users/:id", func(c *gin.Context) {
				id := c.Param("id")
				c.JSON(http.StatusOK, gin.H{
					"message": "users",
					"data":    id,
				})
			})
		}

		// 比赛路由
		gameRouterGroup := v1.Group("/game-api")
		{
			gc := new(v1API.GameAPIController)
			// 比赛列表
			gameRouterGroup.GET("/games", gc.Index)
			// 添加比赛
			gameRouterGroup.POST("/create", gc.Create)
			// 更新比赛
			gameRouterGroup.PUT("/update/:id", gc.Update)
			// 删除比赛
			gameRouterGroup.DELETE("/delete/:id", gc.Delete)
		}

		// 赛题路由 question-api
		questionRouterGroup := v1.Group("/question-api")
		{
			qc := new(v1API.QuestionAPIController)
			questionRouterGroup.GET("/all", qc.Index)
			questionRouterGroup.GET("/single/:id", qc.Show)
			questionRouterGroup.POST("/create", qc.Create)
			questionRouterGroup.DELETE("/delete/:id", qc.Delete)
			questionRouterGroup.PATCH("/update/:id", qc.Update)
		}

		// 赛题路由 challenge-api
		challengeRouterGroup := v1.Group("/challenge-api")
		{
			qc := new(v1API.QuestionAPIController)
			challengeRouterGroup.GET("/all", qc.Index)
			challengeRouterGroup.GET("/item/:id", qc.Index)
		}
	}
}
