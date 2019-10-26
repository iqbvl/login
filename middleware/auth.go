package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/iqbvl/login/models"
)

func RequestBodyConverter(req *http.Request) (*models.User, error) {
	decoder := json.NewDecoder(req.Body)
	var t models.User
	err := decoder.Decode(&t)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return &t, nil
}

func OTPRequestBodyConverter(req *http.Request) (*models.OTP, error) {
	decoder := json.NewDecoder(req.Body)
	var t models.OTP
	err := decoder.Decode(&t)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return &t, nil
}

func WriteResponse(w http.ResponseWriter, response interface{}, code int) {
	jsonResponse, _ := json.Marshal(response)
	w.WriteHeader(code)
	w.Write(jsonResponse)
}
