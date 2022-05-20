package model

type Wallet struct {
	ID     uint32  `json:"id" gorm:"primaryKey"`
	Amount float64 `json:"amount"`
}
