package routes

import (
	"github.com/gin-gonic/gin"
	usercontroller "github.com/pawan1210/ticket-booking-app/api/controllers/user"
	"github.com/pawan1210/ticket-booking-app/api/utils"
	userservice "github.com/pawan1210/ticket-booking-app/internal/services/user"
	"github.com/pawan1210/ticket-booking-app/internal/services/user/types"
	userstore "github.com/pawan1210/ticket-booking-app/internal/stores/user"
)

func initUserRouter(router *gin.RouterGroup) {
	userStore := userstore.NewInMemoryStore()
	userService := userservice.New(userStore)
	userController := usercontroller.New(userService)

	router.POST("", utils.ControllerWrapper(types.AddUserRequest{}, userController.AddUser))
}
