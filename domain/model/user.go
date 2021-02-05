package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// User is
type User struct {
	ID        string    `json:"id" valid:"uuid"`
	Name      string    `json:"name" valid:"notnull"`
	Email     string    `json:"email" valid:"notnull"`
	CreatedAt time.Time `created_at:"id" valid:"-"`
	UpdatedAt time.Time `updated_at:"id" valid:"-"`
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		return err
	}

	return nil
}

// NewUser does
func NewUser(name string, email string) (*User, error) {
	user := User{
		Name:  name,
		Email: email,
	}

	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()

	err := user.isValid()

	if err != nil {
		return nil, err
	}

	return &user, nil
}
