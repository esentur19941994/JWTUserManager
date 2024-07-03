package models

import "time"

type User struct {
	Username string    `json:"username"`
	InitDate time.Time `json:"init_date"`
}
