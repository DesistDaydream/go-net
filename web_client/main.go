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

// Example1 直接使用 http.Get() 来发起请求
func Example1() {
	// net/http 标准库中可以实现作为客户端发送 http 请求
	// Get() 向指定的服务器发送一个 HTTP GET 请求，并返回一个 Response
	if resp, err = http.Get("http://172.38.40.250:8080/index"); err != nil {
		panic(err)
	}
	// 关闭连接
	defer resp.Body.Close()

	// 输出服务端响应的的状态码
	fmt.Println("Response status:", resp.Status)

	// 输出 Response 中 Body 内容
	// 处理响应，并输出 Response Body
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// Example2 先构建一个 Request，再根据这个 Request 发起请求
// 除了 Example1 中的方法，还可以先构建一个 Request，再根据这个 Request 发起 http 请求，这种方式常用来自定义请求内容
func Example2() {
	// 构建 Request
	req, _ = http.NewRequest("GET", "http://172.38.40.250:8080/index", nil)
	// 为构建的 Request 设定请求头信息，可以多次使用 Set() 来设定多个 Header 信息
	req.Header.Set("Content-type", "application/json;charset=utf-8")
	// 查看一下将要发起的请求内容
	fmt.Printf("本次 HTTP Request 为：%v\n请求方法为：%v\n请求头为：%v\n", req, req.Method, req.Header)

	// 根据新构建的 req 来发起请求，并获取响应信息
	if resp, err = (&http.Client{}).Do(req); err != nil {
		panic(err)
	}
	// 关闭连接
	defer resp.Body.Close()

	// 处理响应，并输出 Response Body
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	// 不同的方式建立与服务端的连接
	Example1()
	Example2()
	// 像服务端传递请求体，获取 JSON 数据并处理、输出
	GetJSON()
}
