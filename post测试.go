
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

	// Example for binding JSON ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json list
		fmt.Println(json.Username+"1")
		if err := c.ShouldBindJSON(&json); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("t")
		fmt.Println(err)

		fmt.Println(json.Username)
		if json.Username != "manu" || json.Password != "123" {
		//	c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		//	return
		//}else{

		db.Create(&json)
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in","data":json})
		}
		fmt.Println(json.Username)

	})
	router.Run()
}