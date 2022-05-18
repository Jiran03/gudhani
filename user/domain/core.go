package domain

import "time"

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Address   string
	Gender    string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
