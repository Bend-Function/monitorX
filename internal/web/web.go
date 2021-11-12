package web

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"log"
	"monitorX/internal/database"
	"monitorX/internal/web/controller"
	"net/http"
	"strconv"
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
		user.GET("", func(c *gin.Context) {
			requestUser := RequestUsername(c)
			c.JSON(200, controller.UserInfo(requestUser))
		})
		user.POST("", func(c *gin.Context) {
			var newUser *database.User
			err := c.Bind(&newUser)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			}
		})
		user.GET("/node/list", func(c *gin.Context) {
			requestUser := RequestUsername(c)
			c.JSON(200, controller.UserNode(requestUser))
		})
		user.GET("/node/data/:id/:time_period", func(c *gin.Context) {
			requestUser := RequestUsername(c)
			nodeID, err := strconv.Atoi(c.Param("id"))
			timePeriod := c.Param("time_period")
			if err != nil {
				c.JSON(http.StatusInternalServerError, "id is not int")
			}
			c.JSON(200, controller.QueryNodeData(nodeID, requestUser, timePeriod))
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
