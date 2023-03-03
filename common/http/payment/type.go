package payment

type (
	AddPayment struct {
		Amount float64 `json:"amount"`
		Method string  `json:"method" binding:"required,oneof=CASH QRIS"`
	}
)
