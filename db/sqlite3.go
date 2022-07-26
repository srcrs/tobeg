package db

import (
	"sort"
	"strconv"
	"time"
	"tobeg/model"
	"tobeg/utils"

	"go.uber.org/zap"
)

// 插入支付记录
func InsertFlow(flow model.Flow) {
	zap.S().Info(flow)
	db := utils.GenerateDB()
	defer db.Close()

	stmt, err := db.Prepare("insert into flow (id, tradeId, outTradeId, username, amount, status, createTime, updateTime)" +
		" values (NULL, ?, ?, ?, ?, ?, ?, ?) on conflict(outTradeId) do update set status=excluded.status, updateTime=excluded.updateTime;")
	checkErr(err)

	timeStamp := time.Now().Unix()
	res, err := stmt.Exec(flow.TradeId, flow.OutTradeId, flow.UserName, flow.Amount, flow.Status, timeStamp, timeStamp)
	zap.S().Info("插入结果:", res)
}

// flowId为0 从最后开始找
// count查找个数
func GetFlowList(flowId, count int, pageType string) []model.Flow {
	db := utils.GenerateDB()
	defer db.Close()

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
	zap.S().Info(flowId, count)
	rows, err := db.Query("select * from flow where " + cond + " order by id " + sorted + " limit " + strconv.Itoa(count))
	checkErr(err)
	defer rows.Close()

	arr := make([]model.Flow, 0)

	for rows.Next() {
		flowModel := model.Flow{}
		rows.Scan(&flowModel.Id, &flowModel.TradeId, &flowModel.OutTradeId, &flowModel.UserName, &flowModel.Amount, &flowModel.Status, &flowModel.CreateTime, &flowModel.UpdateTime)
		arr = append(arr, flowModel)
	}
	err = rows.Err()
	checkErr(err)

	formatFlow(arr)

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Id > arr[j].Id
	})
	return arr
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

func CountFlow() int {
	db := utils.GenerateDB()
	defer db.Close()

	rows, err := db.Query("select count(id) count from flow;")
	checkErr(err)
	defer rows.Close()

	var count int
	for rows.Next() {
		rows.Scan(&count)
	}
	return count
}

func checkErr(err error) {
	if err != nil {
		zap.S().Infof("数据库操作失败: %s", err.Error())
		return
	}
}
