package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/DesistDaydream/go-net/pkg/database"
	"github.com/sirupsen/logrus"
)

// StockIn 入库表单处理
func StockIn(w http.ResponseWriter, r *http.Request) {
	// 允许跨域访问，当使用 live server 响应前端代码时方便调试
	w.Header().Set("Access-Control-Allow-Origin", "*")

	logrus.WithFields(logrus.Fields{
		"方法": r.Method,
		"端点": r.RequestURI,
	}).Infof("客户端请求")

	RequestBody, _ := io.ReadAll(r.Body)
	logrus.Infof("请求体：%v", string(RequestBody))

	// 获取 JSON 请求体
	var stockInOrder database.StockInOrder
	err := json.Unmarshal(RequestBody, &stockInOrder)
	if err != nil {
		logrus.Error("解析 JSON 请求体失败：", err)
		return
	}

	// 数据处理
	stockInOrder.AddStockInOrder()
	w.Write([]byte("ok"))
}
