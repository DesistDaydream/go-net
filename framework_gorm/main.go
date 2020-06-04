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

var (
	db       *gorm.DB
	err      error
	products []Product
)

func main() {
	// 初始化数据库
	Init()
	defer db.Close()

	// 插入数据
	Insert()

	// 删除数据
	// Delete()

	// 更新数据
	Update()

	// 查询数据
	Query()
}

// Init 初始化数据库,连接数据库并根据结构体建立数据表
func Init() {
	// 连接mysql数据库,连接格式为：USER:PASSWORD@PROTOCOL(IP:PORT)/DBNAME?AGRS
	db, err = gorm.Open("mysql", "root:mypassword@tcp(10.10.100.200:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	// AutoMigrate 用来刷新数据表，不存在则创建，表名默认为结构体名称的复数，e.g.这里创建完的表名为 products 。
	// 当结构体中增加字段时，会自动在表中增加列，如果字段名中有多个大写字母，则列名使用下划线分隔，e.g.CreatedAt 字段的列名为 cretaed_at 。
	// 但是删除结构体中的字段时，并不会删除列。
	db.AutoMigrate(&Product{})
}

// Insert 插入数据
func Insert() {
	// INSERT INTO products (code,price) VALUES ("L1211","3000");
	db.Create(&Product{Code: "L1212", Price: 1000})
}

// Delete 删除数据
func Delete() {
	db.Delete(&products)
}

// Update 更新数据
func Update() {
	// 更新product的price为2000
	db.Model(&products).Update("Price", 2000)
}

// Query 查询数据
func Query() {
	//	查询 products 变量关联的结构体，所定义的数据表中的所有数据，然后将查询到的数据保存到 Product 结构体中
	// SELECT * FROM products;
	db.Find(&products)
	for _, i := range products {
		fmt.Println(i)
	}
}
