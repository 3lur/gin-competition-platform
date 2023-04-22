package bootstrap

import (
	"competition-backend/app/http/middlewares"
	"competition-backend/pkg/config"
	"competition-backend/pkg/console"
	"competition-backend/pkg/logger"
	"competition-backend/router"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetupRouter() {
	r := gin.New()

	gin.SetMode(gin.ReleaseMode)

	router.RegisterAPIRoutes(r)

	registerGlobalMiddleware(r)

	setup404Handler(r)

	err := r.Run(":" + config.Get("app.port"))
	if err != nil {
		logger.Error(err.Error())
		console.Exit("Unable to start server, error:" + err.Error())
	}
}

func registerGlobalMiddleware(r *gin.Engine) {
	r.Use(
		middlewares.Logger(),
		middlewares.ForceUA(),
	)
}

func setup404Handler(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		as := c.Request.Header.Get("Accept")
		if strings.Contains(as, "text/html") {
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "路由未定义，请确认 url 和请求方法是否正确",
			})
		}
	})
}
