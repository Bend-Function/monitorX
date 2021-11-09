package web

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"log"
	"monitorX/internal/database"
	"monitorX/internal/web/controller"
	"net/http"
)

func nodeRouter(router *gin.Engine) {
	node := router.Group("/node")
	{
		node.POST("/update/data", func(c *gin.Context) {
			var nodeData database.NodeData
			err = c.Bind(&nodeData)
			nodePassword := c.Query("password")
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			}
			c.JSON(200, controller.UpdateData(&nodeData, nodePassword))
		})
	}
}

func userRouter(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.GET("/info", func(c *gin.Context) {
			requestUser := RequestUsername(c)
			c.JSON(200, controller.UserInfo(requestUser))
		})
		user.GET("/node", func(c *gin.Context) {
			requestUser := RequestUsername(c)
			c.JSON(200, controller.UserNode(requestUser))
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
	nodeRouter(router)
	router.Use(Auth(router, timeout).MiddlewareFunc())
	userRouter(router)
	err = router.Run(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatal(err)
	}
}
