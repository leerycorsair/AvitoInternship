package orders_controller

import (
	"AvitoInternship/internal/models"
	"errors"
)

var (
	OrderNotFound       = errors.New("Order Not Found")
	OrderIsAlreadyExist = errors.New("Order Is Already Exist")
	GetOrderError       = errors.New("Bad Order Get")
	WrongStatusError    = errors.New("Status Isn't Right To Change Order Status")
	BadMonthError       = errors.New("Bad Month")
	BadYearError        = errors.New("Bad Year")
)

type OrdersControllerInterface interface {
	GetOrder(orderID, userID, serviceID int) (models.Order, error)
	CreateNewOrder(orderID, userID, serviceID int, value float64, comments string) error
	CheckOrderIsExist(orderID, userID, serviceID int) (bool, error)
	ReserveOrder(orderID, userID, serviceID int) error
	FinishOrder(orderID, userID, serviceID int) error
	ReturnOrder(orderID, userID, serviceID int) (float64, error)
	GetFinanceReports(month, year int) ([]models.Report, error)
}
