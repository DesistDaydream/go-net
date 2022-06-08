package main

import (
	"log"
	"net/http"

	"github.com/DesistDaydream/go-net/cmd/web_server/database"
	"github.com/DesistDaydream/go-net/cmd/web_server/handler"
	"github.com/DesistDaydream/go-net/web"
)

func main() {
	// 设置访问的路由
	// 嵌入静态资源
	http.Handle("/", http.FileServer(http.FS(web.Assets)))
	// TODO: 将前后端分离，不再使用后端的模板渲染前端页面
	// 下面这些接口都将由前端调用，并将结果响应给客户端
	// 登录功能
	http.HandleFunc("/api/login", handler.Login)
	// 按钮跳转、表单处理、数据库的增删改查
	http.HandleFunc("/api/order", handler.Order)
	http.HandleFunc("/api/stock-in", handler.StockIn)
	http.HandleFunc("/api/stock-out", handler.StockOut)
	http.HandleFunc("/api/query", handler.Query)

	// 测试用接口
	// 测试用，可以测试接收到的任何 URL，以及传入的内容
	http.HandleFunc("/test", handler.Test)
	// 基本的 Web Server 端功能
	http.HandleFunc("/index", handler.Index)
	// 处理请求头，并将请求头响应给客户端
	http.HandleFunc("/header", handler.RequestHeader)
	// 处理接收到的 JSON 格式数据，并响应 JSON 格式数据给客户端
	http.HandleFunc("/json", handler.ResponseJSON)
	// 上传下载
	http.HandleFunc("/download", handler.Download)
	// Prometheus 告警接收接口
	http.HandleFunc("/alarmService/api/v1/alerts", handler.AlertmanagerV1)
	http.HandleFunc("/alarmService/api/v2/alerts", handler.AlertmanagerV2)

	db := database.NewSqlite("test.db")
	// 连接数据库
	db.ConnDB()

	// 设置监听的端口
	url := ":18080"
	log.Printf("开始监听 %v 端口\n", url)
	if err := http.ListenAndServe(url, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
