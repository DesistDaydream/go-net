package order

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	// mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

// Query 查询表单处理
func Query(w http.ResponseWriter, r *http.Request) {
	fmt.Println("query当前客户端的请求method为：", r.Method)

	if r.Method == "GET" {
		// 使用query.html响应给客户端
		t, _ := template.ParseFiles("templates/query.html")
		t.Execute(w, nil)
	} else {
		// 数据处理
		db, err := sql.Open("mysql", "root:mypassword@/caredaily?charset=utf8")
		CheckErr(err)

		db.Close()
	}
}
