package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	dsn := "root:huqianlong123@tcp(127.0.0.1:3306)/qing?charset=utf8mb4&parseTime=True&loc=Local"
	// 也可以使用MustConnect连接不成功就panic
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)

	type list struct {
		gorm.Model
		Username string `gorm:"type:varchar(20); not null" json:"username" binding:"required"`
		Password string `gorm:"type:varchar(20);not null" json:"password" binding:"required"`
		Status   string `gorm:"type:varchar(20);not null" json:"status" binding:"required"`
	}

	//test := list{Username: "manu", Password: "123", Status: "1"}
	//result := db.Create(&test)
	//fmt.Println(result)
	router := gin.Default()
	router.GET("/welcome/:id", func(c *gin.Context) {

		id := c.Param("id")

		sqlstr := `select count(username) from list where id = ?`
		var count int
		if err := db.Get(&count, sqlstr, id); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"400": "用户不存在"})
			return
		}
		if count > 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"400": "用户存在"})
			return
		}

	})

	router.Run()
}
