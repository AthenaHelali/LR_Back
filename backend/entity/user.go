package entity

import "time"

type User struct {
	ID          uint
	PhoneNumber string
	Name        string
	// Password always keeps hashed password
	Password  string
	CreatedAt time.Time
	Role      Role
}
