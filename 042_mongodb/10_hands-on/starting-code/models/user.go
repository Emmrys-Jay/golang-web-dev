package models

import (
	"time"
)

type User struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type Session struct {
	Un           string
	LastActivity time.Time
}
