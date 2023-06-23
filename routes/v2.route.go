package routes

import (
	"mini_Atm/controllers/transaction_controllers"

	"github.com/gin-gonic/gin"
)

func V2RouteTransaction(app *gin.RouterGroup) {
	route := app
	transaction := route.Group("transaction")

	transaction.POST("/transfer/:id", transaction_controllers.TransferTransaction)
	transaction.POST("/saving/:id", transaction_controllers.SavingTransaction)
	transaction.POST("/withdraw/:id", transaction_controllers.WithDrawTransaction)
	transaction.GET("/history/:id", transaction_controllers.GetHistoryTransactionUser)
}
