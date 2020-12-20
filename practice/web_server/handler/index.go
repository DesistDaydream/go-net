package handler

import (
	"fmt"
	"net/http"
)

// Index 基本展示功能。w为响应给客户端的信息。r为客户端发起的请求信息。
func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析url传递的参数，对于POST请求则解析响应包的主体(request body)

	// 显示当前请求的方法
	// 这是输出到服务端的信息，由于这是定义在 / 路径下的，所有访问 / 下的任何路径，都会输出这些信息
	fmt.Printf("%v页面当前客户端的请求method为：%v\n", r.Host, r.Method)

	// 这个写入到w的，是输出到客户端的
	fmt.Fprintf(w, "Hello DesistDaydream!")
}
