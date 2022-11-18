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

func (c *UsersController) CheckUserIsExist(userID int) (result bool, err error) {
	_, isUserExistErr := c.repo.GetCurrentBalance(userID)
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

func (c *UsersController) CreateNewUser(userID int) error {
	c.mutex.Lock()
	isUserExist, err := c.CheckUserIsExist(userID)
	if err == nil {
		if !isUserExist {
			err = c.repo.AddNewUser(userID)
		} else {
			err = UserIsExistErr
		}
	}
	c.mutex.Unlock()
	return err
}

func (c *UsersController) CheckBalance(userID int) (float64, error) {
	return c.repo.GetCurrentBalance(userID)
}

func (c *UsersController) CheckAbleToBuyService(userID int, servicePrice float64) (bool, error) {
	var result bool
	balance, err := c.repo.GetCurrentBalance(userID)
	if err == nil {
		if servicePrice <= balance {
			result = true
		}
	}
	return result, err
}

func (c *UsersController) DonateMoney(userID int, value float64) (err error) {
	c.mutex.Lock()
	if value >= 0 {
		err = c.repo.ChangeBalance(userID, value)
	} else {
		err = NegValueError
	}
	c.mutex.Unlock()
	return err
}

func (c *UsersController) SpendMoney(userID int, value float64) error {
	c.mutex.Lock()
	canSpendMoney, err := c.CheckAbleToBuyService(userID, value)
	if err == nil && canSpendMoney {
		err = c.repo.ChangeBalance(userID, -value)
	} else {
		err = NotEnoughMoneyErr
	}
	c.mutex.Unlock()
	return err
}
