// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package models

// Errors contains detailed error messages along with additional details
type Errors struct {
	Message    string
	StatusCode int
	Status     string
	ErrorCode  int
}
