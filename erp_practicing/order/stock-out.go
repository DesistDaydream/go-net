package order

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	// mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

// StockOut 入库表单处理
func StockOut(w http.ResponseWriter, r *http.Request) {
	fmt.Println("stock-out当前客户端的请求method为：", r.Method)

	if r.Method == "GET" {
		// 使用stock-out.html响应给客户端
		t, _ := template.ParseFiles("templates/stock-out.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm() // 解析请求参数
		fmt.Println("购买客户:", r.Form["客户"])
		fmt.Println("当前型号:", r.Form["类型"])
		fmt.Println("该型号尺寸:", r.Form["尺寸"])

		fmt.Fprintf(w, "赚钱啦！！")

		// 数据处理
		db, err := sql.Open("mysql", "root:mypassword@/caredaily?charset=utf8")
		CheckErr(err)

		db.Close()
	}
}
