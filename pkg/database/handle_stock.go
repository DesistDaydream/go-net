package database

import (
	"github.com/sirupsen/logrus"
)

// StockInOrder 入库订单的属性，后面的描述信息用来绑定属性与表单中的字段
type StockInOrder struct {
	// gorm.Model
	Provider string `form:"provider" binding:"required"`
	Commodity
}

// StockOutOrder 出库订单的属性，后面的描述信息用来绑定属性与表单中的字段
type StockOutOrder struct {
	// gorm.Model
	Customer string `form:"customer" binding:"required"`
	Commodity
}

// Commodity 一个商品应该具有的属性
type Commodity struct {
	Product string `form:"product" binding:"required"`
	Type    string `form:"type" binding:"required"`
	Size    string `form:"size" binding:"required"`
	Amount  int    `form:"amount" binding:"required"`
}

// AddStockInOrder 在 stock-in.go 页面中向数据库添加入库订单数据
func (i *StockInOrder) AddStockInOrder() {
	// 数据库处理逻辑，结构体中的数据在调用本方法之前已经处理好了。
	db.Create(i)
}

// QueryStockInOrder 在 stock-query.html 页面中查询数据库中的入库订单数据
func QueryStockInOrder() (StockInOrders []StockInOrder) {
	// 数据库处理逻辑
	db.Find(&StockInOrders)

	// 在后台打印查询的数据，检查数据正确性
	logrus.Info("数据库中的数据为：", StockInOrders)
	return
}

// AddStockOutOrder 在 stock-out.html 页面添加出库订单数据
func (o *StockOutOrder) AddStockOutOrder() {
	logrus.Info("待更新")
}
