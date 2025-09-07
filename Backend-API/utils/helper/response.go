package helper

import (
	"encoding/json"
	"farm-integrated-web3/dto"
	"net/http"
)

func HttpWriter(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, err.Error(), status)
		}
	}
}

func HttpError(w http.ResponseWriter, status int, message string) {
	HttpWriter(w, status, dto.ResponseError{
		Error: message,
	})
}
