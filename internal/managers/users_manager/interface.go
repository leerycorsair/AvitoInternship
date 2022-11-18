package users_manager

type UsersManagerInterface interface {
	AddBalance(userID int, value float64, comments string) error
	GetUserBalance(userID int) (float64, error)
	Transfer(srcUserID, dstUserID int, value float64, comments string) error
}
