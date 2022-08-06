package global

import (
	"gorm.io/gorm"
	"tobeg/config"
)

var (
	Config    *config.Config = &config.Config{}
	PayStatus                = make(map[string]bool)
	DB        *gorm.DB
)
