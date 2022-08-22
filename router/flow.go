package router

import (
	"github.com/gin-gonic/gin"

	"tobeg/api"
)

func InitFlowRouter(Router *gin.RouterGroup) {
	AlipayRouter := Router.Group("flow")
	{
		AlipayRouter.POST("/getflowlist", api.GetFlowList)
	}
}
