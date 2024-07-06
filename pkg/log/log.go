package log

import (
	"sync"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	once   sync.Once
)

func InitLogger() {
	once.Do(func() {
		var err error
		logger, err = zap.NewProduction()

		if err != nil {
			panic(err)
		}
	})
}

func GetLogger() *zap.Logger {
	if logger == nil {
		InitLogger()
	}

	return logger
}
