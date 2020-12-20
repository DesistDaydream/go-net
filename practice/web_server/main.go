package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DesistDaydream/GoWeb/practice/web_server/handler"
)

func main() {
	// 设置访问的路由
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/order", handler.Order)
	http.HandleFunc("/stock-in", handler.StockIn)
	http.HandleFunc("/stock-out", handler.StockOut)
	http.HandleFunc("/query", handler.Query)
	http.HandleFunc("/header", handler.RequestHeader)

	// 设置监听的端口
	fmt.Println("开始监听8080端口")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
