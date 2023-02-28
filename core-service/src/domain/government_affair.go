package domain

import "context"

type GovernmentAffair struct {
	ID            int64  `json:"id"`
	MainAffair    string `json:"main_affair"`
	SubMainAffair string `json:"sub_main_affair"`
}

type GovernmentAffairUsecase interface {
	Fetch(ctx context.Context, params *Request) (res []GovernmentAffair, total int64, err error)
}

type GovernmentAffairRepository interface {
	Fetch(ctx context.Context, params *Request) (res []GovernmentAffair, total int64, err error)
}
