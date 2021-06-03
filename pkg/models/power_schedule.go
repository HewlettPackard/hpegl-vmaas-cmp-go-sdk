// (C) Copyright 2021 Hewlett Packard Enterprise Development LP
package models

import "time"

type GetAllPowerSchedulesSchedules struct {
	ID                     int         `json:"id"`
	Name                   string      `json:"name"`
	Description            interface{} `json:"description"`
	Enabled                bool        `json:"enabled"`
	Scheduletype           string      `json:"scheduleType"`
	Scheduletimezone       string      `json:"scheduleTimezone"`
	Sundayon               float64     `json:"sundayOn"`
	Sundayoff              float64     `json:"sundayOff"`
	Mondayon               float64     `json:"mondayOn"`
	Mondayoff              float64     `json:"mondayOff"`
	Tuesdayon              float64     `json:"tuesdayOn"`
	Tuesdayoff             float64     `json:"tuesdayOff"`
	Wednesdayon            float64     `json:"wednesdayOn"`
	Wednesdayoff           float64     `json:"wednesdayOff"`
	Thursdayon             float64     `json:"thursdayOn"`
	Thursdayoff            float64     `json:"thursdayOff"`
	Fridayon               float64     `json:"fridayOn"`
	Fridayoff              float64     `json:"fridayOff"`
	Saturdayon             float64     `json:"saturdayOn"`
	Saturdayoff            float64     `json:"saturdayOff"`
	Totalmonthlyhourssaved float64     `json:"totalMonthlyHoursSaved"`
	Datecreated            time.Time   `json:"dateCreated"`
	Lastupdated            time.Time   `json:"lastUpdated"`
}

type GetAllPowerSchedules struct {
	Schedules []GetAllPowerSchedulesSchedules `json:"schedules"`
}
