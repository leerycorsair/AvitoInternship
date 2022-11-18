package services_handler

import "net/http"

type ServicesHandlerInterface interface {
	ReserveService(w http.ResponseWriter, r *http.Request)
	AcceptService(w http.ResponseWriter, r *http.Request)
	CancelService(w http.ResponseWriter, r *http.Request)
}
