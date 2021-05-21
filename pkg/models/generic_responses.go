// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package models

// ErrBadRequest
type ErrBadRequest struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// ErrInternalError
type ErrInternalError struct {
	Success string `json:"success"`
	Message string `json:"message"`
}

// ErrNotFound
type ErrNotFound struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// ErrUnauthorized
type ErrUnauthorized struct {
	Error_  string `json:"error"`
	Message string `json:"message"`
}

//Success Or Failure Message
type SuccessOrErrorMessage struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg,omitempty"`
	Message string `json:"Message,omitempty"`
}
