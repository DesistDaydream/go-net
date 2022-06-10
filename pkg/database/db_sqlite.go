package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Sqilte 数据库连接信息
type Sqlite struct {
	Database string
}

func NewSqlite(database string) *Sqlite {
	return &Sqlite{
		Database: database,
	}
}

// ConnDB 连接数据库
func (s *Sqlite) ConnDB() {
	db, err = gorm.Open(sqlite.Open(s.Database), &gorm.Config{})
	if err != nil {
		logrus.Fatalln("failed to connect database, ", err)
	}

	// AutoMigrate 用来刷新数据表，不存在则创建，表名默认为结构体名称的复数，e.g.这里会创建一个名为 products 的表，假如 Product 为 ProductTest，则会创建出一个名为 product_test 的表
	// 结构体中的每个字段都是该表的列，字段名称即是表中列的名称，如果字段名中有多个大写字母，则列名使用下划线分隔，e.g.CreatedAt 字段的列名为 cretaed_at
	// 当结构体中增加字段时，会自动在表中增加列；但是删除结构体中的属性时，并不会删除列
	db.AutoMigrate(&StockInOrder{}, &User{})

	// 创建管理员用户
	if err := createAdminUser(); err != nil {
		logrus.Fatalln(err)
	}
}
