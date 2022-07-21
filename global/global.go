package global

import "tobeg/config"

var (
	Config    *config.Config = &config.Config{}
	PayStatus                = make(map[string]bool)
)
