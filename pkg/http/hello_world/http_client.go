package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 设置一些会用到的全局变量，省的每次都要重新初始化
var (
	req  *http.Request
	resp *http.Response
	err  error
)

// Client1 直接使用 http.Get() 来发起请求
func Client1() {
	// net/http 标准库中可以实现作为客户端发送 http 请求
	// Get() 向指定的服务器发送一个 HTTP GET 请求，并返回一个 Response
	if resp, err = http.Get("http://localhost:8080/hello"); err != nil {
		panic(err)
	}
	// 关闭连接
	defer resp.Body.Close()

	// 处理 Response 中的 Body，并输出响应体的字符串格式内容
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// Client2 先构建一个 Request，再根据这个 Request 发起请求，这种方式常用来自定义请求内容
func Client2() {
	// 构建 Request
	req, _ = http.NewRequest("GET", "http://localhost:8080/hello", nil)
	// 为构建的 Request 设定请求头信息，可以多次使用 Set() 来设定多个 Header 信息
	req.Header.Set("Content-type", "application/json;charset=utf-8")

	// 根据新构建的 req 来发起请求，并获取响应信息
	// 这里的 http.Client{} 中可以设置一些发起 HTTP 请求时的一些信息，比如 TLS 等
	if resp, err = (&http.Client{}).Do(req); err != nil {
		panic(err)
	}
	// 关闭连接
	defer resp.Body.Close()

	// 处理响应，并输出 Response Body
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
