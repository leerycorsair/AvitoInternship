package orders_repository

import (
	"database/sql"
	"sync"
	"time"

	"AvitoInternship/internal/models"
)

type OrderRepo struct {
	mutex sync.RWMutex
	c     *sql.DB
}

func CreateOrdersRepo(c *sql.DB) *OrderRepo {
	return &OrderRepo{c: c}
}

func (repo *OrderRepo) CreateOrder(order models.Order) error {
	repo.mutex.Lock()
	curTime := order.CreatedAt.Format("2006-01-02 15:04:05")
	_, err := repo.c.Exec("insert into orders(`order_id`, `user_id`, `service_id`, `value`, "+
		"`created_at`, `comments`, `status`) values (?, ?, ?, ?, ?, ?, ?);", order.OrderId, order.UserId, order.ServiceId,
		order.Value, curTime, order.Comments, order.Status)
	repo.mutex.Unlock()
	return err
}

func (repo *OrderRepo) GetAllOrders() ([]models.Order, error) {
	allOrders := make([]models.Order, 0)

	repo.mutex.Lock()
	rows, err := repo.c.Query("select order_id, user_id, service_id, value, created_at, comments, status from orders;")
	repo.mutex.Unlock()

	if err == nil {
		for rows.Next() {
			newOrder := &models.Order{}
			var orderTime string
			err = rows.Scan(&newOrder.OrderId, &newOrder.UserId, &newOrder.ServiceId,
				&newOrder.Value, &orderTime, &newOrder.Comments, &newOrder.Status)
			newOrder.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", orderTime)
			if err != nil {
				break
			} else {
				allOrders = append(allOrders, *newOrder)
			}
		}
	}
	return allOrders, err
}

func (repo *OrderRepo) GetOrderByID(orderID, userID, serviceId int) (models.Order, error) {
	foundOrder := models.Order{}

	repo.mutex.Lock()
	row := repo.c.QueryRow("select order_id, user_id, service_id, value, created_at, comments, status from orders where order_id = ? and user_id = ? and service_id = ?;", orderID, userID, serviceId)
	repo.mutex.Unlock()

	var orderTime string

	err := row.Scan(&foundOrder.OrderId, &foundOrder.UserId, &foundOrder.ServiceId,
		&foundOrder.Value, &orderTime, &foundOrder.Comments, &foundOrder.Status)

	foundOrder.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", orderTime)

	return foundOrder, err
}

func (repo *OrderRepo) GetUserOrders(userID int) ([]models.Order, error) {
	allOrders := make([]models.Order, 0)

	repo.mutex.Lock()
	rows, err := repo.c.Query("select order_id, user_id, service_id, value, created_at, comments, status from orders where user_id = ?;", userID)
	repo.mutex.Unlock()

	if err == nil {
		for rows.Next() {
			newOrder := &models.Order{}

			var orderTime string
			err = rows.Scan(&newOrder.OrderId, &newOrder.UserId, &newOrder.ServiceId,
				&newOrder.Value, &orderTime, &newOrder.Comments, &newOrder.Status)
			newOrder.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", orderTime)
			if err != nil {
				break
			} else {
				allOrders = append(allOrders, *newOrder)
			}
		}
	}
	return allOrders, err
}

func (repo *OrderRepo) GetServiceOrders(serviceId int) ([]models.Order, error) {
	allOrders := make([]models.Order, 0)

	repo.mutex.Lock()
	rows, err := repo.c.Query("select order_id, user_id service_id, value, created_at, comments, status from orders where service_id = ?;", serviceId)
	repo.mutex.Unlock()

	if err == nil {
		for rows.Next() {
			newOrder := &models.Order{}
			var orderTime string
			err = rows.Scan(&newOrder.OrderId, &newOrder.UserId, &newOrder.ServiceId,
				&newOrder.Value, &orderTime, &newOrder.Comments, &newOrder.Status)
			newOrder.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", orderTime)
			if err != nil {
				break
			} else {
				allOrders = append(allOrders, *newOrder)
			}
		}
	}
	return allOrders, err
}

func (repo *OrderRepo) ChangeOrderStatus(orderID, userID, serviceId int, status int) error {
	repo.mutex.Lock()
	_, err := repo.c.Exec("update orders set status = ? where order_id = ? and user_id = ? and service_id = ?", status, orderID, userID, serviceId)
	repo.mutex.Unlock()
	return err
}

func (repo *OrderRepo) GetValueOfFinishedServices(month, year int) ([]models.Report, error) {
	allServices := make([]models.Report, 0)

	repo.mutex.Lock()
	rows, err := repo.c.Query("select service_id, SUM(value) from orders where status = 2 and month(created_at) = ? and year(created_at) = ? group by service_id;", month, year)
	repo.mutex.Unlock()

	if err == nil {
		for rows.Next() {
			newServiceReport := models.Report{}
			err = rows.Scan(&newServiceReport.ServiceId, &newServiceReport.Sum)
			if err != nil {
				break
			} else {
				allServices = append(allServices, newServiceReport)
			}
		}
	}
	return allServices, err
}
