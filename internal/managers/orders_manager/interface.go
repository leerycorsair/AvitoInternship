package orders_manager

type OrdersManagerInterface interface {
	ReserveService(userID, orderID, serviceID int, value float64, comments string) error
	AcceptReserve(userID, orderID, serviceID int) error
	CancelReserve(userID, orderID, serviceID int, comments string) error
}
