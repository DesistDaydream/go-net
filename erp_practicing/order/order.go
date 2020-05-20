package order

import (
	"fmt"
	"net/http"
	"text/template"
)

// Homepage 订单表单处理页面
func Homepage(w http.ResponseWriter, r *http.Request) {
	//显示当前请求的方法
	fmt.Println("order当前客户端的请求method为：", r.Method)

	if r.Method == "GET" {
		// 解析一个html文件,并将该页面作为响应，交给客户端
		t, _ := template.ParseFiles("templates/order.html")
		t.Execute(w, nil)
	} else {
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
		fmt.Println("测试按钮惦记", r.FormValue("test"))
	}
}
