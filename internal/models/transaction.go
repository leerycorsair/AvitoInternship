package models

import "time"

const (
	Add      = 0
	Reserve  = 1
	Return   = 2
	Transfer = 3
)

type Transaction struct {
	TransactionId   int
	UserId          int
	TransactionType int
	Value           float64
	CreatedAt       time.Time
	ActionComments  string
	AddComments     string
}
