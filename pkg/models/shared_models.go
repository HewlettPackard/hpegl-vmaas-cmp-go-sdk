// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

package models

type IDModel struct {
	ID int `json:"id,omitempty" tf:"id"`
}

type IDStringModel struct {
	ID string `json:"id,omitempty" tf:"id"`
}

type UserNameModel struct {
	UserName string `json:"username" tf:"username"`
}

type NameModel struct {
	Name string `json:"name" tf:"name"`
}

type Meta struct {
	Max    int `json:"max"`
	Offset int `json:"offset"`
	Size   int `json:"size"`
	Total  int `json:"total"`
}
