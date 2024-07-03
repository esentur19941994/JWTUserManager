package models

import "time"

type User struct {
	Username string    `json:"username" bson:"username"`
	InitDate time.Time `json:"init_date" bson:"init_date"`
}
