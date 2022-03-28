package main

import (
	"fmt"
	"testing"
)

func TestRPCClient(t *testing.T) {
	fmt.Printf("客户端请求一、直接使用 http.Get() 来发起请求\n")
	Client1()
}
