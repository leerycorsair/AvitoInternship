package users_controller

import "errors"

var (
	NegValueError     = errors.New("Bad Value To Update")
	NotEnoughMoneyErr = errors.New("Not Enough Money")
	UserIsExistErr    = errors.New("User Us Already Exist")
	UserNotExistErr   = errors.New("User Is Not Exist")
)

type UsersControllerInterface interface {
	CheckUserIsExist(userID int) (result bool, err error)
	CreateNewUser(userID int) error
	CheckBalance(userID int) (float64, error)
	CheckAbleToBuyService(userID int, servicePrice float64) (bool, error)
	DonateMoney(userID int, value float64) (err error)
	SpendMoney(userID int, value float64) error
}
