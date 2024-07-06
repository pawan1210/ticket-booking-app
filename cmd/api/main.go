package main

import (
	"context"

	"github.com/pawan1210/ticket-booking-app/api"
	"github.com/pawan1210/ticket-booking-app/pkg/log"
	"go.uber.org/zap"
)

func main() {
	log.InitLogger()

	ctx, cancel := context.WithCancel(context.Background())

	logger := log.GetLogger()

	logger.Info("initializing application")

	err := api.Init(ctx, cancel)

	if err != nil {
		logger.Fatal("error in initializing application", zap.Error(err))
	}
}
