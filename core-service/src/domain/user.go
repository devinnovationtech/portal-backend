package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// User ...
type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	UnitID    NullInt64 `json:"unit_id"`
	RoleID    NullInt64 `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// Author ...
type Author struct {
	Name string `json:"name"`
	// FIXME: add unit
}

// UserRepository represent the unit repository contract
type UserRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (User, error)
}