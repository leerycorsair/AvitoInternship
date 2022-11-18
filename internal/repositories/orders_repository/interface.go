package orders_repository

import "AvitoInternship/internal/models"

type OrdersRepoInterface interface {
	CreateOrder(order models.Order) error
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(orderID, userID, serviceType int) (models.Order, error)
	GetUserOrders(userID int) ([]models.Order, error)
	GetServiceOrders(serviceType int) ([]models.Order, error)
	ChangeOrderStatus(orderID, userID, serviceType int, status int) error
	GetValueOfFinishedServices(month, year int) ([]models.Report, error)
}
