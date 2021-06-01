package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"log"
	"math/rand"
	"net/http"
	"time"
)



type User struct {
	gorm.Model
	Name         string  `gorm:"type:varchar(20);not null"`       // string默认长度为255, 使用这种tag重设。
	Telephone          string     `gorm:"varchar(110;not null;unique"` // 自增
	Password 			string  `gorm:"varchar(110;not null "`
	}



func main() {

	db:=initDB()
	defer db.DB()
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {

		//获取参数
		Name := ctx.PostForm("Name")
		Telephone := ctx.PostForm("Telephone")
		Password := ctx.PostForm("Password")


		//数据验证
		if len(Telephone)!=11 {
			ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"手机号必须为11位"})
			return
		}
		if len(Password)<6 {
			ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"密码不能少于6位"})
			return
		}
		if len(Name)==0 {
			Name=RandomString(10)
		}
		if isTelephoneExits(db ,Telephone) {
			ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"用户存在"})
			return
		}

		newUser :=User{
			Name: Name,
			Telephone: Telephone,
			Password: Password,
		}
		db.Create(&newUser)


		log.Println(Name,Telephone,Password)

	 	ctx.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})



	panic(r.Run())// listen and serve on 0.0.0.0:8080
}

func isTelephoneExits(db* gorm.DB,Telephone string) bool {
	var user User
	db.Where("telephone =?", Telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}


func RandomString(n int)  string{
	var letters =[]byte("saSAsaasdsadsa")
	result := make ([]byte,n)
	rand.Seed(time.Now().Unix())
	for i :=range result{
		result[i]=letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func initDB() *gorm.DB {



	dsn := "root:root@tcp(127.0.0.1:3306)/accussm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err!=nil {
		panic("failed to connect database"+err.Error())
	}
	db.AutoMigrate(&User{})
		return  db
	}

