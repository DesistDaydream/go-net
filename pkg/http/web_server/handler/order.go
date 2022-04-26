package handler

import (
	"fmt"
	"net/http"
	"text/template"
)

// CheckErr 检查数据库操作
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Order 订单表单处理页面
func Order(w http.ResponseWriter, r *http.Request) {
	// 允许跨域访问
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//显示当前请求的方法
	fmt.Printf("当前客户端的请求%v页面的 Method 为：%v\n", r.RequestURI, r.Method)

	// 不同的 Request Method 进行不同处理
	switch r.Method {
	case "GET":
		// 解析一个html文件,并将该页面作为响应，交给客户端
		t, _ := template.ParseFiles("./templates/order.html")
		t.Execute(w, nil)
	default:
		// r.FormValue() 与 r.Form 不同，会自动调用r.ParseForm对客户端提交的参数进行解析
		// 这样就不用提前调用r.ParseForm了
		switch r.FormValue("order") {
		case "入库":
			http.Redirect(w, r, "/stock-in", http.StatusFound)
		case "出库":
			http.Redirect(w, r, "/stock-out", http.StatusFound)
		case "查询":
			http.Redirect(w, r, "/query", http.StatusFound)
		}
		fmt.Println("测试按钮点击", r.FormValue("test"))
	}
}
