package main

import (
	"fmt"
	"log"
	"net/http"

	sessioncookie "github.com/DesistDaydream/GoWeb/sessioncookie"
	_ "github.com/DesistDaydream/GoWeb/sessionmemory"
)

var globalSessions *sessioncookie.SessionManager

func init() {
	var err error
	globalSessions, err = sessioncookie.NewSessionManager("sessionmemory", "goSessionid", 3600)
	if err != nil {
		fmt.Println(err)
		return
	}
	go globalSessions.GC()
	fmt.Println("fd")
}

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("name")
	if err == nil {
		fmt.Println(cookie.Value)
		fmt.Println(cookie.Domain)
		fmt.Println(cookie.Expires)
	}
	//fmt.Fprintf(w, "Hello world!\n") //这个写入到w的是输出到客户端的
}
func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	val := sess.Get("username")
	if val != nil {
		fmt.Println(val)
	} else {
		sess.Set("username", "DesistDaydream")
		fmt.Println("set session")
	}
}
func loginOut(w http.ResponseWriter, r *http.Request) {
	//销毁
	globalSessions.SessionDestroy(w, r)
	fmt.Println("session destroy")
}

func main() {
	http.HandleFunc("/", sayHelloHandler) //	设置访问路由
	http.HandleFunc("/login", login)
	http.HandleFunc("/loginout", loginOut) //销毁
	log.Fatal(http.ListenAndServe(":8080", nil))
}
