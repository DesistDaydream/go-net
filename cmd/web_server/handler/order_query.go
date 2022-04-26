package handler

import (
	"fmt"
	"net/http"
	"text/template"
)

// Query 查询表单处理
func Query(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("当前客户端的请求%v页面的 Method 为：%v\n", r.RequestURI, r.Method)

	switch r.Method {
	case "GET":
		// 使用query.html响应给客户端
		t, _ := template.ParseFiles("./templates/query.html")
		t.Execute(w, nil)
	default:
		// 数据处理
	}
}
