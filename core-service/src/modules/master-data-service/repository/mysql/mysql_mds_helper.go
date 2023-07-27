package mysql

import (
	"fmt"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

/**
 * this block of code is used to generate the query for the news
 * need to be refactored later to be more generic and reduce the code complexity (go generic tech debt)
 */
func filterMdsQuery(params *domain.Request, binds *[]interface{}) string {
	var query string

	if params.Keyword != "" {
		*binds = append(*binds, `%`+params.Keyword+`%`)
		query = fmt.Sprintf(`%s AND ms.service_name LIKE ?`, query)
	}

	if v, ok := params.Filters["status"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND mds.status = ?`, query)
	}

	if v, ok := params.Filters["created_by"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND mds.created_by = ?`, query)
	}

	if v, ok := params.Filters["opd_name"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND units.name = ?`, query)
	}

	if v, ok := params.Filters["service_user"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND ms.service_user = ?`, query)
	}

	if v, ok := params.Filters["technical"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND ms.technical = ?`, query)
	}

	if params.StartDate != "" && params.EndDate != "" {
		*binds = append(*binds, params.StartDate, params.EndDate)
		query = fmt.Sprintf(`%s AND (DATE(mds.updated_at) BETWEEN ? AND ?)`, query)
	}

	return query
}
