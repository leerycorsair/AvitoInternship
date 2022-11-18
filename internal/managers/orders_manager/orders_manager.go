package orders_manager

import (
	"AvitoInternship/internal/controllers/orders_controller"
	"AvitoInternship/internal/controllers/transactions_controller"
	"AvitoInternship/internal/controllers/users_controller"
)

type OrdersManager struct {
	usersController        users_controller.UsersControllerInterface
	ordersController       orders_controller.OrdersControllerInterface
	transactionsController transactions_controller.TransactionsControllerInterface
}

func CreateOrdersManager(uc users_controller.UsersControllerInterface, oc orders_controller.OrdersControllerInterface,
	tc transactions_controller.TransactionsControllerInterface) *OrdersManager {
	return &OrdersManager{
		usersController:        uc,
		ordersController:       oc,
		transactionsController: tc,
	}
}

func (m *OrdersManager) ReserveService(userID, orderID, serviceID int, value float64, comment string) error {
	canUserBuy, err := m.checkUserCanBuyService(userID, value)
	if err == nil {
		if canUserBuy {
			err = m.createOrder(userID, orderID, serviceID, value, comment)
			if err == nil {
				err = m.ordersController.ReserveOrder(orderID, userID, serviceID)
				if err == nil {
					err = m.usersController.SpendMoney(userID, value)
					if err == nil {
						err = m.transactionsController.AddNewRecordReserveService(userID, value, serviceID, comment)
					}
				}
			}
		} else {
			err = users_controller.NotEnoughMoneyErr
		}
	}
	return err
}

func (m *OrdersManager) checkUserCanBuyService(userID int, value float64) (bool, error) {
	var canBuy bool
	isUserExist, err := m.usersController.CheckUserIsExist(userID)
	if err == nil {
		if value <= 0 {
			err = users_controller.NegValueError
		} else if !isUserExist {
			err = users_controller.UserNotExistErr
		} else {
			canBuy, err = m.usersController.CheckAbleToBuyService(userID, value)
		}
	}
	return canBuy, err
}

func (m *OrdersManager) createOrder(userID, orderID, serviceID int, value float64, comment string) error {
	isOrderExist, err := m.ordersController.CheckOrderIsExist(orderID, userID, serviceID)
	if err == nil {
		if !isOrderExist {
			err = m.ordersController.CreateNewOrder(orderID, userID, serviceID, value, comment)
		} else {
			err = orders_controller.OrderIsAlreadyExist
		}
	}
	return err
}

func (m *OrdersManager) AcceptReserve(userID, orderID, serviceID int) error {
	isOrderExist, err := m.ordersController.CheckOrderIsExist(orderID, userID, serviceID)
	if err == nil {
		if isOrderExist {
			err = m.ordersController.FinishOrder(orderID, userID, serviceID)
			if err == nil {
				m.transferMoneyToCompanyAccount()
			}
		} else {
			err = orders_controller.OrderNotFound
		}
	}
	return err
}

func (*OrdersManager) transferMoneyToCompanyAccount() {}

func (m *OrdersManager) CancelReserve(userID, orderID, serviceID int, comment string) error {
	isOrderExist, err := m.ordersController.CheckOrderIsExist(orderID, userID, serviceID)
	if err == nil {
		if isOrderExist {
			value, cancelOrderErr := m.ordersController.ReturnOrder(orderID, userID, serviceID)
			if cancelOrderErr == nil {
				err = m.usersController.DonateMoney(userID, value)
				if err == nil {
					err = m.transactionsController.AddNewRecordReturnService(userID, value, serviceID, comment)
				}
			} else {
				err = cancelOrderErr
			}
		} else {
			err = orders_controller.OrderNotFound
		}
	}
	return err
}
