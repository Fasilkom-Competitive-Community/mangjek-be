package payment

import "time"

const (
	CashMethod Method = "CASH"
	QRISMethod Method = "QRIS"

	PaidStatus   Status = "PAID"
	UnpaidStatus Status = "UNPAID"
)

type (
	Method string
	Status string

	Payment struct {
		ID      string
		OrderID string
		Amount  float64
		Status  Status
		Method  Method

		CreatedAt time.Time
		UpdatedAt time.Time
	}

	AddPayment struct {
		ID     string
		Amount float64
		Status Status
		Method Method

		QrString string
	}
)
