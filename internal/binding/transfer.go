package binding

type Transfer struct {
	Sender   uint32  `json:"sender" binding:"required"`
	Receiver uint32  `json:"receiver" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`
}
