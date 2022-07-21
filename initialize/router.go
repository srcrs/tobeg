package initialize

import (
	"tobeg/middlerwares"
	"tobeg/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	// 解决跨域
	Router.Use(middlerwares.Cors())

	ApiGroup := Router.Group("v1")
	router.InitPayRouter(ApiGroup)

	return Router
}
