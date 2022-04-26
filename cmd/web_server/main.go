package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DesistDaydream/go-net/cmd/web_server/handler"
)

func main() {
	// 设置访问的路由
	// 测试用，可以测试接收到的任何 URL，以及传入的内容
	http.HandleFunc("/", handler.Test)
	// 基本的 Web Server 端功能
	http.HandleFunc("/index", handler.Index)
	// 处理请求头，并将请求头响应给客户端
	http.HandleFunc("/header", handler.RequestHeader)
	// 处理接收到的 JSON 格式数据，并响应 JSON 格式数据给客户端
	http.HandleFunc("/json", handler.ResponseJSON)
	// 登录功能
	http.HandleFunc("/login", handler.Login)
	// 按钮跳转、表单处理、数据库的增删改查
	http.HandleFunc("/order", handler.Order)
	http.HandleFunc("/stock-in", handler.StockIn)
	http.HandleFunc("/stock-out", handler.StockOut)
	http.HandleFunc("/query", handler.Query)
	// 上传下载
	http.HandleFunc("/download", handler.Download)
	// Prometheus 告警接收接口
	http.HandleFunc("/alarmService/api/v1/alerts", handler.AlertmanagerV1)
	http.HandleFunc("/alarmService/api/v2/alerts", handler.AlertmanagerV2)

	// 设置监听的端口
	url := ":18080"
	fmt.Printf("开始监听 %v 端口\n", url)
	if err := http.ListenAndServe(url, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
