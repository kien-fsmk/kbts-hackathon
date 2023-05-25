package handlers

import (
	"encoding/json"
	"net/http"
)

func returnResponse(rw http.ResponseWriter, responseBody interface{}, responseCode int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(responseCode)
	responseJSON, _ := json.Marshal(responseBody)
	_, _ = rw.Write(responseJSON)
}
