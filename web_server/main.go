package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DesistDaydream/GoWeb/web_server/handler"
)

func main() {
	// 设置访问的路由
	// 测试用，可以测试接收到的任何 URL，以及传入的内容
	http.HandleFunc("/", handler.Test)
	// 基础服务端功能
	http.HandleFunc("/index", handler.Index)
	// 登录功能
	http.HandleFunc("/login", handler.Login)
	// 按钮跳转、数据库操作
	http.HandleFunc("/order", handler.Order)
	http.HandleFunc("/stock-in", handler.StockIn)
	http.HandleFunc("/stock-out", handler.StockOut)
	http.HandleFunc("/query", handler.Query)
	// 输出请求头信息
	http.HandleFunc("/header", handler.RequestHeader)
	// 接收 JSON 格式请求体，并响应 JSON 格式内容
	http.HandleFunc("/json", handler.ResponseJSON)

	// 设置监听的端口
	fmt.Println("开始监听8080端口")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
