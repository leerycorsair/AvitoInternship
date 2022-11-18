package reports_handler

import "net/http"

type ReportsHandlerInterface interface {
	GetFinanceReport(w http.ResponseWriter, r *http.Request)
	GetUserReport(w http.ResponseWriter, r *http.Request)
}
