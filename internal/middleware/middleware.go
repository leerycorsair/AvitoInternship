package middleware

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func Log(logger *logrus.Entry, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("New Request: Method - %s, URL - %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func Panic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Recovered", err)
				http.Error(w, "Internal Server Error", 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
