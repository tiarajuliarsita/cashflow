package routes

import (
	"mini_Atm/controllers/user_controllers"

	"github.com/gin-gonic/gin"
)

func V1RouteUser(app *gin.RouterGroup) {
	route := app

	user:=route.Group("user")
	user.GET("/", user_controllers.GetAllUsers)
	user.GET("/:id", user_controllers.GetUserByID)
	user.POST("/create", user_controllers.CreasteUser)
	user.DELETE("/delete/:id", user_controllers.DeleteUser)
	user.PATCH("/update/:id", user_controllers.UpdatedUser)
}
