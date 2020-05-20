package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// DBInfo 数据库连接信息
type DBInfo struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	DBname   string `json:"DBname"`
	Addr     string `json:"addr"`
}

// Init 初始化
func Init(info *DBInfo) {
	fmt.Println("开始连接数据库")
	var err error
	db, err = gorm.Open("mysql", info.Name+":"+info.Password+"@tcp("+info.Addr+")/"+info.DBname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln("failed to connect database, ", err)
	}

	db.DB().SetConnMaxLifetime(60 * time.Second)
	// db.LogMode(true)
	fmt.Println("开始创建数据表")
	db.AutoMigrate(&User{}, &Role{}, &Commodity{}, &CustormerOrder{}, &CustormerGoods{}, &PurchaseOrder{}, &PurchaseGoods{})

	db.Model(&Role{}).AddForeignKey("user_id", "users(user_id)", "no action", "no action")

	if err := createAdminUser(); err != nil {
		log.Fatalln(err)
	}
}
