package api

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pawan1210/ticket-booking-app/api/routes"
)

func Init(ctx context.Context, cancel context.CancelFunc) error {
	router := gin.New()

	routes.Init(router)

	return router.Run(fmt.Sprintf(":%d", 8000))
}
