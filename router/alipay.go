package router

import (
	"github.com/gin-gonic/gin"
	"tobeg/api"
)

func InitAlipayRouter(Router *gin.RouterGroup) {
	AlipayRouter := Router.Group("alipay")
	{
		AlipayRouter.POST("trade/precreate", api.TradePrecreate)
		AlipayRouter.POST("trade/paysuccess", api.TradePaySuccess)
		AlipayRouter.POST("trade/tradeclose", api.TradeClose)
		AlipayRouter.POST("trade/query", api.TradeQuery)
	}
}
