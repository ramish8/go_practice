package model

import (
	"go_practice/domain/entity"
	"time"
)

type User struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FromDomain(e *entity.User) *User {
	return &User{
		ID:   e.GetID(),
		Name: e.GetName(),
	}
}

func (u *User) ToDomain() (*entity.User, error) {
	return entity.ReconstructUser(u.ID, u.Name, u.CreatedAt, u.UpdatedAt)
}
