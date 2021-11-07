package web

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"monitorX/internal/web/controller"
	"net/http"
)

func userRouter(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.GET("/info", func(c *gin.Context) {
			requestUser := RequestUsername(c
			c.JSON(200, controller.UserInfo(requestUser))
		})
		user.GET("/node", func(c *gin.Context) {
			requestUser := RequestUsername(c)
			c.JSON(200, controller.UserInfo(requestUser))
		})
	}
}

func pingRouter(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "pong",
		})
	})
}
func Start(host string, port, timeout int) {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	pingRouter(router)
	router.Use(Auth(router, timeout).MiddlewareFunc())
	userRouter(router)
	router.Run(fmt.Sprintf("%s:%d", host, port))
}
