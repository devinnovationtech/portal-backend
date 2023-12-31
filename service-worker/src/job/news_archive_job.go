package job

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/utils"
	"github.com/sirupsen/logrus"
)

// NewsArchiveJob ...
func NewsArchiveJob(conn *utils.Conn, cfg *config.Config) {
	logrus.Println("NewsArchiveJob is running")

	// Get news ids from news will be archived
	var newsIDs []string
	query := `SELECT id FROM news
		WHERE status='PUBLISHED'
		AND is_live=1
		AND duration <> -1
		AND end_date < NOW()`

	rows, err := conn.Mysql.Query(query)
	if err != nil {
		logrus.Error(err)
		return
	}
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			logrus.Error(err)
			return
		}
		newsIDs = append(newsIDs, id)
	}

	if len(newsIDs) == 0 {
		return
	}

	// Archive news
	updateQuery := fmt.Sprintf(`UPDATE news SET is_live=0, status='ARCHIVED' WHERE id IN (%s)`, strings.Join(newsIDs, ","))
	res, err := conn.Mysql.Exec(updateQuery)
	if err != nil {
		logrus.Error(err)
	}

	// rows affected
	if ra, err := res.RowsAffected(); err != nil {
		logrus.Error("ErrNewsArchiveJob: ", err)
	} else {
		logrus.Println("NewsArchiveJob: Rows affected: ", ra)
	}

	// elasticapi update body set is_live=0 where body id in (news ids) and domain=news
	doc := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"terms": map[string]interface{}{
							"id": newsIDs,
						},
					},
					{
						"match": map[string]interface{}{
							"domain": "news",
						},
					},
				},
			},
		},
		"script": map[string]interface{}{
			"source": "ctx._source.is_active=params.is_active;",
			"lang":   "painless",
			"params": map[string]interface{}{
				"is_active": false,
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		logrus.Error("Error encoding doc", err)
	}

	// update elasticsearch
	esRes, esErr := conn.Elastic.UpdateByQuery(
		[]string{cfg.ELastic.IndexContent},
		conn.Elastic.UpdateByQuery.WithBody(&buf),
		conn.Elastic.UpdateByQuery.WithContext(context.Background()),
	)

	if esErr != nil {
		logrus.Error("Error Update response", err)
	}
	defer esRes.Body.Close()
	fmt.Println(esRes.String())
}
