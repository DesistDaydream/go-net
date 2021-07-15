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
	// 注意，Get() 方法本质上，还是调用的是 http.do() 方法，http.do() 方法的示例在本代码的 Client2 中进行展示。
	if resp, err = http.Get("http://localhost:8080/index"); err != nil {
		panic(err)
	}
	// 关闭连接
	defer resp.Body.Close()

	// 输出服务端响应的的状态码
	fmt.Println("Response status:", resp.Status)

	// 处理 Response 中的 Body，并输出响应体的字符串格式内容
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// Client2 先构建一个 Request，再根据这个 Request 发起请求，这种方式常用来自定义请求内容
func Client2() {
	// 构建 Request
	req, _ = http.NewRequest("GET", "http://localhost:8080/index", nil)
	// 为构建的 Request 设定请求头信息，可以多次使用 Set() 来设定多个 Header 信息
	req.Header.Set("Content-type", "application/json;charset=utf-8")
	// 查看一下将要发起的请求内容
	fmt.Printf("查看本次 HTTP 请求的信息：\n请求内容：%+v\n请求方法：%v\n请求头：%v\n", req, req.Method, req.Header)

	// 根据新构建的 req 来发起请求，并获取响应信息
	// http.Client{} 结构体就是一个 HTTP 客户端。结构体中可以设置一些发起 HTTP 请求时的一些信息，比如 TLS 等
	http := &http.Client{}
	if resp, err = http.Do(req); err != nil {
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
	// ############################################
	// ########直接使用 http.Get() 来发起请求########
	// ############################################
	fmt.Printf("客户端请求一、直接使用 http.Get() 来发起请求\n")
	// Client1()
	// ############################################################
	// ########先构建一个 Request，再根据这个 Request 发起请求########
	// ############################################################
	fmt.Printf("\n客户端请求二、先构建一个 Request，再根据这个 Request 发起请求\n")
	Client2()
	// ##########################################################
	// ########向服务端传递请求体，获取 JSON 数据并处理、输出########
	// ##########################################################
	fmt.Printf("\n客户端请求三、向服务端传递请求体，获取 JSON 数据并处理、输出\n")
	// GetJSON()
}
