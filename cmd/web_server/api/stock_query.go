package api

import (
	"encoding/json"
	"net/http"

	"github.com/DesistDaydream/go-net/cmd/web_server/database"
	"github.com/sirupsen/logrus"
)

// StockQuery 查询表单处理
func StockQuery(w http.ResponseWriter, r *http.Request) {
	logrus.WithFields(logrus.Fields{
		"方法": r.Method,
		"端点": r.RequestURI,
	}).Infof("客户端请求")

	// 查询库存
	d := database.QueryStockInOrder()
	// 将库存信息转换为 JSON 格式的字符串
	jsonByte, _ := json.Marshal(d)
	// 响应 JSON 数据给前端
	w.Write(jsonByte)
}
