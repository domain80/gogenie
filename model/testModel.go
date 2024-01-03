package model

import (
	"time"
)

//go:generate go run ../main.go model -n User
type User struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
