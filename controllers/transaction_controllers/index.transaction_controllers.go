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
	currentBalanceSenderUser := senderUser.Saldo - trfReq.Amount
	if trfReq.Amount > senderUser.Saldo {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "your balance is not enough",
		})
		return
	}
	senderUser.Saldo = currentBalanceSenderUser
	err = database.DB.Table("users").Where("id = ?", id).Update("saldo", senderUser.Saldo).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	//change the balance on recipient user

	currentBalanceRecipientUser := recipientUser.Saldo + trfReq.Amount
	recipientUser.Saldo = currentBalanceRecipientUser
	err = database.DB.Table("users").Where("account_number = ?", trfReq.RecipientAccountNumber).Update("saldo", recipientUser.Saldo).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	UserHistory := models.History{
		Amount:          trfReq.Amount,
		Date:            time.Now(),
		UserID:          idInt,
		TransactionID:   trfExist.ID,
		TypeTransaction: "transfer",
	}

	err = database.DB.Table("history").Create(&UserHistory).Error
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
			"saldo":                    senderUser.Saldo,
		},
	})
}

func SavingTransaction(c *gin.Context) {
	id := c.Param("id")
	savReq := new(transaction_request.SavingReq)
	user := new(models.Users)
	err := c.ShouldBind(&savReq)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "data required",
		})
		return
	}

	err = database.DB.Table("users").Where("id = ?", id).Find(&user).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}
	if user.ID == 0 {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "user not found",
		})
		return
	}

	if savReq.Pin != user.Pin {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "invalid pin",
		})
		return
	}

	userSaldo := user.Saldo + savReq.Amount
	user.Saldo = userSaldo
	idInt, _ := strconv.Atoi(id)
	saving := models.Saving{
		UserID: idInt,
		Amount: savReq.Amount,
		Date:   time.Now(),
	}
	err = database.DB.Table("saving").Create(&saving).Error
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	err = database.DB.Table("users").Where("id=?", id).Update("saldo", user.Saldo).Error
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	UserHistory := models.History{
		Amount:          savReq.Amount,
		Date:            time.Now(),
		UserID:          idInt,
		TransactionID:   saving.ID,
		TypeTransaction: "saving",
	}

	err = database.DB.Table("history").Create(&UserHistory).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "saving successfully",
		"transaction": gin.H{
			"amount": savReq.Amount,
			"saldo":  user.Saldo,
		},
	})
}

func WithDrawTransaction(c *gin.Context) {
	id := c.Param("id")
	WDReq := new(transaction_request.WithdrawReq)
	user := new(models.Users)
	err := c.ShouldBind(&WDReq)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "data required",
		})
		return
	}

	err = database.DB.Table("users").Where("id = ?", id).Find(&user).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}
	if user.ID == 0 {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "user not found",
		})
		return
	}

	if WDReq.Pin != user.Pin {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "invalid pin",
		})
		return
	}

	if WDReq.Amount > user.Saldo {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "your balance is not enough",
		})
		return
	}

	userSaldo := user.Saldo - WDReq.Amount
	user.Saldo = userSaldo
	idInt, _ := strconv.Atoi(id)
	withDraw := models.Saving{
		UserID: idInt,
		Amount: WDReq.Amount,
		Date:   time.Now(),
	}
	err = database.DB.Table("withdraw").Create(&withDraw).Error
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	UserHistory := models.History{
		Amount:          WDReq.Amount,
		Date:            time.Now(),
		UserID:          idInt,
		TransactionID:   withDraw.ID,
		TypeTransaction: "withdraw",
	}

	err = database.DB.Table("history").Create(&UserHistory).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	err = database.DB.Table("users").Where("id=?", id).Update("saldo", user.Saldo).Error
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "saving successfully",
		"transaction": gin.H{
			"amount": WDReq.Amount,
			"saldo":  user.Saldo,
		},
	})
}

func GetHistoryTransactionUser(c *gin.Context) {
	id := c.Param("id")
	// idInt, _ := strconv.Atoi("id")
	History := new([]models.History)
	user := new(models.Users)

	err := database.DB.Table("users").Where("id = ? ", id).Find(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": "internal server error",
		})
		return
	}
	if user.ID == 0 {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "user not found",
		})
		return
	}

	err = database.DB.Table("history").Where("user_id = ?", id).Find(&History).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}
	historyexist := new(models.History)
	err = database.DB.Table("history").Where("user_id = ?", id).Find(&historyexist).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}
	if historyexist.UserID == 0 {
		c.AbortWithStatusJSON(400, gin.H{
			"message": "no transaction history",
			
		})
		return
	}

	c.JSON(200, gin.H{
		"data": History,
	})
}
