package repository

import (
	"context"
	"go_practice/domain/entity"
	"go_practice/domain/repository"
	"go_practice/infrastructure/model"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) (int64, error) {
	model := model.FromDomain(user)
	row := r.db.QueryRowContext(ctx, "INSERT INTO users (name, created_at, updated_at) VALUES ($1, $2, $3) RETURNING id", model.Name, time.Now(), time.Now())

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, name FROM users WHERE id = $1", id)

	user := model.User{}
	if err := row.Scan(&user.ID, &user.Name); err != nil {
		return nil, err
	}

	return user.ToDomain()
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) (int64, error) {
	row := r.db.QueryRowContext(ctx, "UPDATE users SET name = $1 WHERE id = $2 RETURNING id", user.GetName(), user.GetID())

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id)
	return err
}
