package api

import (
	"net/http"
	"strconv"

	"tobeg/db"

	"github.com/gin-gonic/gin"
)

func GetFlowList(ctx *gin.Context) {
	flowId, _ := strconv.Atoi(ctx.PostForm("flowId"))
	count, _ := strconv.Atoi(ctx.PostForm("count"))
	pageType := ctx.PostForm("type")

	end := false
	flows := db.GetFlowList(flowId, count, pageType)
	headFlowId := 0
	endFlowId := 0
	if len(flows) > 0 {
		headFlowId = flows[0].Id
		endFlowId = flows[len(flows)-1].Id
	}
	if len(flows) < count {
		end = true
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "成功",
		"data": map[string]any{
			"total":      db.CountFlow(),
			"flows":      flows,
			"isEnd":      end,
			"headFlowId": headFlowId,
			"endFlowId":  endFlowId,
		},
	})
}
