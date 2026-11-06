package entity

type OrderItem struct {
	ID        int
	OrderID   int
	ProductID int
	Quantity  int
	UnitPrice float64
}
