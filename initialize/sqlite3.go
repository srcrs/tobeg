package initialize

import (
	"os"

	"go.uber.org/zap"
	"tobeg/utils"
)

func InitSQLite3() {

	// 数据库已存在
	if _, err := os.Stat("./flow.db"); err == nil {
		zap.S().Info("数据库表已存在")
		return
	}

	db := utils.GenerateDB()
	defer db.Close()

	sqlStmt := `CREATE TABLE flow (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    tradeId VARCHAR(32) NOT NULL,
    outTradeId VARCVHAR(32) NOT NULL,
    userName VARCHAR(64) NULL,
    amount INT NOT NULL,
    status VARCHAR(32) NOT NULL,
    createTime BIGINT NOT NULL,
    updateTime BIGINT NOT NULL
    );
    CREATE UNIQUE INDEX idx_outtradeid on flow(outTradeId);
  `

	_, err := db.Exec(sqlStmt)
	if err != nil {
		zap.S().Errorf("数据库执行错误%q: %s\n", err, sqlStmt)
		return
	}

}
