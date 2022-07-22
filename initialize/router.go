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

	// 加载tmplates目录下模版文件
	Router.LoadHTMLGlob("./templates/*")

	Router.Static("/public", "./public")

	ApiGroup := Router.Group("v1")
	router.InitAlipayRouter(ApiGroup)
	router.InitTmplRouter(Router.Group(""))
	return Router
}
