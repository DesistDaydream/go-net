package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"web/login"
	"web/order"
	// "web/test"
)

// SayHelloName 基本展示功能。w为响应给客户端的信息。r为客户端发起的请求信息。
func SayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析url传递的参数，对于POST请求则解析响应包的主体(request body)

	// 这些信息是输出到服务器端的信息
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	// 这个写入到w的，是输出到客户端的
	fmt.Fprintf(w, "Hello DesistDaydream!")
}

func main() {
	// 设置访问的路由
	http.HandleFunc("/", SayHelloName)
	http.HandleFunc("/login", login.Login)
	http.HandleFunc("/order", order.Homepage)
	http.HandleFunc("/stock-in", order.StockIn)
	http.HandleFunc("/stock-out", order.StockOut)
	http.HandleFunc("/query", order.Query)
	// http.HandleFunc("/test", test.Test)

	// 设置监听的端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
