package login

import (
	"fmt"
	"net/http"
	"text/template"
)

// Login 登录相关的表单处理功能
func Login(w http.ResponseWriter, r *http.Request) {
	//显示当前请求的方法
	fmt.Println("login页面当前客户端的请求method为：", r.Method)

	if r.Method == "GET" {
		// 解析一个html文件,并将该页面作为响应，交给客户端
		t, _ := template.ParseFiles("login/login.html")
		t.Execute(w, nil)
	} else {
		//解析客户端的请求信息。由于是要识别客户端输入的信息，则必须要解析，否则请求中的body无法识别。
		r.ParseForm()
		// r.Form 用来获取客户端输入的内容，[] 中的字符串对应前端页面给定的name中的标识符。该内容是一个切片。
		fmt.Println("用户输入的username为:", r.Form["username"])
		fmt.Println("用户输入的password为:", r.Form["password"])
		fmt.Println("备注信息：", r.Form["note"])

		// 验证客户端输入的内容,如果输入错误则返回错误信息。
		if r.Form["username"][0] != "zhangna" {
			fmt.Fprintf(w, "用户名错误，请重新输入")
			return
		}

		// 登录之后，跳转到order首页页面来继续处理客户端的请求
		http.Redirect(w, r, "/order", http.StatusFound)
	}
}
