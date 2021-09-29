// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package models

type Meta struct {
	Max    int `json:"max"`
	Offset int `json:"offset"`
	Size   int `json:"size"`
	Total  int `json:"total"`
}
