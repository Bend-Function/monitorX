package web

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"monitorX/internal/web/controller"
	"net/http"
)

func userRouter(router *gin.Engine) {
	user := router.Group("/trojan/user")
	{
		user.GET("", func(c *gin.Context) {
			requestUser := RequestUsername(c)
			c.JSON(200, controller.UserInfo(requestUser))
		})
		//user.GET("/page", func(c *gin.Context) {
		//	curPageStr := c.DefaultQuery("curPage", "1")
		//	pageSizeStr := c.DefaultQuery("pageSize", "10")
		//	curPage, _ := strconv.Atoi(curPageStr)
		//	pageSize, _ := strconv.Atoi(pageSizeStr)
		//	c.JSON(200, controller.PageUserList(curPage, pageSize))
		//})
		//user.POST("", func(c *gin.Context) {
		//	username := c.PostForm("username")
		//	password := c.PostForm("password")
		//	c.JSON(200, controller.CreateUser(username, password))
		//})
		//user.POST("/update", func(c *gin.Context) {
		//	sid := c.PostForm("id")
		//	username := c.PostForm("username")
		//	password := c.PostForm("password")
		//	id, _ := strconv.Atoi(sid)
		//	c.JSON(200, controller.UpdateUser(uint(id), username, password))
		//})
		//user.POST("/expire", func(c *gin.Context) {
		//	sid := c.PostForm("id")
		//	sDays := c.PostForm("useDays")
		//	id, _ := strconv.Atoi(sid)
		//	useDays, _ := strconv.Atoi(sDays)
		//	c.JSON(200, controller.SetExpire(uint(id), uint(useDays)))
		//})
		//user.DELETE("/expire", func(c *gin.Context) {
		//	sid := c.Query("id")
		//	id, _ := strconv.Atoi(sid)
		//	c.JSON(200, controller.CancelExpire(uint(id)))
		//})
		//user.DELETE("", func(c *gin.Context) {
		//	stringId := c.Query("id")
		//	id, _ := strconv.Atoi(stringId)
		//	c.JSON(200, controller.DelUser(uint(id)))
		//})
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
