package main

import (
	"fmt"
	"food_delivery/component"
	"food_delivery/modules/websites/tranport/do_gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dst := "root@tcp(127.0.0.1:3306)/users"
	db, err := gorm.Open(mysql.Open(dst), &gorm.Config{})
	router := gin.Default()
	appCtx := component.NewAppCtx(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	router.POST("/signup", do_gin.Register(appCtx))
	router.GET("/get", do_gin.GetList(appCtx))
	router.Run()
}
