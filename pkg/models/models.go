package models

import (
	"errors"
	"time"
)

var (
	// ErrNoRecords ...
	ErrNoRecords = errors.New("models: no matching record found")
	// ErrInvalidCredentials ...
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// ErrDuplicatedEmail ...
	ErrDuplicatedEmail = errors.New("models: duplicated email")
)

// Snippet ...
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// User ...
type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}
