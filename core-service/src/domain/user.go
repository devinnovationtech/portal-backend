package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// User ...
type User struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Photo     NullString `json:"photo"`
	Unit      Unit       `json:"unit"`
	UnitName  string     `json:"unit_name"`
	RoleID    NullInt64  `json:"role_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt time.Time  `json:"deleted_at"`
}

// UserInfo ...
type UserInfo struct {
	ID       uuid.UUID  `json:"id"`
	Name     string     `json:"name"`
	Username string     `json:"username"`
	Email    string     `json:"email"`
	Photo    NullString `json:"photo"`
	Unit     Unit       `json:"unit"`
	RoleID   NullInt64  `json:"role_id"`
}

// Author ...
type Author struct {
	Name     string `json:"name"`
	UnitName string `json:"unit_name"`
}

// UserRepository represent the unit repository contract
type UserRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	Store(context.Context, *User) error
}

// UserUsecase ...
type UserUsecase interface {
	Store(context.Context, *User) error
}
