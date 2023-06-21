package user_controllers

import (
	"mini_Atm/database"
	"mini_Atm/models"
	"mini_Atm/request/user_req"
	"mini_Atm/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreasteUser(c *gin.Context) {
	userReq := new(user_req.UsersReq)
	err := c.ShouldBind(&userReq)
	if err != nil {
		c.JSON(400, gin.H{

			"message": "required data",
		})

		return
	}

	userExist := new(models.Users)
	database.DB.Table("users").Where("user_name = ?", userReq.UserName).First(&userExist)
	if userExist.UserName != "" {
		c.JSON(200, gin.H{
			"message": "username all ready exist",
		})
		return
	}

	NumberAccount := utils.RandomNumberAccount(7)

	user := models.Users{
		Email:         userReq.Email,
		AccountNumber: NumberAccount,
		Balance:       userReq.Balance,
		BornDate:      userReq.BornDate,
		UserName:      userReq.UserName,
		PhoneNumber:   userReq.PhoneNumber,
		Pin:           userReq.Pin,
	}

	err = database.DB.Table("users").Create(&user).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		// log.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"message": "sucessfully",
		"user":    user,
	})

}

func GetAllUsers(c *gin.Context) {
	users := new([]models.Users)
	err := database.DB.Table("users").Find(&users).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}
	c.JSON(200, gin.H{
		"Data": users,
	})

}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user := new(models.Users)
	err := database.DB.Table("users").Where("id = ? ", id).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": user,
	})

}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	user := new(models.Users)
	err := database.DB.Table("users").Where("id=?", id).Find(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}
	err = database.DB.Table("users").Where("id =?", id).Unscoped().Delete(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSONP(200, gin.H{
		"message": "deleted succesfully",
	})

}

func UpdatedUser(c *gin.Context) {
	id := c.Param("id")
	userReq := new(user_req.UsersReq)

	err := c.ShouldBind(&userReq)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "user not found",
		})
	}
	user := new(models.Users)
	err = database.DB.Table("users").Where("id=?", id).Find(&user).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}
	if user.ID == 0 {
		c.JSON(404, gin.H{
			"message": "user not found",
		})
		return
	}

	user.UserName = userReq.UserName
	user.Balance = userReq.Balance
	user.PhoneNumber = userReq.PhoneNumber
	user.Email = userReq.Email
	user.BornDate = userReq.BornDate
	user.Pin = userReq.Pin

	err = database.DB.Table("users").Where("id=?", id).Updates(&user).Error
	if err!= nil{
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
	}
	c.JSON(200, gin.H{
		"message":"updated succesfully",
		"data":user,
	})

}
