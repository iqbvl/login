package middleware

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

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

func ValidateBasicAuthHeader(r *http.Request) (*models.User, error) {
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		return nil, errors.New("Error Credentials")
	}
	payload, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		return nil, errors.New("Error Credentials")
	}
	userCredentials := models.User{
		Username: pair[0],
		Password: pair[1],
	}
	return &userCredentials, nil
}
