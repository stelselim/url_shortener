package model

type BaseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BaseResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}
