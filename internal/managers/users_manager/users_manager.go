package users_manager

import (
	"AvitoInternship/internal/controllers/users_controller"
	"fmt"
)

type UsersManager struct {
	usersController        users_controller.UsersControllerInterface
	transactionsController tc.TransactionsControllerInterface
}

func CreateNewUsersManager(usersController users_controller.UsersControllerInterface,
	transactionsController tc.TransactionsControllerInterface) *UsersManager {
	return &UsersManager{
		usersController:        usersController,
		transactionsController: transactionsController,
	}
}

func (m *UsersManager) AddBalance(userID int, value float64, comments string) error {
	isUserExist, err := m.usersController.CheckUserIsExist(userID)
	if err == nil {
		if !isUserExist {
			err = m.usersController.CreateNewUser(userID)
		}
		if err == nil {
			err = m.usersController.DonateMoney(userID, value)
			if err == nil {
				err = m.transactionsController.AddNewRecordAddBalance(userID, value, comments)
			}
		}
	}
	return err
}

func (m *UsersManager) GetUserBalance(userID int) (float64, error) {
	isUserExist, err := m.usersController.CheckUserIsExist(userID)
	if err == nil {
		if isUserExist {
			return m.usersController.CheckBalance(userID)
		} else {
			return 0, users_controller.UserNotExistErr
		}
	} else {
		return 0, fmt.Errorf("managers call: %w", err)
	}
}

func (m *UsersManager) Transfer(srcUserID, dstUserID int, value float64, comment string) error {
	isUserExist, err := m.checkAllUsersAreExists(srcUserID, dstUserID)
	if err == nil {
		if isUserExist {
			canBuy, checkBuyErr := m.checkUserCanBuyService(srcUserID, value)
			if checkBuyErr == nil {
				if canBuy {
					err = m.usersController.SpendMoney(srcUserID, value)
					if err == nil {
						err = m.usersController.DonateMoney(dstUserID, value)
						if err == nil {
							err = m.makeReportsForAllUsers(srcUserID, dstUserID, value, comment)
						}
					}
				} else {
					err = users_controller.NotEnoughMoneyErr
				}
			}
		} else {
			err = users_controller.UserNotExistErr
		}
	}
	return err
}

func (m *UsersManager) checkAllUsersAreExists(srcUserID, dstUserID int) (bool, error) {
	var result bool
	var err error
	isUserExist, firstCheck := m.usersController.CheckUserIsExist(srcUserID)
	if firstCheck == nil {
		if isUserExist {
			isSecUserExist, secCheckErr := m.usersController.CheckUserIsExist(dstUserID)
			if secCheckErr == nil {
				if isSecUserExist {
					result = true
				} else {
					err = users_controller.UserNotExistErr
				}
			}
		} else {
			err = users_controller.UserNotExistErr
		}
	}
	return result, err
}

func (m *UsersManager) checkUserCanBuyService(userID int, value float64) (bool, error) {
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

func (m *UsersManager) makeReportsForAllUsers(srcUserID, dstUserID int, value float64, comment string) error {
	err := m.transactionsController.AddNewRecordTransferTo(srcUserID, dstUserID, value, comment)
	if err == nil {
		err = m.transactionsController.AddNewRecordTransferFrom(dstUserID, srcUserID, value, comment)
	}
	return err
}
