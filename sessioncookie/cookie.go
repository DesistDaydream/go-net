package seesioncookie

import (
	"fmt"
	"net/http"
)

// SetCookie 设置 cookie
func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "username", Value: "DesistDaydream"}
	http.SetCookie(w, &cookie)
}

// GetCookie 获取 cookie
func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("username")
	fmt.Println(cookie)
	fmt.Fprint(w, cookie)
}
