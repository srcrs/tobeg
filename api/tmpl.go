package api

import (
	"net/http"
	"tobeg/global"

	"github.com/gin-gonic/gin"
)

// 首页
func WebIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"baseUrl": global.Config.BaseConfig.Url,
		"toSells": global.Config.BaseConfig.ToSells,
		"favicon": global.Config.BaseConfig.Favicon,
		"title":   global.Config.BaseConfig.Title,
	})
}
