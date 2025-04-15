package api

import "time"

// NOTE: 面倒なのでハードコードするが、本来はcodegenで生成する
type CreateUserRequest struct {
	Name string `json:"name"`
}

type CreateUserResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
