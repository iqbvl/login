package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/iqbvl/login/handler"
)

var TokenAuth *jwtauth.JWTAuth

func main() {
	// secret key, nanti diubah ke config
	secretKey := "aXFiYWwgYWJkdXJyYWhtYW4="

	TokenAuth = jwtauth.New("HS256", []byte(secretKey), nil)
	ctx := context.Background()
	// use 0.0.0.0 instead of localhost or 127.0.0.1 ntar ga ke expose coy
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", handler.Router(ctx, TokenAuth)))
}
