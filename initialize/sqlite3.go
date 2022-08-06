package initialize

import (
	"os"
	"tobeg/global"
	"tobeg/model"

	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitSQLite3() {

	// 数据库已存在
	if _, err := os.Stat("./tobeg.db"); err == nil {
		zap.S().Info("数据库表已存在")
	} else {
		zap.S().Info("正在初始化数据库")
	}

	var err error
	global.DB, err = gorm.Open(sqlite.Open("tobeg.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		zap.S().Errorf("连接数据库失败: %s", err.Error())
		return
	}

	global.DB.AutoMigrate(&model.Flow{})

}
