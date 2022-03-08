package job

import (
	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/utils"
	"github.com/sirupsen/logrus"
)

// NewsPublishingJob ...
func NewsPublishingJob(conn *utils.Conn) {
	logrus.Println("NewsPublishingJob is running")
	res, err := conn.Mysql.Exec("UPDATE news SET is_live=1, published_at = now() WHERE status='PUBLISHED' AND is_live=0 AND start_date <= NOW()")
	if err != nil {
		logrus.Error(err)
	}

	// rows affected
	ra, err := res.RowsAffected()

	if err != nil {
		logrus.Error(err)
	}

	logrus.Println("NewsPublishingJob: Rows affected: ", ra)

	return
}