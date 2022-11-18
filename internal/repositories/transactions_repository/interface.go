package transactions_repository

import "AvitoInternship/internal/models"

type TransactionsRepoInterface interface {
	AddNewTransaction(newTransaction models.Transaction) error
	GetAllTransactions() ([]models.Transaction, error)
	GetUserTransactions(userID int, orderBy string, limit int, offset int) ([]models.Transaction, error)
	GetTransactionByID(transactionID int) (models.Transaction, error)
}
