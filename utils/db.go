package utils

import (
	"database/sql"
	"time"

	"go.uber.org/zap"

	_ "github.com/mattn/go-sqlite3"
)

// 获取DB操作实例
func GenerateDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./flow.db")
	checkErr(err)
	return db
}

func UnixToStr(timeUnix int64, layout string) string {
	timeStr := time.Unix(timeUnix, 0).Format(layout)
	return timeStr
}

func checkErr(err error) {
	if err != nil {
		zap.S().Errorf("sqlite实例创建失败: %s", err.Error())
	}
}
