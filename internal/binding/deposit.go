package binding

type Deposit struct {
	Receiver uint32  `json:"receiver" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`
}
