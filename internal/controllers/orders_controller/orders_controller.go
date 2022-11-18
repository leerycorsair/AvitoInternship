package orders_controller

import (
	"AvitoInternship/internal/models"
	"AvitoInternship/internal/repositories/orders_repository"

	"database/sql"
	"sync"
	"time"
)

type OrdersController struct {
	mutex sync.RWMutex
	repo  orders_repository.OrdersRepoInterface
}

func CreateNewOrdersController(repo orders_repository.OrdersRepoInterface) *OrdersController {
	return &OrdersController{mutex: sync.RWMutex{}, repo: repo}
}

func (c *OrdersController) GetOrder(orderID, userID, serviceID int) (models.Order, error) {
	return c.repo.GetOrderByID(orderID, userID, serviceID)
}

func (c *OrdersController) CreateNewOrder(orderID, userID, serviceID int, value float64, comment string) error {
	isExist, err := c.CheckOrderIsExist(orderID, userID, serviceID)
	if err == nil {
		if !isExist {
			newOrder := models.Order{
				OrderId:   orderID,
				UserId:    userID,
				ServiceId: serviceID,
				Value:     value,
				CreatedAt: time.Now(),
				Comments:  comment,
				Status:    models.REGISTRATED,
			}
			c.mutex.Lock()
			err = c.repo.CreateOrder(newOrder)
			c.mutex.Unlock()
		} else {
			err = OrderIsAlreadyExist
		}
	}
	return err
}

func (c *OrdersController) CheckOrderIsExist(orderID, userID, serviceID int) (bool, error) {
	var result bool

	c.mutex.Lock()
	foundOrder, err := c.repo.GetOrderByID(orderID, userID, serviceID)
	c.mutex.Unlock()

	if err == nil {
		if foundOrder.OrderId == orderID && foundOrder.UserId == userID && foundOrder.ServiceId == serviceID {
			result = true
		} else {
			result = false
		}
	} else if err == sql.ErrNoRows {
		err = nil
		result = false
	}
	return result, err
}

func (c *OrdersController) ReserveOrder(orderID, userID, serviceID int) error {
	isOrderExist, err := c.CheckOrderIsExist(orderID, userID, serviceID)

	if err == nil {
		if isOrderExist {
			curOrder, getOrderErr := c.GetOrder(orderID, userID, serviceID)
			if getOrderErr == nil && curOrder.Status == models.REGISTRATED {
				err = c.repo.ChangeOrderStatus(orderID, userID, serviceID, models.RESERVED)
			} else {
				if getOrderErr != nil {
					err = GetOrderError
				} else {
					err = WrongStatusError
				}
			}
		}
	}
	return err
}

func (c *OrdersController) FinishOrder(orderID, userID, serviceID int) error {
	isOrderExist, err := c.CheckOrderIsExist(orderID, userID, serviceID)

	if err == nil {
		if isOrderExist {
			curOrder, getOrderErr := c.GetOrder(orderID, userID, serviceID)
			if getOrderErr == nil && curOrder.Status == models.RESERVED {
				err = c.repo.ChangeOrderStatus(orderID, userID, serviceID, models.FINISHED)
			} else {
				if getOrderErr != nil {
					err = GetOrderError
				} else {
					err = WrongStatusError
				}
			}
		}
	}
	return err
}

func (c *OrdersController) ReturnOrder(orderID, userID, serviceID int) (float64, error) {
	var value float64
	isOrderExist, err := c.CheckOrderIsExist(orderID, userID, serviceID)

	if err == nil {
		if isOrderExist {
			curOrder, getOrderErr := c.GetOrder(orderID, userID, serviceID)
			if getOrderErr == nil && curOrder.Status == models.RESERVED {
				err = c.repo.ChangeOrderStatus(orderID, userID, serviceID, models.RETURNED)
				if err == nil {
					value = curOrder.Value
				}
			} else {
				if getOrderErr != nil {
					err = GetOrderError
				} else {
					err = WrongStatusError
				}
			}
		}
	}
	return value, err
}

func (c *OrdersController) GetFinanceReports(month, year int) ([]models.Report, error) {
	var report = make([]models.Report, 0)
	var err error = nil
	if month < 1 || month > 2 {
		err = BadMonthError
	} else if year < 1960 || year > 2022 {
		err = BadYearError
	} else {
		report, err = c.repo.GetValueOfFinishedServices(month, year)
	}
	return report, err
}
