package main

import (
	"acussm/demo/common"
	 "github.com/gin-gonic/gin"
)







func main() {


	db:=common.InitDB()
	defer func() {
		mysqldb,err:=db.DB()
		if err != nil {
			panic("failed to close database"+err.Error())
		}
		mysqldb.Close()
	}()
	r := gin.Default()
		r =collectRoutE(r)
	panic(r.Run())// listen and serve on 0.0.0.0:8080
}








