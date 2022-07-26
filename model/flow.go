package model

type Flow struct {
	Id         int
	TradeId    string
	OutTradeId string
	Amount     float64
	UserName   string
	Status     string
	CreateTime string
	UpdateTime string
}
