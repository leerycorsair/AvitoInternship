package models

import (
	"encoding/json"
	"net/http"
)

func SendShortResponse(w http.ResponseWriter, code int, comments string) {
	var msg = ShortResponseMessage{comments}
	result, err := json.Marshal(msg)
	if err == nil {
		w.WriteHeader(code)
		_, err = w.Write(result)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func BalanceResponse(w http.ResponseWriter, balance float64, comments string) {
	var msg = BalanceResponseMessage{balance, comments}
	result, err := json.Marshal(msg)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(result)
	} else {
		SendShortResponse(w, http.StatusInternalServerError, "Internal Server Problems")
	}
}

func FinanceReportResponse(w http.ResponseWriter, fileURL string) {
	var msg = FinanceReportResponseMessage{fileURL}
	result, err := json.Marshal(msg)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(result)
	} else {
		SendShortResponse(w, http.StatusInternalServerError, "Internal Server Problems")
	}
}

func UserReportResponse(w http.ResponseWriter, allTransactions []Transaction) {
	var msg = UserReportResponseMessage{allTransactions}
	result, err := json.Marshal(msg)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(result)
	} else {
		SendShortResponse(w, http.StatusInternalServerError, "Internal Server Problems")
	}
}
