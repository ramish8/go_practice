package entity

import (
	"errors"
	"time"
)

type User struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func nameValidate(name string) error {
	if name == "" {
		return errors.New("name is required")
	}

	if len(name) > 100 {
		return errors.New("name is too long")
	}

	return nil
}

func NewUser(name string) (*User, error) {
	if err := nameValidate(name); err != nil {
		return nil, err
	}

	return &User{
		Name: name,
	}, nil
}

func ReconstructUser(id int64, name string, createdAt time.Time, updatedAt time.Time) (*User, error) {
	if err := nameValidate(name); err != nil {
		return nil, err
	}

	return &User{ID: id, Name: name, CreatedAt: createdAt, UpdatedAt: updatedAt}, nil
}

func (u *User) UpdateName(name string) error {
	if err := nameValidate(name); err != nil {
		return err
	}

	u.Name = name

	return nil
}
