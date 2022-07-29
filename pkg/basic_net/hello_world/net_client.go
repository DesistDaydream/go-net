package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func Client() {
	// 通过 TCP 与目标建立连接，Dial() 方法返回一个 net.Conn，对目标拨号
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

	// 从连接中读取数据
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Printf("error: %v", err)
	}

	log.Println(response)
}
