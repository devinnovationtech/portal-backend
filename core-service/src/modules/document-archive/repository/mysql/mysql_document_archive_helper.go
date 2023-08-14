package mysql

import (
	"fmt"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

/**
 * this block of code is used to generate the query for the event
 * need to be refactored later to be more generic and reduce the code complexity (go generic tech debt)
 */
func filterDocArchiveQuery(params *domain.Request, binds *[]interface{}) string {
	var query string

	if params.Keyword != "" {
		*binds = append(*binds, `%`+params.Keyword+`%`)
		query += ` AND d.title LIKE ?`
	}

	if v, ok := params.Filters["category"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND d.category = ?`, query)
	}

	if v, ok := params.Filters["status"]; ok && v != "" {
		*binds = append(*binds, v)
		query = fmt.Sprintf(`%s AND d.status = ?`, query)
	}

	return query
}
