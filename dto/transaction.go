package dto

import "time"

type Transaction struct {
	Id              string
	Name            string
	Number          string
	ExpirationMonth int32
	ExpirationYear  int32
	CVV             int32
	Amount          float64
	Description     string
	Store           string
	CreatedAt       time.Time
}
