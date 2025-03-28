package models

import "time"

type User struct {
	USER_ID        int       `json:"user_id" db:"user_id"`
	LAST_NAME string    `json:"last_name" db:"last_name"`
	FIRST_NAME string    `json:"first_name" db:"first_name"`
	MIDDLE_NAME string    `json:"middle_name" db:"middle_name"`
	EMAIL string    `json:"email" db:"email"`
	PASSWORD string    `json:"password" db:"password"`
	ROLE string    `json:"role" db:"role"`
	STATUS string    `json:"status" db:"status"`
	CREATED time.Time `json:"created" db:"created"`
	UPDATED time.Time `json:"updated" db:"updated"`
	LTO_CLIENT_ID string    `json:"lto_client_id" db:"lto_client_id"`
}