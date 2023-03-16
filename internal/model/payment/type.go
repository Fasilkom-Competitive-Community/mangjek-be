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
		ID     string
		Amount float64
		Status Status
		Method Method

		QrString string

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

	QRIS struct {
		ID         string
		ExternalID string
		Amount     float64
		QrString   string
		Status     string
	}

	AddQRIS struct {
		ExternalID string
		Amount     float64
	}
)
