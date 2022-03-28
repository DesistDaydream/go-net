package main

import (
	"fmt"
	"testing"
)

// 不同的方式建立与服务端的连接

// 直接使用 http.Get() 来发起请求
func TestHTTPClient1(t *testing.T) {
	fmt.Printf("客户端请求一、直接使用 http.Get() 来发起请求\n")
	Client1()
}

// 先构建一个 Request，再根据这个 Request 发起请求
func TestHTTPClient2(t *testing.T) {
	fmt.Printf("\n客户端请求二、先构建一个 Request，再根据这个 Request 发起请求\n")
	Client2()
}
