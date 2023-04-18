package domain

import (
	"context"
	"database/sql"
	"time"
)

type MasterDataPublication struct {
	ID                    int64                  `json:"id"`
	DefaultInformation    DefaultInformation     `json:"default_information"`
	ServiceDescription    ServiceDescription     `json:"service_description"`
	AdditionalInformation PublicationInformation `json:"additional_information"`
	Status                string                 `json:"status"`
	UpdatedAt             time.Time              `json:"updated_at"`
	CreatedAt             time.Time              `json:"created_at"`
}

type DefaultInformation struct {
	MdsID             int64      `json:"mds_id"`
	OpdName           string     `json:"opd_name"`
	ServiceForm       string     `json:"form"`
	ServiceName       string     `json:"name"`
	ProgramName       string     `json:"program_name"`
	Description       string     `json:"description"`
	ServiceUser       string     `json:"user"`
	PortalCategory    string     `json:"portal_category"`
	OperationalStatus string     `json:"operational_status"`
	Technical         string     `json:"technical"`
	Benefits          NullString `json:"benefits"`
	Facilities        NullString `json:"facilities"`
	Slug              string     `json:"slug"`
}

type ServiceDescription struct {
	Cover              NullString     `json:"cover"`
	Images             NullString     `json:"images"`
	TermsAndConditions NullString     `json:"terms_and_conditions"`
	ServiceProcedures  NullString     `json:"service_procedures"`
	ServiceFee         string         `json:"service_fee"`
	OperationalTimes   NullString     `json:"operational_times"`
	HotlineNumber      string         `json:"hotline_number"`
	HotlineMail        string         `json:"hotline_mail"`
	InfoGraphics       NullString     `json:"infographics"`
	Locations          NullString     `json:"locations"`
	Application        PubApplication `json:"application"`
	Links              NullString     `json:"links"`
	SocialMedia        NullString     `json:"social_media"`
}

type PubApplication struct {
	ID       int64      `json:"id"`
	Name     string     `json:"name"`
	Status   string     `json:"status"`
	Title    string     `json:"title"`
	Features NullString `json:"features"`
}

type PublicationInformation struct {
	Keywords NullString `json:"keywords"`
	FAQ      NullString `json:"faq"`
}

type StoreMasterDataPublication struct {
	ID                 int64 `json:"id"`
	DefaultInformation struct {
		MdsID          int64     `json:"mds_id"`
		PortalCategory string    `json:"portal_category"`
		Slug           string    `json:"slug"`
		Benefits       MdsObject `json:"benefits"`
		Facilities     MdsObject `json:"facilities"`
	} `json:"default_information" validate:"required"`
	ServiceDescription struct {
		Cover              CoverPublication       `json:"cover"`
		Images             []DetailMetaDataImage  `json:"images"`
		InfoGraphics       PublicationInfographic `json:"infographics"`
		TermsAndConditions MdsObjectCover         `json:"terms_and_conditions"`
		ServiceProcedures  MdsObjectCover         `json:"service_procedures"`
		Application        MdsApplication         `json:"application"`
	} `json:"service_description" validate:"required"`
	AdditionalInformation struct {
		Keywords []string       `json:"keywords"`
		FAQ      PublicationFAQ `json:"faq"`
	} `json:"additional_information" validate:"required"`
	Status    string    `json:"status" validate:"required,eq=PUBLISH|eq=ARCHIVE"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CoverPublication struct {
	Video string              `json:"video"`
	Image DetailMetaDataImage `json:"image"`
}

type PublicationInfographic struct {
	IsActive int8                  `json:"is_active"`
	Images   []DetailMetaDataImage `json:"images"`
}

type PublicationFAQ struct {
	IsActive int8 `json:"is_active"`
	Items    []struct {
		Question string `json:"question"`
		Answer   string `json:"answer"`
	}
}

type DetailPublicationResponse struct {
	DefaultInformation    DetailDefaultInformation    `json:"default_information"`
	ServiceDescription    DetailServiceDescription    `json:"service_description"`
	AdditionalInformation DetailAdditionalInformation `json:"additional_information"`
	Status                string                      `json:"status"`
	UpdatedAt             time.Time                   `json:"updated_at"`
	CreatedAt             time.Time                   `json:"created_at"`
}

type DetailDefaultInformation struct {
	OpdName           string    `json:"opd_name"`
	ServiceForm       string    `json:"service_form"`
	ServiceName       string    `json:"service_name"`
	ProgramName       string    `json:"program_name"`
	Description       string    `json:"description"`
	ServiceUser       string    `json:"service_user"`
	PortalCategory    string    `json:"portal_category"`
	OperationalStatus string    `json:"operator_status"`
	Technical         string    `json:"technical"`
	Benefits          MdsObject `json:"benefits"`
	Facilities        MdsObject `json:"facilities"`
	Slug              string    `json:"slug"`
}

type DetailServiceDescription struct {
	Cover              CoverPublication       `json:"cover"`
	Images             []DetailMetaDataImage  `json:"images"`
	TermsAndConditions MdsObjectCover         `json:"terms_and_conditions"`
	ServiceProcedures  MdsObjectCover         `json:"service_procedures"`
	ServiceFee         string                 `json:"service_fee"`
	OperationalTimes   []OperationalTimeMds   `json:"operational_times"`
	HotlineNumber      string                 `json:"hotline_number"`
	HotlineMail        string                 `json:"hotline_mail"`
	InfoGraphics       PublicationInfographic `json:"infographics"`
	Locations          []LocationMds          `json:"locations"`
	Application        MdsApplication         `json:"application"`
	Links              []LinkMds              `json:"links"`
	SocialMedia        []SocialMediaMds       `json:"social_media"`
}

type DetailAdditionalInformation struct {
	Keywords []string       `json:"keywords"`
	FAQ      PublicationFAQ `json:"faq"`
}

type MasterDataPublicationUsecase interface {
	Store(ctx context.Context, body *StoreMasterDataPublication) (err error)
	Fetch(ctx context.Context, au *JwtCustomClaims, params *Request) (res []MasterDataPublication, total int64, err error)
	Delete(ctx context.Context, id int64) (err error)
	GetByID(ctx context.Context, ID int64) (res MasterDataPublication, err error)
}

type MasterDataPublicationRepository interface {
	Store(ctx context.Context, body *StoreMasterDataPublication) (err error)
	GetTx(context.Context) (*sql.Tx, error)
	Fetch(ctx context.Context, params *Request) (res []MasterDataPublication, total int64, err error)
	Delete(ctx context.Context, id int64) (err error)
	GetByID(ctx context.Context, ID int64) (res MasterDataPublication, err error)
}
