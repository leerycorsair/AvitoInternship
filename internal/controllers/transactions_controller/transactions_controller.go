package transactions_controller

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"AvitoInternship/internal/models"
	"AvitoInternship/internal/repositories/transactions_repository"
)

type TransactionsController struct {
	mutex sync.RWMutex
	repo  transactions_repository.TransactionsRepoInterface
	cnt   int64
}

func CreateTransactionsController(repo transactions_repository.TransactionsRepoInterface) *TransactionsController {
	return &TransactionsController{
		mutex: sync.RWMutex{},
		repo:  repo,
	}
}

func (c *TransactionsController) GetTransactionByID(transactionID int) (models.Transaction, error) {
	return c.repo.GetTransactionByID(transactionID)
}

func (c *TransactionsController) AddNewRecordAddBalance(userID int, value float64, comments string) error {
	newTransact := models.Transaction{
		TransactionId:   int(c.cnt),
		UserId:          userID,
		TransactionType: models.Add,
		Value:           value,
		CreatedAt:       time.Now(),
		ActionComments:  "Balance was increased",
		AddComments:     comments,
	}
	c.mutex.Lock()
	err := c.repo.AddNewTransaction(newTransact)
	c.mutex.Unlock()
	return err
}

func (c *TransactionsController) AddNewRecordReserveService(userID int, value float64, serviceID int, comments string) error {
	newTransact := models.Transaction{
		TransactionId:   int(c.cnt),
		UserId:          userID,
		TransactionType: models.Reserve,
		Value:           value,
		CreatedAt:       time.Now(),
		ActionComments:  "Service " + strconv.Itoa(serviceID) + " was payed",
		AddComments:     comments,
	}

	c.mutex.Lock()
	err := c.repo.AddNewTransaction(newTransact)
	atomic.AddInt64(&c.cnt, 1)
	c.mutex.Unlock()

	return err
}

func (c *TransactionsController) AddNewRecordReturnService(userID int, value float64, serviceID int, comments string) error {
	newTransact := models.Transaction{
		TransactionId:   int(c.cnt),
		UserId:          userID,
		TransactionType: models.Return,
		Value:           value,
		CreatedAt:       time.Now(),
		ActionComments:  "Return for Service: " + strconv.Itoa(serviceID),
		AddComments:     comments,
	}

	c.mutex.Lock()
	err := c.repo.AddNewTransaction(newTransact)
	atomic.AddInt64(&c.cnt, 1)
	c.mutex.Unlock()
	return err
}

func (c *TransactionsController) AddNewRecordTransferTo(srcUserID, dstUserID int, value float64, comments string) error {
	newTransact := models.Transaction{
		TransactionId:   int(c.cnt),
		UserId:          srcUserID,
		TransactionType: models.Transfer,
		Value:           value,
		CreatedAt:       time.Now(),
		ActionComments:  "Transaction to User: " + fmt.Sprintf("%d", dstUserID),
		AddComments:     comments,
	}
	c.mutex.Lock()
	err := c.repo.AddNewTransaction(newTransact)
	atomic.AddInt64(&c.cnt, 1)
	c.mutex.Unlock()
	return err
}

func (c *TransactionsController) AddNewRecordTransferFrom(srcUserID, dstUserID int, value float64, comments string) error {
	newTransact := models.Transaction{
		TransactionId:   int(c.cnt),
		UserId:          srcUserID,
		TransactionType: models.Transfer,
		Value:           value,
		CreatedAt:       time.Now(),
		ActionComments:  "Transaction from User: " + fmt.Sprintf("%d", dstUserID),
		AddComments:     comments,
	}

	c.mutex.Lock()
	err := c.repo.AddNewTransaction(newTransact)
	atomic.AddInt64(&c.cnt, 1)
	c.mutex.Unlock()
	return err
}

func (c *TransactionsController) GetUserTransactions(userID int, orderBy string,
	limit, offset int) ([]models.Transaction, error) {
	return c.repo.GetUserTransactions(userID, orderBy, limit, offset)
}
