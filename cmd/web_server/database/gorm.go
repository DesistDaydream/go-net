package database

import (
	"errors"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	UserName string
	Password string
}

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// AutoMigrate 用来刷新数据表，不存在则创建，表名默认为结构体名称的复数，e.g.这里会创建一个名为 products 的表，假如 Product 为 ProductTest，则会创建出一个名为 product_test 的表
	// 结构体中的每个字段都是该表的列，字段名称即是表中列的名称，如果字段名中有多个大写字母，则列名使用下划线分隔，e.g.CreatedAt 字段的列名为 cretaed_at
	// 当结构体中增加字段时，会自动在表中增加列；但是删除结构体中的属性时，并不会删除列
	db.AutoMigrate(&UserInfo{})

	return db
}

// 插入数据
func Insert(db *gorm.DB) {
	db.Create(&UserInfo{UserName: "lichenhao", Password: "123456"})
}

// 查询数据
func QueryUser(db *gorm.DB, username string, password string) error {
	// 声明一个 Product 数组，用来存放查询结果
	var userinfo []UserInfo
	// 查询数据，查找 code 列的值为 D42 的所有记录，并将查询结果存放到 products 中
	db.Find(&userinfo, "user_name = ?", username)
	for _, user := range userinfo {
		log.Println(user)
		if user.Password == password {
			return nil
		}
	}

	if userinfo == nil {
		// 没有查询到数据，返回一个错误
		return errors.New("没有该用户")
	} else {
		return errors.New("密码错误")
	}
}

// 更新数据
func Update(db *gorm.DB) {
	// 根据条件更新数据
	// UPDATE products SET price=300 WHERE code="D42";
	db.Model(&UserInfo{}).Where("code = ?", "D42").Update("price", 300)
}

// 删除数据
func Delete(db *gorm.DB) {
	// 根据 id 删除数据
	// DELETE FROM products WHERE id=10;
	db.Delete(&UserInfo{}, 1)
}
