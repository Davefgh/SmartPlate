package models

import "time"

type Vehicle struct {
	ID      int       `json:"id" db:"id"`
	Make    string    `json:"make" db:"make"`
	Model   string    `json:"model" db:"model"`
	Plate   string    `json:"plate" db:"plate"`
	Created time.Time `json:"created" db:"created"`
	Updated time.Time `json:"updated" db:"updated"`
}