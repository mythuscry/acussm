package respond
import (
	"github.com/gin-gonic/gin"
	 "net/http"
	)


func Respond(ctx *gin.Context,httpstatus,code int ,data gin.H,msg string)  {
	ctx.JSON(httpstatus,gin.H{"code":code,"data":data,"msg":msg})
}

func Success(ctx *gin.Context,data gin.H,msg string)  {
	Respond(ctx,http.StatusOK,200,data,msg)
}

func Fali(ctx *gin.Context,data gin.H,msg string)  {
	Respond(ctx ,http.StatusOK,400,data,msg)
}
