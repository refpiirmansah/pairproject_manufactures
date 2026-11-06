package entity

import "time"

type Assembly struct {
	ID          int
	OrderItemID int
	Quantity    int
	SatartDate  *time.Time
	EndaDate    *time.Time
	Status      string
}
