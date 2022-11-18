package users_handler

import "net/http"

type UsersHandlerInterface interface {
	GetBalance(w http.ResponseWriter, r *http.Request)
	AddBalance(w http.ResponseWriter, r *http.Request)
	Transfer(w http.ResponseWriter, r *http.Request)
}
