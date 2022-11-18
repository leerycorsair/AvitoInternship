package users_repository

import (
	"database/sql"
	"sync"
)

type UsersRepo struct {
	mutex sync.RWMutex
	c     *sql.DB
}

func CreateUsersRepo(c *sql.DB) *UsersRepo {
	return &UsersRepo{c: c}
}

func (repo *UsersRepo) AddNewUser(userID int) error {
	repo.mutex.Lock()
	_, err := repo.c.Exec("insert into users(`id`, `balance`) values (?, 0);", userID)
	repo.mutex.Unlock()
	return err
}

func (repo *UsersRepo) GetCurrentBalance(userID int) (balance float64, err error) {
	repo.mutex.Lock()

	row := repo.c.QueryRow("select balance from users where id = ?;", userID)
	err = row.Scan(&balance)
	repo.mutex.Unlock()

	if err == sql.ErrNoRows {
		err = UserNotExist
	}

	return balance, err
}

func (repo *UsersRepo) ChangeBalance(userID int, delta float64) error {
	repo.mutex.Lock()

	_, err := repo.c.Exec("update users set balance = balance + ? where id = ?;", delta, userID)
	repo.mutex.Unlock()

	if err == sql.ErrTxDone {
		err = UserNotExist
	}

	return err
}
