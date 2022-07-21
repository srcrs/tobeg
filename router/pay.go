package router

import (
	"github.com/gin-gonic/gin"
	"tobeg/api"
)

func InitPayRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("alipay")
	{
		BaseRouter.POST("trade/precreate", api.TradePrecreate)
		BaseRouter.POST("trade/paysuccess", api.TradePaySuccess)
		BaseRouter.POST("trade/tradeclose", api.TradeClose)
		BaseRouter.POST("trade/query", api.TradeQuery)
	}
}
