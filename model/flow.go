package model

type Flow struct {
	Id         int     `gorm:"column:id;primarykey"`
	TradeId    string  `gorm:"column:tradeId;type:varchar(32);not null"`
	OutTradeId string  `gorm:"column:outTradeId;Index:idx_outtradeid;unique;type:varchar(32);not null"`
	Amount     float64 `gorm:"column:amount;type:int;not null"`
	UserName   string  `gorm:"column:userName;type:varchar(64);not null"`
	Status     string  `gorm:"column:status;type:varchar(32);not null"`
	CreateTime string  `gorm:"column:createTime;type:bigint;not null"`
	UpdateTime string  `gorm:"column:updateTime;type:bigint;not null"`
}
