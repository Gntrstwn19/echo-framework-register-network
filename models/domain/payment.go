package domain

import "time"

type Payment struct {
	payment_id     int
	customer_id    int
	payment_date   time.Time
	payment_amount int
	payment_status string
	payment_method string
}

func (Payment) TableName() string {
	return "payment"
}
