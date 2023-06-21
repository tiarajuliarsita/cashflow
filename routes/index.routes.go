package routes

import (
	"mini_Atm/controllers/transaction_controllers"
	"mini_Atm/controllers/user_controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {
	routes := app
	routes.GET("/users", user_controllers.GetAllUsers)
	routes.GET("/user/:id", user_controllers.GetUserByID)
	routes.POST("/user/create", user_controllers.CreasteUser)
	routes.DELETE("/user/delete/:id", user_controllers.DeleteUser)
	routes.PATCH("/user/update/:id", user_controllers.UpdatedUser)

	routes.POST("/user/transfer/:id", transaction_controllers.TransferTransaction)
}
