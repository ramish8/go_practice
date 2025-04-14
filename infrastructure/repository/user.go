package repository

import (
	"context"
	"database/sql"
	"go_practice/domain/entity"
	"go_practice/domain/repository"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) (int64, error) {
	row := r.db.QueryRowContext(ctx, "INSERT INTO users (name) VALUES (?) RETURNING id", user.Name)

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, name FROM users WHERE id = ?", id)

	var user entity.User
	if err := row.Scan(&user.ID, &user.Name); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) (int64, error) {
	row := r.db.QueryRowContext(ctx, "UPDATE users SET name = ? WHERE id = ? RETURNING id", user.Name, user.ID)

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = ?", id)
	return err
}
