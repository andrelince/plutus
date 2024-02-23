package logger

import (
	"github.com/sirupsen/logrus"
)

type Logger *logrus.Logger

type Settings struct {
}

func New(settings Settings) Logger {
	return logrus.New()
}
