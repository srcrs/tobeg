package initialize

import (
	"tobeg/global"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile("config.yml")
	if err := v.ReadInConfig(); err != nil {
		zap.S().Panicf("读取配置文件失败: %s", err.Error())
	}
	v.ReadInConfig()
	v.Unmarshal(global.Config)
	zap.S().Infof("配置信息: %v", global.Config)
}
