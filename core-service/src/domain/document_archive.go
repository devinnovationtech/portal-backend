package domain

import (
	"context"
	"time"
)

// DocumentArchive Struct ...
type DocumentArchive struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Excerpt     string    `json:"excerpt" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Source      string    `json:"source"`
	Mimetype    string    `json:"mimetype"`
	Status      string    `json:"status"`
	Category    string    `json:"category" validate:"required"`
	CreatedBy   User      `json:"created_by"`
	UpdatedBy   User      `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ListDocumentArchive ...
type ListDocumentArchive struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Excerpt     string    `json:"excerpt"`
	Description string    `json:"description"`
	Source      string    `json:"source"`
	Mimetype    string    `json:"mimetype"`
	Status      string    `json:"status"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DocumentArchiveRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Source      string `json:"source" validate:"required"`
	Mimetype    string `json:"mimetype" validate:"required"`
	Status      string `json:"status" validate:"required,eq=DRAFT|eq=PUBLISHED|eq=ARCHIVED"`
	Category    string `json:"category" validate:"required"`
}

type UpdateStatusDocumentArchiveRequest struct {
	Status string `json:"status" validate:"required,eq=DRAFT|eq=PUBLISHED|eq=ARCHIVED"`
}

// DocumentArchiveUsecase ...
type DocumentArchiveUsecase interface {
	Fetch(ctx context.Context, params *Request) ([]DocumentArchive, int64, error)
	Store(ctx context.Context, body *DocumentArchiveRequest, createdBy string) error
	Update(ctx context.Context, body *DocumentArchiveRequest, UpdatedBy string, ID int64) error
	Delete(ctx context.Context, ID int64) error
	GetByID(ctx context.Context, ID int64) (DocumentArchive, error)
	TabStatus(ctx context.Context) ([]TabStatusResponse, error)
	UpdateStatus(ctx context.Context, body *UpdateStatusDocumentArchiveRequest, updatedBy string, ID int64) error
}

// DocumentArchiveRepository ...
type DocumentArchiveRepository interface {
	Fetch(ctx context.Context, params *Request) ([]DocumentArchive, int64, error)
	Store(ctx context.Context, body *DocumentArchiveRequest, createdBy string) error
	Update(ctx context.Context, body *DocumentArchiveRequest, UpdatedBy string, ID int64) error
	Delete(ctx context.Context, ID int64) error
	GetByID(ctx context.Context, ID int64) (DocumentArchive, error)
	TabStatus(ctx context.Context) ([]TabStatusResponse, error)
	UpdateStatus(ctx context.Context, body *UpdateStatusDocumentArchiveRequest, updatedBy string, ID int64) error
}
