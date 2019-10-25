package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/iqbvl/login/helper"
	"github.com/iqbvl/login/middleware"
	"github.com/iqbvl/login/models"
)

const (
	invalidEmailFormat      = "You are inputting wrong email address format"
	emptyLoginField         = "Username and Password cant be empty"
	postMethodSupported     = "Only Post Allowed"
	errorConvertRequestBody = "Error when converting request body"
	tokenError              = "Error Generate Token"
)

var TokenAuth *jwtauth.JWTAuth

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	loginResponse := models.LoginResponse{}
	rsp := models.Response{}
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		responses := &models.Response{
			Data:          "",
			ErrorMessages: postMethodSupported,
			IsSuccess:     false,
		}
		middleware.WriteResponse(w, responses, http.StatusMethodNotAllowed)
		return
	}
	var user *models.User // defaul
	user, err := middleware.RequestBodyConverter(r)
	if err != nil {
		responses := &models.Response{
			Data:          "",
			ErrorMessages: errorConvertRequestBody,
			IsSuccess:     false,
		}
		middleware.WriteResponse(w, responses, http.StatusOK)
		return
	}

	if user.Username == "" || user.Password == "" {
		responses := &models.Response{
			Data:          "",
			ErrorMessages: emptyLoginField,
			IsSuccess:     false,
		}
		middleware.WriteResponse(w, responses, http.StatusOK)
		return
	}

	_, tokenString, _ := TokenAuth.Encode(jwt.MapClaims{
		"id":        1,
		"username":  user.Username,
		"claimDate": time.Now().Format(`2006-01-02T15:04:05.000-07:00`),
		"expire":    helper.ParseDuration("P1Y"),
	})
	loginResponse.Token = tokenString

	ctx := context.WithValue(r.Context(), "TokenAuth", TokenAuth)
	r.WithContext(ctx)

	//Success
	rsp.Data = tokenString
	rsp.ErrorMessages = ""
	rsp.IsSuccess = true

	response, _ := json.Marshal(rsp)
	w.Write(response)

}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var responses models.Response
	var user *models.User // defaul
	user, err := middleware.RequestBodyConverter(r)
	if err != nil {
		responses := &models.Response{
			Data:          "",
			ErrorMessages: errorConvertRequestBody,
			IsSuccess:     false,
		}
		middleware.WriteResponse(w, responses, http.StatusOK)
		return
	}

	if !helper.ValidateEmailFormat(user.EmailAddress) {
		responses := &models.Response{
			Data:          "",
			ErrorMessages: invalidEmailFormat,
			IsSuccess:     false,
		}
		middleware.WriteResponse(w, responses, http.StatusOK)
		return
	}

	responses.Data = "OK"
	responses.ErrorMessages = ""
	responses.IsSuccess = true

	response, _ := json.Marshal(responses)
	w.Write(response)
}

func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var responses models.Response
	var user *models.User // defaul
	user, err := middleware.RequestBodyConverter(r)
	if err != nil {
		responses := &models.Response{
			Data:          "",
			ErrorMessages: errorConvertRequestBody,
			IsSuccess:     false,
		}
		middleware.WriteResponse(w, responses, http.StatusOK)
		return
	}

	if !helper.ValidateEmailFormat(user.EmailAddress) {
		responses := &models.Response{
			Data:          "",
			ErrorMessages: invalidEmailFormat,
			IsSuccess:     false,
		}
		middleware.WriteResponse(w, responses, http.StatusOK)
		return
	}

	responses.Data = "OK"
	responses.ErrorMessages = ""
	responses.IsSuccess = true

	response, _ := json.Marshal(responses)
	w.Write(response)
}

func SendOTPHandler(w http.ResponseWriter, r *http.Request) {
	var responses models.Response
	var otp *models.OTP // defaul
	otp, err := middleware.OTPRequestBodyConverter(r)
	if err != nil {
		responses := &models.Response{
			Data:          "",
			ErrorMessages: errorConvertRequestBody,
			IsSuccess:     false,
		}
		middleware.WriteResponse(w, responses, http.StatusOK)
		return
	}

	responses.Data = "OTP " + strconv.Itoa(otp.OTP) + " has been processed successfully"
	responses.ErrorMessages = ""
	responses.IsSuccess = true

	response, _ := json.Marshal(responses)
	w.Write(response)
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	var responses models.Response
	_, claims, _ := jwtauth.FromContext(r.Context())
	responses.Data = fmt.Sprintf("protected area. hi %v", claims["username"])
	responses.ErrorMessages = ""
	responses.IsSuccess = true

	response, _ := json.Marshal(responses)
	w.Write(response)
}

func Router(ctx context.Context, tokenAuth *jwtauth.JWTAuth) http.Handler {
	TokenAuth = tokenAuth

	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(chimiddleware.RequestID)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)

	// Basic CORS
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(cors.Handler)
	r.Use(chimiddleware.Timeout(60 * time.Second))

	//public
	r.Group(func(r chi.Router) {
		r.Post("/login", LoginHandler)
		r.Post("/register", RegisterHandler)
		r.Post("/forgotpassword", ForgotPasswordHandler)
		r.Post("/sendotp", SendOTPHandler)
	})

	//protected
	r.Group(func(r chi.Router) {

		r.Use(jwtauth.Verifier(TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/dashboard", DashboardHandler)
	})

	return r
}
