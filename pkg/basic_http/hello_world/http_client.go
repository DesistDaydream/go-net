package main

import (
	"crypto/tls"
	"fmt"
	"io"
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
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// Client2 先构建一个 Request，再根据这个 Request 发起请求，这种方式常用来自定义请求内容
func Client2() {
	// 构建 [http.Request]，主要是构建 URL、Method、Header、Body 等信息
	req, _ = http.NewRequest("GET", "http://localhost:8080/hello", nil)
	// 为构建的 [http.Request] 设定请求头信息，可以多次使用 Set() 来设定多个 Header 信息
	req.Header.Set("Content-type", "application/json;charset=utf-8")

	// ######## 构建用于发起 HTTP 请求的 Client，可以指定 TLS 配置、超时时间等 ########

	// Transport(传输器) 实现了 [http.RoundTripper] 接口，负责 HTTP 请求真正的底层执行细节。包括：
	// - 建立/管理 TCP 链接
	// - TLS 握手
	// - 代理选择
	// - 协议协商
	// - etc.
	// 接口中的 [http.Transport.RoundTrip] 方法会执行单个 HTTP 事务，返回所提供请求的响应。该方法在 [http.send()] 中被调用。
	// 标准库里的 [http.Transport] 结构体有默认值。我们直接使用默认值，然后修改/添加其中的属性的值。
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true, // 跳过 HTTPS 证书验证
	}
	// 实例化一个 [http.Client] 结构体。之后，通过这个结构体的 [http.Client.Do] 方法发起 HTTP 请求
	// Client(客户端) 是基于 Transport 的更高层封装，负责策略性的东西。包括：
	// - 处理重定向
	// - Cookie
	// - 超时
	// - etc.
	// 标准库里的 [http.DefaultClient] 无默认值（i.e. `var DefaultClient = &Client{}`）。
	client := &http.Client{
		Transport: tr,
		Jar:       nil,
		Timeout:   0,
	}

	// #############################################################################

	// 根据新构建的 req，使用 client 实例发起请求。获取响应结果
	if resp, err = client.Do(req); err != nil {
		panic(err)
	}
	// 关闭连接
	defer resp.Body.Close()

	// 处理响应，并输出 Response Body
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
