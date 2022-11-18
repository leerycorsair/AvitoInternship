package reports_manager

import "AvitoInternship/internal/models"

type ReportsManagerInterface interface {
	GetFinanceReport(month, year int, url string) error
	GetUserReport(userID int, orderBy string, limit, offset int) ([]models.Transaction, error)
}
