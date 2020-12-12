package main

import (
	"fmt"
	"net/http"
)

// HelloWorld 处理客户端请求 /hello 时的具体逻辑
func HelloWorld(w http.ResponseWriter, req *http.Request) {
	// 将 Hello DesistDaydream! 这一串字符写入到 Response 中，并响应给客户端
	fmt.Fprintf(w, "Hello DesistDaydream!")
}

func main() {
	// 设置访问的路由
	// 当客户端发起 http 请求，访问 http://IP:8080/hello ，由 HelloWorld 函数处理该请求。
	http.HandleFunc("/hello", HelloWorld)

	// 设置监听的端口
	http.ListenAndServe(":8080", nil)
}
