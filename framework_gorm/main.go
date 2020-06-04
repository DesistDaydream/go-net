package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Product struct {
	// gorm.Model 是一个包含了ID, CreatedAt, UpdatedAt, DeletedAt四个字段的结构体。
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// 连接mysql数据库,连接格式为：USER:PASSWORD@PROTOCOL(IP:PORT)/DBNAME?AGRS
	db, err := gorm.Open("mysql", "root:mypassword@tcp(10.10.100.200:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// AutoMigrate 用来刷新数据表，不存在则创建，表名默认为结构体名称的复数，e.g.这里创建完的表名为 products 。
	// 当结构体中增加字段时，会自动在表中增加列，如果字段名中有多个大写字母，则列名使用下划线分隔，e.g.CreatedAt 字段的列名为 cretaed_at 。
	// 但是删除结构体中的字段时，并不会删除列。
	db.AutoMigrate(&Product{})

	// 插入数据
	db.Create(&Product{Code: "L1212", Price: 1000})

	// 读取
	var products Product
	rows, _ := db.Find(&products).Rows()
	for rows.Next() {
		rows.Scan(products)
		fmt.Println(products)
	}

	// 更新 - 更新product的price为2000
	db.Model(&products).Update("Price", 2000)

	// 删除 - 删除product
	db.Delete(&products)
}
