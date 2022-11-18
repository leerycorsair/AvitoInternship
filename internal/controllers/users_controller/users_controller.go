package users_controller

import (
	"AvitoInternship/internal/repositories/users_repository"
	"sync"
)

type UsersController struct {
	mutex sync.RWMutex
	repo  users_repository.UsersRepoInterface
}

func CreateUsersController(repo users_repository.UsersRepoInterface) *UsersController {
	return &UsersController{mutex: sync.RWMutex{}, repo: repo}
}

func (m *UsersController) CheckUserIsExist(userID int) (result bool, err error) {
	_, isUserExistErr := m.repo.GetCurrentBalance(userID)
	switch isUserExistErr {
	case users_repository.UserNotExist:
		result = false
	case nil:
		result = true
	default:
		err = isUserExistErr
	}
	return result, err
}

func (m *UsersController) CreateNewUser(userID int) error {
	m.mutex.Lock()
	isUserExist, err := m.CheckUserIsExist(userID)
	if err == nil {
		if !isUserExist {
			err = m.repo.AddNewUser(userID)
		} else {
			err = UserIsExistErr
		}
	}
	m.mutex.Unlock()
	return err
}

func (m *UsersController) CheckBalance(userID int) (float64, error) {
	return m.repo.GetCurrentBalance(userID)
}

func (m *UsersController) CheckAbleToBuyService(userID int, servicePrice float64) (bool, error) {
	var result bool
	balance, err := m.repo.GetCurrentBalance(userID)
	if err == nil {
		if servicePrice <= balance {
			result = true
		}
	}
	return result, err
}

func (m *UsersController) DonateMoney(userID int, value float64) (err error) {
	m.mutex.Lock()
	if value >= 0 {
		err = m.repo.ChangeBalance(userID, value)
	} else {
		err = NegValueError
	}
	m.mutex.Unlock()
	return err
}

func (m *UsersController) SpendMoney(userID int, value float64) error {
	m.mutex.Lock()
	canSpendMoney, err := m.CheckAbleToBuyService(userID, value)
	if err == nil && canSpendMoney {
		err = m.repo.ChangeBalance(userID, -value)
	} else {
		err = NotEnoughMoneyErr
	}
	m.mutex.Unlock()
	return err
}
