package transactions_repository

import (
	"AvitoInternship/internal/models"
	"database/sql"
	"sync"
	"time"
)

type TransactionsRepo struct {
	mutex sync.RWMutex
	c     *sql.DB
}

func CreateTransactionsRepo(c *sql.DB) *TransactionsRepo {
	return &TransactionsRepo{c: c}
}

func (repo *TransactionsRepo) AddNewTransaction(newTransaction models.Transaction) error {
	repo.mutex.Lock()
	curTime := newTransaction.CreatedAt.Format("2006-01-02 15:04:05")
	_, err := repo.c.Exec("insert into transactions(`transaction_id`, `user_id`, `transaction_type`, `value`, "+
		"`created_at`, `action_comments`, `add_comments`) values (?, ?, ?, ?, ?, ?, ?);", newTransaction.TransactionId, newTransaction.UserId,
		newTransaction.TransactionType, newTransaction.Value, curTime,
		newTransaction.ActionComments, newTransaction.AddComments)
	repo.mutex.Unlock()
	return err
}

func (repo *TransactionsRepo) GetAllTransactions() ([]models.Transaction, error) {
	allTransactions := make([]models.Transaction, 0)

	repo.mutex.Lock()
	rows, err := repo.c.Query("select transaction_id, user_id, transaction_type, value, created_at," +
		" action_comments, add_comments from transactions")
	repo.mutex.Unlock()

	if err == nil {
		for rows.Next() {
			newTransact := models.Transaction{}
			var transactTime string
			err = rows.Scan(&newTransact.TransactionId, &newTransact.UserId, &newTransact.TransactionType,
				&newTransact.Value, &transactTime, &newTransact.ActionComments, &newTransact.AddComments)
			newTransact.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", transactTime)
			if err != nil {
				break
			} else {
				allTransactions = append(allTransactions, newTransact)
			}
		}
	}
	return allTransactions, err
}

func (repo *TransactionsRepo) GetUserTransactions(userID int, orderBy string, limit, offset int) ([]models.Transaction, error) {
	allTransactions := make([]models.Transaction, 0)

	repo.mutex.Lock()
	rows, err := repo.c.Query("select transaction_id, user_id, transaction_type, value, created_at,"+
		" action_comments, add_comments from transactions where user_id = ? order by ? desc limit ? offset ?", userID, orderBy, limit, offset)
	repo.mutex.Unlock()

	if err == nil {
		for rows.Next() {
			newTransact := models.Transaction{}
			var transactTime string
			err = rows.Scan(&newTransact.TransactionId, &newTransact.UserId, &newTransact.TransactionType,
				&newTransact.Value, &transactTime, &newTransact.ActionComments, &newTransact.AddComments)
			newTransact.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", transactTime)
			if err != nil {
				break
			} else {
				allTransactions = append(allTransactions, newTransact)
			}
		}
	}
	return allTransactions, err
}

func (repo *TransactionsRepo) GetTransactionByID(transactionID int) (models.Transaction, error) {
	curTransact := models.Transaction{}

	repo.mutex.Lock()
	row := repo.c.QueryRow("select transaction_id, user_id, transaction_type, value, created_at,"+
		" action_comments, add_comments from transactions where transaction_id = ?", transactionID)
	repo.mutex.Unlock()

	var transactTime string
	err := row.Scan(&curTransact.TransactionId, &curTransact.UserId, &curTransact.TransactionType,
		&curTransact.Value, &transactTime, &curTransact.ActionComments, &curTransact.AddComments)
	curTransact.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", transactTime)

	return curTransact, err
}
