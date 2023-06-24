package routes

import (
	"mini_Atm/controllers/user_controllers"
	"mini_Atm/database"
	"mini_Atm/repository"
	"mini_Atm/service"

	"github.com/gin-gonic/gin"
)

func V1RouteUser(app *gin.RouterGroup) {
	route := app

	user := route.Group("user")
	userRepository := repository.NewRepository(database.DB)
	userService := service.NewService(userRepository)
	userHandler := user_controllers.NewUserHandler(userService)

	user.GET("/", userHandler.GetAllUsers)
	user.GET("/:id", userHandler.GetUserByID)
	user.POST("/create", user_controllers.CreasteUser)
	user.DELETE("/delete/:id", user_controllers.DeleteUser)
	user.PATCH("/update/:id", user_controllers.UpdatedUser)
}
