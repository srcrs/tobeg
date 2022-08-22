package db

import (
	"sort"
	"strconv"
	"time"

	"tobeg/global"
	"tobeg/model"
	"tobeg/utils"

	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

func InsertFlow(flow *model.Flow) {
	zap.S().Info(flow)
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	flow.CreateTime, flow.UpdateTime = timeStamp, timeStamp
	result := global.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "outTradeId"}},
		DoUpdates: clause.AssignmentColumns([]string{"status", "updateTime"}),
	}).Create(flow)

	if result.Error != nil {
		zap.S().Info("插入支付流水异常: %s", result.Error.Error())
	}
}

// flowId为0 从最后开始找
// count查找个数
func GetFlowList(flowId, count int, pageType string) []model.Flow {
	var flows []model.Flow

	cond, sorted := "", ""
	if flowId == 0 {
		cond, sorted = "1=1", "desc"
	} else {
		if pageType == "prev" {
			cond, sorted = "id > "+strconv.Itoa(flowId), ""
		} else if pageType == "next" {
			cond, sorted = "id < "+strconv.Itoa(flowId), "desc"
		}
	}

	global.DB.Limit(count).Order("id " + sorted).Where(cond).Find(&flows)

	formatFlow(flows)

	sort.Slice(flows, func(i, j int) bool {
		return flows[i].Id > flows[j].Id
	})

	return flows
}

func formatFlow(arr []model.Flow) {
	for i := 0; i < len(arr); i++ {
		//timeStamp, _ := strconv.ParseInt(arr[i].CreateTime, 10, 64)
		//arr[i].CreateTime = utils.UnixToStr(timeStamp, "2006-01-02 15:04:05")
		timeStamp, _ := strconv.ParseInt(arr[i].UpdateTime, 10, 64)
		timeStamp += 8 * 3600
		arr[i].UpdateTime = utils.UnixToStr(timeStamp, "2006-01-02 15:04:05")
	}
}

func CountFlow() int64 {
	var count int64
	global.DB.Model(&model.Flow{}).Count(&count)
	return count
}
