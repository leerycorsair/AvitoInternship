package services_handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"AvitoInternship/internal/controllers/users_controller"
	"AvitoInternship/internal/models"

	"github.com/sirupsen/logrus"
)

type ServicesHandler struct {
	l *logrus.Entry
	m orders_manager.OrdersManagerInterface
}

func CreateServicesHandler(m orders_manager.OrdersManagerInterface) *ServicesHandler {
	l := logrus.WithFields(logrus.Fields{})
	logrus.SetReportCaller(false)
	logrus.SetFormatter(&logrus.TextFormatter{PadLevelText: false, DisableLevelTruncation: false})
	return &ServicesHandler{l: l, m: m}
}

func (h *ServicesHandler) ReserveService(w http.ResponseWriter, r *http.Request) {
	var reserveParams models.ReserveServiceMessage
	var statusCode int
	var handleMessage string

	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	unmarshalError := json.Unmarshal(body, &reserveParams)
	if unmarshalError != nil {
		http.Error(w, "Unmarshal Error", http.StatusBadRequest)
		return
	}

	reserveError := h.m.ReserveService(reserveParams.UserID, reserveParams.OrderID, reserveParams.ServiceID,
		reserveParams.Value, reserveParams.Comments)

	switch reserveError {
	case nil:
		statusCode = http.StatusOK
		handleMessage = "OK"
	case users_controller.UserNotExistErr:
		statusCode = http.StatusUnauthorized
		handleMessage = fmt.Sprintf("%v", users_controller.UserNotExistErr)
	case users_controller.NotEnoughMoneyErr:
		statusCode = http.StatusUnprocessableEntity
		handleMessage = fmt.Sprintf("%v", users_controller.NotEnoughMoneyErr)
	case users_controller.NegValueError:
		statusCode = http.StatusUnprocessableEntity
		handleMessage = fmt.Sprintf("%v", users_controller.NegValueError)
	default:
		statusCode = http.StatusInternalServerError
		handleMessage = fmt.Sprintf("Internal Server Error")
	}
	models.SendShortResponse(w, statusCode, handleMessage)
	h.l.Infof("Request: Method - %s,  URL - %s, Result: Status_Code = %d, Text = %s, Err = %v",
		r.Method, r.URL.Path, statusCode, handleMessage, reserveError)
}

func (h *ServicesHandler) AcceptService(w http.ResponseWriter, r *http.Request) {
	var acceptParams models.AcceptServiceMessage
	var statusCode int
	var handleMessage string

	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	unmarshalError := json.Unmarshal(body, &acceptParams)
	if unmarshalError != nil {
		http.Error(w, "Invalid Body Params", http.StatusBadRequest)
		return
	}

	acceptReserveErr := h.m.AcceptReserve(acceptParams.UserID, acceptParams.OrderID, acceptParams.ServiceID)

	switch acceptReserveErr {
	case nil:
		statusCode = http.StatusOK
		handleMessage = "OK"
	case users_controller.UserNotExistErr:
		statusCode = http.StatusUnauthorized
		handleMessage = fmt.Sprintf("%v", users_controller.UserNotExistErr)
	case orders_controller.OrderNotFound:
		statusCode = http.StatusNotFound
		handleMessage = fmt.Sprintf("%v", orders_controller.OrderNotFound)
	case orders_controller.WrongStateError:
		statusCode = http.StatusForbidden
		handleMessage = fmt.Sprintf("%v", orders_controller.WrongStateError)
	default:
		statusCode = http.StatusInternalServerError
		handleMessage = fmt.Sprintf("Internal Server Error")
	}
	models.SendShortResponse(w, statusCode, handleMessage)
	h.l.Infof("Request: Method - %s,  URL - %s, Result: Status_Code = %d, Text = %s, Err = %v",
		r.Method, r.URL.Path, statusCode, handleMessage, acceptReserveErr)
}

func (h *ServicesHandler) CancelService(w http.ResponseWriter, r *http.Request) {
	var cancelParams models.CancelServiceMessage
	var statusCode int
	var handleMessage string

	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	unmarshalError := json.Unmarshal(body, &cancelParams)
	if unmarshalError != nil {
		http.Error(w, "Invalid Body Params", http.StatusBadRequest)
		return
	}

	cancelReserveErr := h.m.CancelReserve(cancelParams.UserID, cancelParams.OrderID, cancelParams.ServiceID,
		cancelParams.Comments)

	switch cancelReserveErr {
	case nil:
		statusCode = http.StatusOK
		handleMessage = "OK"
	case users_controller.UserNotExistErr:
		statusCode = http.StatusUnauthorized
		handleMessage = fmt.Sprintf("%v", users_controller.UserNotExistErr)
	case orders_controller.OrderNotFound:
		statusCode = http.StatusNotFound
		handleMessage = fmt.Sprintf("%v", orders_controller.OrderNotFound)
	case orders_controller.WrongStateError:
		statusCode = http.StatusForbidden
		handleMessage = fmt.Sprintf("%v", orders_controller.WrongStateError)
	default:
		statusCode = http.StatusInternalServerError
		handleMessage = fmt.Sprintf("Internal Server Error")
	}
	models.SendShortResponse(w, statusCode, handleMessage)
	h.l.Infof("Request: Method - %s,  URL - %s, Result: Status_Code = %d, Text = %s, Err = %v",
		r.Method, r.URL.Path, statusCode, handleMessage, cancelReserveErr)
}
