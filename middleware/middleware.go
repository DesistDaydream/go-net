package main

import (
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello DesistDaydream!"))
	// 让 Handler 运行 2 秒，模拟一下处理时间以便观察中间件的处理逻辑
	time.Sleep(2 * time.Second)
}

// timeMiddleware 手动实现了一个中间件框架
func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// 声明一个当前时间，以便计算各种代码逻辑的耗时
		t := time.Now()

		// next handler
		// 任何方法实现了 ServerHTTP，即是一个合法的 http.Handler
		next.ServeHTTP(w, req)

		log.Println("中间加+handler 共消耗时间：", time.Since(t))
	})
}

func main() {
	http.Handle("/", timeMiddleware(http.HandlerFunc(hello)))
	http.ListenAndServe(":8080", nil)
}
