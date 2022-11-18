package users_repository

import "errors"

var UserNotExist = errors.New("User Not Exist")

type UsersRepoInterface interface {
	AddNewUser(userID int) error
	GetCurrentBalance(userID int) (balance float64, err error)
	ChangeBalance(userID int, delta float64) error
}
