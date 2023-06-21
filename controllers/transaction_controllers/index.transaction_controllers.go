package transaction_controllers

import (
	"fmt"
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


	// trf := new(transaction_request.TranferReq)
	idInt, _ := strconv.Atoi(id)
	trfExist := models.Transfer{
		UserID:                 idInt,
		Date:                   time.Now(),
		Amount:                 trfReq.Amount,
		RecipientAccountNumber: trfReq.RecipientAccountNumber,
	}
	fmt.Println("ini datanya", trfExist)

	err = database.DB.Table("transfer").Create(&trfExist).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message":     "transfer succesfully",
		"transaction": trfReq,
	})
}
