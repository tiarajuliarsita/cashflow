package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {
	routes := app.Group("")

	V1RouteUser(routes)
	V2RouteTransaction(routes)

}
