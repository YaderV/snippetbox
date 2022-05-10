package mock

import (
	"time"

	"example.com/snippetbox/pkg/models"
)

var mockUser = &models.User{
	ID:      1,
	Name:    "Yader",
	Email:   "yader@netlandish.com",
	Created: time.Now(),
	Active:  true,
}

// UserModel ...
type UserModel struct{}

// Insert ...
func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case mockUser.Email:
		return models.ErrDuplicatedEmail
	default:
		return nil
	}
}

// Authenticate ...
func (m *UserModel) Authenticate(email, password string) (int, error) {
	switch email {
	case mockUser.Email:
		return mockUser.ID, nil
	default:
		return 0, models.ErrInvalidCredentials
	}
}

// Get ...
func (m *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case mockUser.ID:
		return mockUser, nil
	default:
		return nil, models.ErrNoRecords
	}

}
