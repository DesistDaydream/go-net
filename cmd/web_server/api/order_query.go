package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DesistDaydream/go-net/cmd/web_server/database"
)

// 响应查询库存的结构体
type QueryStock struct {
	Supplier    string `json:"supplier"`
	Type        string `json:"type"`
	Size        string `json:"size"`
	Count       int    `json:"count"`
	StockInTime string `json:"stock_in_time"`
}

// Query 查询表单处理
func Query(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("当前客户端的请求 %v 页面的 Method 为: %v\n", r.RequestURI, r.Method)
	// 查询库存
	d := database.QueryStockInOrder()
	// 将库存信息转换为 JSON 格式的字符串
	jsonStr, _ := json.Marshal(d)
	// 响应 JSON 格式的字符串
	w.Write(jsonStr)
}
