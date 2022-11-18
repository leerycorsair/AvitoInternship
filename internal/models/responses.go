package models

import (
	"encoding/json"
	"net/http"
)

type ShortResponseMessage struct {
	Comments string `json:"comments"`
}

type BalanceResponseMessage struct {
	Balance  float64 `json:"balance"`
	Comments string  `json:"comments"`
}

type FinanceReportResponseMessage struct {
	FileURL string `json:"fileURL"`
}

type UserReportResponseMessage struct {
	AllTransactions []Transaction `json:"transactions"`
}

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
		SendShortResponse(w, http.StatusInternalServerError, "internal server problems")
	}
}

func FinanceReportResponse(w http.ResponseWriter, fileURL string) {
	var msg = FinanceReportResponseMessage{fileURL}
	result, err := json.Marshal(msg)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(result)
	} else {
		SendShortResponse(w, http.StatusInternalServerError, "internal server problems")
	}
}

func UserReportResponse(w http.ResponseWriter, allTransactions []Transaction) {
	var msg = UserReportResponseMessage{allTransactions}
	result, err := json.Marshal(msg)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(result)
	} else {
		SendShortResponse(w, http.StatusInternalServerError, "internal server problems")
	}
}
