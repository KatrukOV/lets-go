package models

import (
	"time"
)

// Snippet ...
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type Snippets []*Snippet

// User ...
type UserCreate struct {
	Name     string
	Email    string
	Password string
}

// UserLogin
type UserLogin struct {
	Email    string
	Password string
}
