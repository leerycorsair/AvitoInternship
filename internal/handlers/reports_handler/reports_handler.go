package reports_handler

import (
	"fmt"
	"net/http"
	"strconv"

	"AvitoInternship/internal/controllers/orders_controller"
	"AvitoInternship/internal/controllers/users_controller"

	"AvitoInternship/internal/models"
	"AvitoInternship/internal/tools"

	"github.com/sirupsen/logrus"
)

type ReportsHandler struct {
	l *logrus.Entry
	m reports_manager.ReportManagerInterface
}

func CreateReportsHandler(m reports_manager.ReportManagerInterface) *ReportsHandler {
	contextLogger := logrus.WithFields(logrus.Fields{})
	logrus.SetReportCaller(false)
	logrus.SetFormatter(&logrus.TextFormatter{PadLevelText: false, DisableLevelTruncation: false})
	return &ReportsHandler{m: m, l: contextLogger}
}

func (h *ReportsHandler) GetFinanceReport(w http.ResponseWriter, r *http.Request) {
	var statusCode int
	var handleMessage string

	monthFromQuery := r.URL.Query().Get("month")

	if monthFromQuery == "" {
		models.SendShortResponse(w, http.StatusBadRequest, "Month Not Found")
		return
	}

	yearFromQuery := r.URL.Query().Get("year")

	if yearFromQuery == "" {
		models.SendShortResponse(w, http.StatusBadRequest, "Year Not Found")
		return
	}

	month, castMonthErr := strconv.Atoi(monthFromQuery)
	if castMonthErr != nil {
		models.SendShortResponse(w, http.StatusBadRequest, "Incorrect Month")
		return
	}

	year, castYearErr := strconv.Atoi(yearFromQuery)
	if castYearErr != nil {
		models.SendShortResponse(w, http.StatusBadRequest, "Incorrect Year")
		return
	}

	fileURL := "report.csv"

	getFinanceReportErr := h.m.GetFinanceReport(month, year, fileURL)
	switch getFinanceReportErr {
	case nil:
		models.FinanceReportResponse(w, "report.csv")
		return
	case orders_controller.BadYearError:
		statusCode = http.StatusBadRequest
		handleMessage = fmt.Sprintf("%v", orders_controller.BadYearError)
	case orders_controller.BadMonthError:
		statusCode = http.StatusBadRequest
		handleMessage = fmt.Sprintf("%v", orders_controller.BadMonthError)
	default:
		statusCode = http.StatusInternalServerError
		handleMessage = fmt.Sprintf("Internal Server Error")
	}
	models.SendShortResponse(w, statusCode, handleMessage)
	h.l.Infof("Request: Method - %s,  URL - %s, Result: Status_Code = %d, Text = %s, Err = %v",
		r.Method, r.URL.Path, statusCode, handleMessage, getFinanceReportErr)
}

func (h *ReportsHandler) GetUserReport(w http.ResponseWriter, r *http.Request) {
	var statusCode int
	var handleMessage string

	userIDFromQuery := r.URL.Query().Get("userID")

	if userIDFromQuery == "" {
		models.SendShortResponse(w, http.StatusBadRequest, "UserID Not Found")
		return
	}

	userID, err := strconv.Atoi(userIDFromQuery)
	if err != nil {
		models.SendShortResponse(w, http.StatusBadRequest, "UserID Isn't Number")
		return
	}

	orderBy := r.URL.Query().Get("orderBy")

	switch orderBy {
	case "id":
	case "time":
	case "value":
	default:
		orderBy = "id"
	}

	limit, getLimitErr := tools.GetOptionalIntParam(r, "Limit")
	if getLimitErr != nil {
		models.SendShortResponse(w, http.StatusBadRequest, "Limit Isn't Number")
		return
	}

	offset, getOffsetErr := tools.GetOptionalIntParam(r, "Offset")
	if getOffsetErr != nil {
		models.SendShortResponse(w, http.StatusBadRequest, "Offset Isn't Number")
		return
	}

	userReport, getReportErr := h.m.GetUserReport(userID, orderBy, limit, offset)
	switch getReportErr {
	case nil:
		models.UserReportResponse(w, userReport)
		return
	case users_controller.UserNotExistErr:
		statusCode = http.StatusUnauthorized
		handleMessage = fmt.Sprintf("%v", users_controller.UserNotExistErr)
	default:
		statusCode = http.StatusInternalServerError
		handleMessage = fmt.Sprintf("Internal Server Error")
	}
	models.SendShortResponse(w, statusCode, handleMessage)
	h.l.Infof("Request: Method - %s,  URL - %s, Result: Status_Code = %d, Text = %s, Error = %v",
		r.Method, r.URL.Path, statusCode, handleMessage, getReportErr)
}
