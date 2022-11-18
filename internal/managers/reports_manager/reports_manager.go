package reports_manager

import (
	"AvitoInternship/internal/controllers/orders_controller"
	"AvitoInternship/internal/controllers/reports_controller"
	"AvitoInternship/internal/controllers/transactions_controller"

	"AvitoInternship/internal/controllers/users_controller"
	"AvitoInternship/internal/models"
)

type ReportsManager struct {
	usersController        users_controller.UsersControllerInterface
	ordersController       orders_controller.OrdersControllerInterface
	transactionsController transactions_controller.TransactionsControllerInterface
}

func CreateReportsManager(uc users_controller.UsersControllerInterface, oc orders_controller.OrdersControllerInterface,
	tc transactions_controller.TransactionsControllerInterface) *ReportsManager {
	return &ReportsManager{
		usersController:        uc,
		ordersController:       oc,
		transactionsController: tc,
	}
}

func (m *ReportsManager) GetFinanceReport(month, year int, url string) error {
	dataToReport, err := m.ordersController.GetFinanceReports(month, year)
	if err == nil {
		reportController := reports_controller.CreateNewReportsController()
		err = reportController.CreateFinancialReportCSV(dataToReport, url)
	}
	return err
}

func (m *ReportsManager) GetUserReport(userID int, orderBy string, limit, offset int) ([]models.Transaction, error) {
	var allTransactions = make([]models.Transaction, 0)
	var err error

	if limit == -1 {
		limit = 25
	}
	if offset == -1 {
		offset = 0
	}
	if orderBy == "" {
		orderBy = "id"
	}

	isAccExist, checkAccountError := m.usersController.CheckUserIsExist(userID)
	if checkAccountError == nil {
		if isAccExist {
			allTransactions, err = m.transactionsController.GetUserTransactions(userID, orderBy, limit, offset)
		} else {
			err = users_controller.UserNotExistErr
		}
	} else {
		err = checkAccountError
	}

	return allTransactions, err
}
