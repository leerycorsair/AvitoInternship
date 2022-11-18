package transactions_controller

import "AvitoInternship/internal/models"

type TransactionsControllerInterface interface {
	GetTransactionByID(transactionID int) (models.Transaction, error)
	AddNewRecordAddBalance(userID int, value float64, comments string) error
	AddNewRecordReserveService(userID int, value float64, serviceID int, comments string) error
	AddNewRecordReturnService(userID int, value float64, serviceID int, comments string) error
	AddNewRecordTransferTo(srcUserID, dstUserID int, value float64, comments string) error
	AddNewRecordTransferFrom(srcUserID, dstUserID int, value float64, comments string) error
	GetUserTransactions(userID int, orderBy string, limit, offset int) ([]models.Transaction, error)
}
