package routers

import (
	"go_crud/pkg/setting"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter 创建路由
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})

	return r
}
