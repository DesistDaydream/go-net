package main

import (
	"fmt"
	"net/http"
)

// SetCookie 设置 Cookie
func SetCookie(w http.ResponseWriter, r *http.Request) {
	// 将第二个参数给定的 Cookie 放在 http.ResponseWriter 的 Header 中,响应给客户端
	// 第二个参数是一个名为 Cookie 的结构体,直接传值即可。这些值就是这个 Cookie 的属性
	// 这个例子定义了一个名为 username 的 Cookie，值为 DesistDaydream，超时时间为20秒。
	http.SetCookie(w, &http.Cookie{Name: "username", Value: "DesistDaydream", MaxAge: 20})
}

// GetCookie 获取 Cookie
func GetCookie(w http.ResponseWriter, r *http.Request) {
	// 检查本次请求是否携带了名为 key_cookie 的 Cookie,并返回该 Cookie 的 value 属性。
	// 若不存在该 Cookie，则 err 会报错 "http: named cookie not present"
	cookie, _ := r.Cookie("username")
	fmt.Println(cookie)
	// 将 cookie 信息响应给客户端
	fmt.Fprint(w, cookie)
}

func main() {
	// 设置访问的路由与处理器
	http.HandleFunc("/setcookie", SetCookie)
	http.HandleFunc("/getcookie", GetCookie)

	fmt.Println("go web 启动，监听在 8080")
	// 设置监听的端口
	http.ListenAndServe(":8080", nil)
}
