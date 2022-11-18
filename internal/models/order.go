package models

import "time"

const (
	REGISTRATED = 0
	RESERVED    = 1
	FINISHED    = 2
	RETURNED    = 3
)

type Order struct {
	UserId    int
	ServiceId int
	OrderId   int
	Value     float64
	CreatedAt time.Time
	Comments  string
	Status    int
}
