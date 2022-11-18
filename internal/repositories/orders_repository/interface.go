package orders_repository

import "AvitoInternship/internal/models"

type OrdersRepoInterface interface {
	CreateOrder(order models.Order) error
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(orderID, userID, serviceId int) (models.Order, error)
	GetUserOrders(userID int) ([]models.Order, error)
	GetServiceOrders(serviceId int) ([]models.Order, error)
	ChangeOrderStatus(orderID, userID, serviceId int, status int) error
	GetValueOfFinishedServices(month, year int) ([]models.Report, error)
}
