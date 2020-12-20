package handler

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

	switch r.Method {
	case "GET":
		// 使用query.html响应给客户端
		t, _ := template.ParseFiles("practice/templates/query.html")
		t.Execute(w, nil)
	default:
		// 数据处理
		db, err := sql.Open("mysql", "root:mypassword@/caredaily?charset=utf8")
		CheckErr(err)

		db.Close()
	}
}
