package err_resp

import "github.com/gin-gonic/gin"

var UserNotFound = gin.H{
	"message": "user not found",
}
