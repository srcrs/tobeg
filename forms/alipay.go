package forms

type TradePrecreate struct {
	Subject string  `form:"subject" json:"subject"`
	Amount  float32 `form:"amount" json:"amount" binding:"required,min=0.01,max=1000"`
}
