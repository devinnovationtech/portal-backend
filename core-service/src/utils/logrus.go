package utils

import (
	_goLog "github.com/jabardigitalservice/golog/logger"
	_ "github.com/newrelic/go-agent/v3/integrations/nrmysql"
)

type Logrus struct {
	logger *_goLog.Logger
}

func NewLogrus() *Logrus {
	// set formatter logs
	log := _goLog.Init()

	return &Logrus{
		logger: log,
	}
}

func (l *Logrus) Info(logsField *_goLog.LoggerData, message string) {
	l.logger.Info(logsField, message)
}

func (l *Logrus) Error(logsField *_goLog.LoggerData, e error) {
	l.logger.Error(logsField, e)
}
