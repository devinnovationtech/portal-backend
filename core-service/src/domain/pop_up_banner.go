package domain

import (
	"context"
	"time"
)

type PopUpBanner struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	ButtonLabel string     `json:"button_label,omitempty"`
	Image       NullString `json:"image,omitempty"`
	Link        string     `json:"link"`
	Status      string     `json:"status"`
	Duration    int64      `json:"duration,omitempty"`
	IsLive      int8       `json:"is_live"`
	StartDate   *time.Time `json:"start_date,omitempty"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type ListPopUpBannerResponse struct {
	ID        int64       `json:"id"`
	Title     string      `json:"title"`
	Image     ImageBanner `json:"image,omitempty"`
	Link      string      `json:"link"`
	Duration  int64       `json:"duration,omitempty"`
	StartDate *time.Time  `json:"start_date,omitempty"`
	Status    string      `json:"status"`
	IsLive    int8        `json:"is_live"`
}

type DetailPopUpBannerResponse struct {
	ID            int64               `json:"id"`
	Title         string              `json:"title"`
	ButtonLabel   string              `json:"button_label"`
	Image         ImageBanner         `json:"image,omitempty"`
	ImageMetaData ImageMetaDataBanner `json:"image_metadata,omitempty"`
	Link          string              `json:"link"`
	Status        string              `json:"status"`
	IsLive        int8                `json:"is_live"`
	Duration      int64               `json:"duration"`
	StartDate     *time.Time          `json:"start_date"`
	EndDate       *time.Time          `json:"end_date,omitempty"`
	UpdateAt      time.Time           `json:"updated_at"`
}

type ImageBanner struct {
	Desktop string `json:"desktop,omitempty"`
	Mobile  string `json:"mobile,omitempty"`
}

type MetaDetailPopUpBannerResponse struct {
	ID            int64               `json:"id"`
	Title         string              `json:"title"`
	ButtonLabel   string              `json:"button_label"`
	ImageMetaData ImageMetaDataBanner `json:"image,omitempty"`
	Link          string              `json:"link"`
	Status        string              `json:"status"`
	IsLive        int8                `json:"is_live"`
	Duration      int64               `json:"duration"`
	StartDate     *time.Time          `json:"start_date"`
	EndDate       *time.Time          `json:"end_date,omitempty"`
	UpdateAt      time.Time           `json:"updated_at"`
}

type ImageMetaDataBanner struct {
	Desktop DetailMetaDataImage `json:"desktop,omitempty"`
	Mobile  DetailMetaDataImage `json:"mobile,omitempty"`
}

type DetailMetaDataImage struct {
	FileName        string `json:"file_name"`
	FileDownloadUri string `json:"file_download_uri"`
	Size            int64  `json:"size"`
}

type StorePopUpBannerRequest struct {
	ID           int64                `json:"id"`
	Title        string               `json:"title" validate:"required,max=255"`
	CustomButton CustomButtonlabel    `json:"custom_button,omitempty"`
	Scheduler    SchedulerPopUpBanner `json:"scheduler,omitempty"`
	Image        ImageBanner          `json:"image" validate:"required"`
}

type UpdateStatusPopUpBannerRequest struct {
	Status   string `json:"status" validate:"required,eq=ACTIVE|eq=NON-ACTIVE"`
	IsLive   int64  `json:"is_live,omitempty"`
	Duration int64  `json:"duration,omitempty"`
}

type CustomButtonlabel struct {
	Label string `json:"label"`
	Link  string `json:"link"`
}

type SchedulerPopUpBanner struct {
	Duration    int64   `json:"duration"`
	StartDate   *string `json:"start_date"`
	IsScheduled int64   `json:"is_scheduled"`
	Status      string  `json:"status"`
}

type LiveBannerResponse struct {
	ID          int64       `json:"id"`
	Title       string      `json:"title"`
	ButtonLabel string      `json:"button_label"`
	Image       ImageBanner `json:"image,omitempty"`
	Link        string      `json:"link"`
	Status      string      `json:"status"`
	IsLive      int8        `json:"is_live"`
	Duration    int64       `json:"duration"`
	StartDate   *time.Time  `json:"start_date"`
	EndDate     *time.Time  `json:"end_date,omitempty"`
	UpdateAt    time.Time   `json:"updated_at"`
}
type PopUpBannerUsecase interface {
	Fetch(ctx context.Context, auth *JwtCustomClaims, params *Request) (res []PopUpBanner, total int64, err error)
	GetByID(ctx context.Context, id int64) (res PopUpBanner, err error)
	Store(ctx context.Context, auth *JwtCustomClaims, body *StorePopUpBannerRequest) (err error)
	Delete(ctx context.Context, id int64) (err error)
	UpdateStatus(ctx context.Context, id int64, body *UpdateStatusPopUpBannerRequest) (err error)
	Update(ctx context.Context, auth *JwtCustomClaims, id int64, body *StorePopUpBannerRequest) (err error)
	GetMetaDataImage(ctx context.Context, link string) (meta DetailMetaDataImage, err error)
	LiveBanner(ctx context.Context) (res PopUpBanner, err error)
}

type PopUpBannerRepository interface {
	Fetch(ctx context.Context, params *Request) (res []PopUpBanner, total int64, err error)
	GetByID(ctx context.Context, id int64) (res PopUpBanner, err error)
	Store(ctx context.Context, body *StorePopUpBannerRequest) (err error)
	Delete(ctx context.Context, id int64) (err error)
	UpdateStatus(ctx context.Context, id int64, body *UpdateStatusPopUpBannerRequest) (err error)
	DeactiveStatus(ctx context.Context) (err error)
	Update(ctx context.Context, id int64, body *StorePopUpBannerRequest) (err error)
	LiveBanner(ctx context.Context) (res PopUpBanner, err error)
}
