package controller

import (
	"acussm/demo/common"
	"acussm/demo/dtoo"
	"acussm/demo/model"
	"acussm/demo/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"acussm/demo/respond"
)

func Register(ctx *gin.Context) {

db:=common.GetDB()

//获取参数
Name := ctx.PostForm("Name")
Telephone := ctx.PostForm("Telephone")
Password := ctx.PostForm("Password")


//数据验证
if len(Telephone)!=11 {
	respond.Respond(ctx,http.StatusUnprocessableEntity,422,nil,"手机号必须为11位")
	return
}

if len(Password)<6 {
	respond.Respond(ctx,http.StatusUnprocessableEntity,422,nil,"密码不能少于6位")
return
}

if len(Name)==0 {
Name=util.RandomString(10)
}

if isTelephoneExits(db ,Telephone) {
	respond.Respond(ctx,http.StatusUnprocessableEntity,422,nil,"用户存在")

return
}

	hasedpassword ,err:=bcrypt.GenerateFromPassword([]byte (Password),bcrypt.DefaultCost)
	if err != nil {
		respond.Respond(ctx,http.StatusUnprocessableEntity,500,nil,"加密错误")
		return
	}

	newUser :=model.User{
	Name: Name,
	Telephone: Telephone,
	Password: string(hasedpassword),
}
	db.Create(&newUser)


	log.Println(Name,Telephone,Password)


	respond.Success(ctx,nil,"注册成功")

}

func isTelephoneExits(db* gorm.DB,Telephone string) bool {
	var user model.User
	db.Where("telephone =?", Telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
func Login(ctx *gin.Context)  {
	db:=common.GetDB()
	//传参
	Telephone := ctx.PostForm("Telephone")
	Password := ctx.PostForm("Password")


	//数据验证
	if len(Telephone)!=11 {
		respond.Respond(ctx,http.StatusUnprocessableEntity,500,nil,"手机号必须为11位")
		return
	}
	if len(Password)<6 {
		respond.Respond(ctx,http.StatusUnprocessableEntity,500,nil,"密码不能少于6位")

		return

	}
	//判断手机号是否存在

	var user model.User
	db.Where("telephone =?", Telephone).First(&user)
	if user.ID==0 {
		respond.Respond(ctx,http.StatusUnprocessableEntity,422,nil,"用户不存在")
		return
	}




//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(Password) ); err != nil {
		respond.Respond(ctx,http.StatusUnprocessableEntity,400,nil,"密码错误")

		return
	}


	newUser :=model.User{

		Telephone: Telephone,
		Password:  Password,
	}
	db.Create(&newUser)




	token ,err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{"code":500,"msg":"系统异常"})
		log.Printf("token generate error",err)
		return
	}
	//返回结果

	respond.Success(ctx,gin.H{"token":token},"登陆成功")

}

func Info(ctx *gin.Context)  {
	user , _:=ctx.Get("user")
	ctx.JSON(http.StatusOK,gin.H{"code":200,"data":gin.H{"user":dtoo.Touser(user.(model.User))}})

}