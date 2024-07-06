package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Init(
	router *gin.Engine,
) error {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	internalRouter := router.Group("/v1/internal")
	userRouter := internalRouter.Group("/user")

	initUserRouter(userRouter)

	return router.Run(fmt.Sprintf(":%d", 8000))
}
