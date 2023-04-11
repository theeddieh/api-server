package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func formatResponseBody(w http.ResponseWriter, _ *http.Request, statusCode int, b interface{}) {

	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(b)
	if err != nil {
		fmt.Println("unable to encode json:", b)
	}
}

func sendErrorResponse(w http.ResponseWriter, _ *http.Request, e error) {

	w.WriteHeader(http.StatusInternalServerError)

	body := make(map[string]string, 1)
	body["error"] = fmt.Sprint(e)

	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		fmt.Println(err)
	}
}
