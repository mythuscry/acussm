package middlleware

import (
	"acussm/demo/common"
	"acussm/demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Authmiddleware()gin.HandlerFunc  {

	return func(context *gin.Context) {
	tokenstring :=context.GetHeader("authorization")

		if tokenstring=="" ||strings.HasPrefix(tokenstring,"bearer ") {
			context.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			context.Abort()
			return
		}
			tokenstring=tokenstring[7:]

			token,claims ,err:= common.Parsetoken(tokenstring)

			if err != nil||!token.Valid {
			context.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			context.Abort()
			return
		}

		userid :=claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user,userid)

		if userid ==0{
			context.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			context.Abort()
			return
		}
		context.Set("user",user)
		context.Next()

	}

}