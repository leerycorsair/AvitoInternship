package users_handler

import (
	"AvitoInternship/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

type UsersHandler struct {
	l *logrus.Entry
	m users_manager.UsersManagerInterface
}

func CreateUsersHandler(newManager users_manager.UsersManagerInterface) *UsersHandler {
	contextLogger := logrus.WithFields(logrus.Fields{})
	logrus.SetReportCaller(false)
	logrus.SetFormatter(&logrus.TextFormatter{PadLevelText: false, DisableLevelTruncation: false})
	return &UsersHandler{l: contextLogger, m: newManager}
}

func (h *UsersHandler) AddBalance(w http.ResponseWriter, r *http.Request) {
	var statusCode int
	var handleMessage string

	var addParams models.AddMessage

	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	unmarshalError := json.Unmarshal(body, &addParams)
	if unmarshalError != nil {
		http.Error(w, "Invalid Body Params", http.StatusBadRequest)
		return
	}

	addError := h.m.AddBalance(addParams.UserID, addParams.Value, addParams.Comments)

	switch addError {
	case nil:
		statusCode = http.StatusOK
		handleMessage = "OK"
	case users_controller.NegValueError:
		statusCode = http.StatusUnprocessableEntity
		handleMessage = fmt.Sprintf("%v", users_controller.NegValueError)
	default:
		statusCode = http.StatusInternalServerError
		handleMessage = fmt.Sprintf("Internal Server Error")
	}
	models.SendShortResponse(w, statusCode, handleMessage)
	h.l.Infof("Request: Method - %s,  UDL - %s, Result: Status_Code = %d, Text = %s, Err = %v",
		r.Method, r.URL.Path, statusCode, handleMessage, addError)
}

func (h *UsersHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	var statusCode int
	var handleMessage string

	var transferParams models.TransferMessage

	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		http.Error(w, "Server Problems", http.StatusInternalServerError)
		return
	}

	unmarshalError := json.Unmarshal(body, &transferParams)
	if unmarshalError != nil {
		http.Error(w, "Invalid Body Params", http.StatusBadRequest)
		return
	}

	transferError := h.m.Transfer(transferParams.SrcUserID, transferParams.DstUserID,
		transferParams.Value, transferParams.Comments)

	switch transferError {
	case nil:
		statusCode = http.StatusOK
		handleMessage = "OK"
	case users_controller.UserNotExistErr:
		statusCode = http.StatusUnauthorized
		handleMessage = fmt.Sprintf("%v", users_controller.UserNotExistErr)
	case users_controller.NotEnoughMoneyErr:
		statusCode = http.StatusUnprocessableEntity
		handleMessage = fmt.Sprintf("%v", users_controller.NotEnoughMoneyErr)
	default:
		statusCode = http.StatusInternalServerError
		handleMessage = fmt.Sprintf("Internal Server Error: %v", transferError)
	}
	models.SendShortResponse(w, statusCode, handleMessage)
	h.l.Infof("Request: Method - %s,  URL - %s, Result: Status_Code = %d, Text = %s",
		r.Method, r.URL.Path, statusCode, handleMessage)
}

func (h *UsersHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	var statusCode int
	var handleMessage string

	userIDFromQuery := r.URL.Query().Get("userID")

	if userIDFromQuery == "" {
		models.SendShortResponse(w, http.StatusBadRequest, "UserID Wasn't Found")
		return
	}

	userID, err := strconv.Atoi(userIDFromQuery)
	if err != nil {
		models.SendShortResponse(w, http.StatusBadRequest, "UserID Isn't A Number")
		return
	}

	balance, getBalanceErr := h.m.GetUserBalance(int(userID))
	switch getBalanceErr {
	case nil:
		models.BalanceResponse(w, balance, "OK")
		return
	case users_controller.UserNotExistErr:
		statusCode = http.StatusUnauthorized
		handleMessage = fmt.Sprintf("%v", users_controller.UserNotExistErr)
	default:
		statusCode = http.StatusInternalServerError
		handleMessage = fmt.Sprintf("Internal Server Error")
	}
	models.SendShortResponse(w, statusCode, handleMessage)
	h.l.Infof("Request: Method - %s,  URL - %s, Result: Status_Code = %d, Text = %s",
		r.Method, r.URL.Path, statusCode, handleMessage)
}
