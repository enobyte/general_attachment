package utils

import (
	"encoding/json"
	"net/http"
)

func Message(code int, status bool, message string) map[string]interface{} {
	return map[string]interface{}{"code": code, "status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func BadReq() int {
	return 400
}

func SuccesReq() int {
	return 200
}
