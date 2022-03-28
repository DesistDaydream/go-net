package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func Client1() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer client.Close()

	var reply string
	// 调用名为 HelloSerevice 服务下的 Hello() 方法
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
