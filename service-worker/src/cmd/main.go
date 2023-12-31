package main

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/job"
	"github.com/jabardigitalservice/portal-jabar-services/service-worker/src/utils"
	"github.com/robfig/cron/v3"
)

type worker struct {
	ctx  context.Context
	cfg  *config.Config
	conn *utils.Conn
}

func newWorker(ctx context.Context, cfg *config.Config, conn *utils.Conn) *worker {
	return &worker{
		ctx:  ctx,
		cfg:  cfg,
		conn: conn,
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// init worker
	w := newWorker(context.TODO(), config.NewConfig(), utils.NewDBConn(config.NewConfig()))
	go w.listenForMessages()

	// set job runner
	c := cron.New()
	cfg := config.NewConfig()

	// @every 2h is mean will run jobs every 2 hours of the day
	c.AddFunc("@every 2h", func() { job.NewsArchiveJob(w.conn, cfg) })
	c.AddFunc("@every 2h", func() { job.NewsPublishingJob(w.conn, cfg) })
	c.AddFunc("@every 2h", func() { job.PopUpBannerDeactivateJob(w.conn, cfg) })
	c.AddFunc("@every 2h", func() { job.PopUpBannerActivateJob(w.conn, cfg) })

	fmt.Println("service-worker is running")

	// start the cron job
	c.Start()
	runtime.Goexit()
}

func (w *worker) listenForMessages() {
	for {
		fmt.Println("service-worker is listening for messages")
		result, err := w.conn.Redis.BLPop(w.ctx, 0*time.Second, "email-queue").Result()

		if err != nil {
			fmt.Println(err.Error())
		} else {

			params := map[string]interface{}{}
			err := json.NewDecoder(strings.NewReader(string(result[1]))).Decode(&params)

			if err != nil {
				fmt.Println(err.Error())
			} else {
				// FIXME: make data type of params to stuct
				job.SendEmailJob(params["to"].(string), params["subject"].(string), params["body"].(string))
			}
		}
	}
}
