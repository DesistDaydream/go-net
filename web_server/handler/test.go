package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Test is
func Test(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Prometheus Push 的 URL 为：%v\n", req.URL)
	fmt.Printf("Prometheus Push 的 Method 为 %v\n", req.Method)
	body, _ := ioutil.ReadAll(req.Body)
	fmt.Printf("Prometheus Push 的 内容 为：%v\n", string(body))
	fmt.Fprintf(w, "测试页面!")
}
