package model

type Response struct {
	Code    int       `json:"code" binding:"required"`
	Message string    `json:"message" binding:"required"`
	Data    []Summary `json:"data,omitempty"`
	Errors  []string  `json:"errors,omitempty"`
}
