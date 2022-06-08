package handler

import (
	"fmt"
	"net/http"
	"text/template"
)

// StockIn 入库表单处理
func StockIn(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("当前客户端的请求%v页面的 Method 为：%v\n", r.RequestURI, r.Method)

	switch r.Method {
	case "GET":
		// 使用stock-in.html响应给客户端
		t, _ := template.ParseFiles("./templates/stock-in.html")
		t.Execute(w, nil)
	default:
		r.ParseForm() // 解析请求参数
		fmt.Println("当前型号:", r.Form["类型"])
		fmt.Println("该型号尺寸:", r.Form["尺寸"])
		fmt.Fprintf(w, "恭喜又进货了！")

		// 数据处理
		// 在此处写数据处理的具体逻辑
	}
}

// StockOut 出库表单处理
func StockOut(w http.ResponseWriter, r *http.Request) {
	fmt.Println("stock-out当前客户端的请求method为：", r.Method)

	switch r.Method {
	case "GET":
		// 使用stock-out.html响应给客户端
		t, _ := template.ParseFiles("./templates/stock-out.html")
		t.Execute(w, nil)
	default:
		r.ParseForm() // 解析请求参数
		fmt.Println("购买客户:", r.Form["客户"])
		fmt.Println("当前型号:", r.Form["类型"])
		fmt.Println("该型号尺寸:", r.Form["尺寸"])

		fmt.Fprintf(w, "赚钱啦！！")
	}
}
