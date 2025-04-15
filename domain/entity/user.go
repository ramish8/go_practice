package entity

import (
	"errors"
	"time"
)

type User struct {
	id        int64     `db:"id"`
	name      string    `db:"name"`
	createdAt time.Time `db:"created_at"`
	updatedAt time.Time `db:"updated_at"`
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
		name: name,
	}, nil
}

func ReconstructUser(id int64, name string, createdAt time.Time, updatedAt time.Time) (*User, error) {
	if err := nameValidate(name); err != nil {
		return nil, err
	}

	return &User{id: id, name: name, createdAt: createdAt, updatedAt: updatedAt}, nil
}

func (u *User) UpdateName(name string) error {
	if err := nameValidate(name); err != nil {
		return err
	}

	u.name = name

	return nil
}

func (u *User) GetID() int64 {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *User) GetUpdatedAt() time.Time {
	return u.updatedAt
}
