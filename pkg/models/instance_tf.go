// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

// This file contains struct models related to terraform, but not the API
package models

type TFInstance struct {
	History    []GetInstanceHistoryProcesses `tf:"history"`
	Containers []GetInstanceContainer        `tf:"containers"`
}
