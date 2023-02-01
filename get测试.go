
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/schema"
	"net/http"
	"time"

	"gorm.io/gorm"
)

func main() {
	fmt.Println(time.Now())
	dsn := "root:huqianlong123@tcp(127.0.0.1:3306)/qing?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	fmt.Println(db)
	fmt.Println(err)
	sqlDB, err := db.DB()
	fmt.Println(time.Second)
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10*time.Second)
	type list struct {
		gorm.Model
		Username string `gorm:"type:varchar(20); not null" json:"username" binding:"required"`
		Password string `gorm:"type:varchar(20);not null" json:"password" binding:"required"`
		Status string `gorm:"type:varchar(20);not null" json:"status" binding:"required"`
	}

	//test := list{Username: "manu", Password: "123", Status: "1"}
	//result := db.Create(&test)
	//fmt.Println(result)
	router := gin.Default()
	router.GET("/welcome/:id", func(c *gin.Context) {
		var data []list

		id := c.Param("id")
		db.Where("id=?",id).Find(&data)
		fmt.Println(data)
		if len(data) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"400": "失败"})
			return
		}else{

			c.JSON(http.StatusUnauthorized, gin.H{"200": "成功","data":data})
		}
	})


	router.Run()
}