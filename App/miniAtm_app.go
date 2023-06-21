package app

import (
	"mini_Atm/database"
	"mini_Atm/routes"

	"github.com/gin-gonic/gin"
)

func App() {
	// connect to database
	app := gin.Default()
	database.ConnectDB()
	routes.InitRoutes(app)
	app.Run(":8080")
}
