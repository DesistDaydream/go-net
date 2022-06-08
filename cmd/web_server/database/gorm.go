package database

import (
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	UserName string
	Password string
}

// 插入数据
func Insert(db *gorm.DB) {
	db.Create(&UserInfo{UserName: "lichenhao", Password: "123456"})
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
