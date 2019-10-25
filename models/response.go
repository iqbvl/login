package models

type Response struct {
	Data          string `json:"data"`
	ErrorMessages string `json:"errorMessages"`
	IsSuccess     bool   `json:"isSuccess"`
	//StatusCode    int    `json:"statusCode"`
}
