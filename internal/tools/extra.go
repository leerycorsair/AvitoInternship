package tools

import (
	"net/http"
	"strconv"
)

func GetOptionalIntParam(r *http.Request, paramName string) (result int, err error) {
	strParam := r.URL.Query().Get(paramName)
	if strParam == "" {
		result = -1
	} else {
		result, err = strconv.Atoi(strParam)
	}
	return result, err
}
