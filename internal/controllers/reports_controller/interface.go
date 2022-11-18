package reports_controller

import "AvitoInternship/internal/models"

type ReportControllerInterface interface {
	CreateFinancialReportCSV([]models.Report, string) error
}
