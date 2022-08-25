package log

import (
	"go.uber.org/zap"
)

type Log struct {
	Prefix  string
	Default *zap.SugaredLogger
}

func NewLog(prefix string) *Log {
	logger, _ := zap.NewProduction()

	return &Log{
		Prefix:  prefix,
		Default: logger.Sugar(),
	}
}

func (f Log) Info(msg string) {
	f.Default.Info(msg)
}

func (f Log) Warn(msg string) {
	f.Default.Warn(msg)
}

func (f Log) Error(msg string) {
	f.Default.Error(msg)
}

func (f Log) Fatal(msg string) {
	f.Default.Fatal(msg)
}
