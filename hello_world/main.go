package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// HelloWorld 处理客户端请求 / 路径时的具体逻辑
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	// 解析客户端发起的 Request，默认是不会解析的
	r.ParseForm()

	// 打印客户端发起的 Request 信息
	// 这些 http Request 将会输出到服务端，也就是服务器的标准输出中
	fmt.Println(r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	// 将 Hello DesistDaydream! 写入到 Response 中，并响应给客户端
	// 这样，浏览器上就会显示 Hello DesistDaydream!
	fmt.Fprintf(w, "Hello DesistDaydream!")
}

func main() {
	// 设置访问的路由
	// 当客户端发起访问请求，访问 / 路径时，由 HelloWorld 函数处理该请求。
	http.HandleFunc("/", HelloWorld)

	// 设置监听的端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
