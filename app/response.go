package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func formatResponseBody(w http.ResponseWriter, _ *http.Request, d interface{}, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(d)
	if err != nil {
		fmt.Println("unable to encode json:", d)
	}
}
