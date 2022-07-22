package router

import (
	"github.com/gin-gonic/gin"
	"tobeg/api"
)

func InitTmplRouter(Router *gin.RouterGroup) {
	TmplRouter := Router.Group("/")
	{
		TmplRouter.GET("/", api.WebIndex)
	}
}
