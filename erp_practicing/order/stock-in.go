package order

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	// mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

// StockIn 入库表单处理
func StockIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("stock-in当前客户端的请求method为：", r.Method)

	if r.Method == "GET" {
		// 使用stock-in.html响应给客户端
		t, _ := template.ParseFiles("templates/stock-in.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm() // 解析请求参数
		fmt.Println("当前型号:", r.Form["类型"])
		fmt.Println("该型号尺寸:", r.Form["尺寸"])

		fmt.Fprintf(w, "恭喜又进货了！")

		// 数据处理
		db, err := sql.Open("mysql", "root:mypassword@/caredaily?charset=utf8")
		CheckErr(err)

		stmt, err := db.Prepare("INSERT inventory SET type=?,size=?,inventory=?")
		CheckErr(err)

		res, err := stmt.Exec(r.Form["类型"][0], r.Form["尺寸"][0], r.Form["数量"][0])
		CheckErr(err)

		id, err := res.LastInsertId()
		CheckErr(err)

		fmt.Println(id)

		db.Close()
	}
}
