package mysql

import (
	"context"
	"database/sql"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/helpers"
)

type mysqlMainServiceRepository struct {
	Conn *sql.DB
}

// NewMysqlMainServiceRepository will create an object that represent the GeneralInformation.Repository interface
func NewMysqlMainServiceRepository(Conn *sql.DB) domain.MainServiceRepository {
	return &mysqlMainServiceRepository{Conn}
}

func (m *mysqlMainServiceRepository) Store(ctx context.Context, ms *domain.StoreMasterDataService) (ID int64, err error) {
	query := `
	INSERT main_services SET opd_name=?, government_affair=?, sub_government_affair=?, service_form=?, 
	service_type=?, sub_service_type=?, service_name=?, program_name=?, description=?, service_user=?, 
	sub_service_spbe=?, operational_status=?, technical=?, benefits=?, facilities=?, website=?, links=?, 
	terms_and_condition=?, service_procedures=?, service_fee=?, operational_time=?, hotline_number=?, 
	hotline_mail=?, location=?
	`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx,
		&ms.Services.Information.OpdName,
		&ms.Services.Information.GovernmentAffair,
		&ms.Services.Information.SubGovernmentAffair,
		&ms.Services.Information.ServiceForm,
		&ms.Services.Information.ServiceType,
		&ms.Services.Information.SubServiceType,
		&ms.Services.Information.ServiceName,
		&ms.Services.Information.ProgramName,
		&ms.Services.Information.Description,
		&ms.Services.Information.ServiceUser,
		&ms.Services.Information.SubServiceSpbe,
		&ms.Services.Information.OperationalStatus,
		&ms.Services.Information.Technical,
		helpers.GetStringFromObject(&ms.Services.Information.Benefits),
		helpers.GetStringFromObject(&ms.Services.Information.Facilities),
		&ms.Services.Information.Website,
		helpers.GetStringFromObject(&ms.Services.Information.Links),
		helpers.GetStringFromObject(&ms.Services.ServiceDetail.TermsAndConditions),
		helpers.GetStringFromObject(&ms.Services.ServiceDetail.ServiceProcedures),
		&ms.Services.ServiceDetail.ServiceFee,
		helpers.GetStringFromObject(&ms.Services.ServiceDetail.OperationalTime),
		&ms.Services.ServiceDetail.HotlineNumber,
		&ms.Services.ServiceDetail.HotlineMail,
		helpers.GetStringFromObject(&ms.Services.Location),
	)
	if err != nil {
		return
	}
	ID, err = res.LastInsertId()
	if err != nil {
		return
	}

	return
}
