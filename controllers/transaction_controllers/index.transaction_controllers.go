package transaction_controllers

import (
	"mini_Atm/database"
	"mini_Atm/models"
	"mini_Atm/request/transaction_request"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func TransferTransaction(c *gin.Context) {
	id := c.Param("id")
	trfReq := new(transaction_request.TranferReq)
	senderUser := new(models.Users)

	err := c.ShouldBind(&trfReq)
	if err != nil {
		c.JSON(404, gin.H{
			"message": err,
		})
		return
	}

	// check account number of recipient user
	recipientUser := new(models.Users)
	err = database.DB.Table("users").Where("account_number = ?", trfReq.RecipientAccountNumber).Find(&recipientUser).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	if trfReq.RecipientAccountNumber != recipientUser.AccountNumber {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "invalid recipient account number",
		})
		return
	}

	// check pin of sender user
	err = database.DB.Table("users").Where("id = ?", id).Find(&senderUser).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}
	if senderUser.ID == 0 {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "user not found",
		})
		return
	}

	if senderUser.Pin != trfReq.Pin {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "invalid pin",
		})
		return
	}

	// set values
	idInt, _ := strconv.Atoi(id)
	trfExist := models.Transfer{
		UserID:                 idInt,
		Date:                   time.Now(),
		Amount:                 trfReq.Amount,
		RecipientAccountNumber: trfReq.RecipientAccountNumber,
	}
	// insert to database
	err = database.DB.Table("transfer").Create(&trfExist).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}
	// change the balance on sender user
	currentBalanceSenderUser := senderUser.Balance - trfReq.Amount
	if trfReq.Amount > senderUser.Balance {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "your balance is not enough",
		})
		return
	}
	senderUser.Balance = currentBalanceSenderUser
	err = database.DB.Table("users").Where("id = ?", id).Update("balance", senderUser.Balance).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	//change the balance on recipient user

	currentBalanceRecipientUser := recipientUser.Balance + trfReq.Amount
	recipientUser.Balance = currentBalanceRecipientUser
	err = database.DB.Table("users").Where("account_number = ?", trfReq.RecipientAccountNumber).Update("balance", recipientUser.Balance).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "transfer succesfully",
		"transaction": gin.H{
			"recipient account number": trfReq.RecipientAccountNumber,
			"amount":                   trfReq.Amount,
			"currentBalance":           senderUser.Balance,
		},
	})
}
