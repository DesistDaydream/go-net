package main

import (
	"net"

	"log"
)

// 处理连接
func handleConn(conn net.Conn) {
	defer conn.Close()
	// 定义缓冲区
	buf := make([]byte, 1024)

	// 读取客户端数据
	conn.Read(buf[:1024])

	// 将数据写回客户端
	len, err := conn.Write([]byte("hello,i am server"))
	if err != nil {
		log.Panicf("error: %v", err)
	}

	log.Println(len)
}

func main() {
	// 实例化监听器
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	// 监听并接受连接
	for {
		// 等待客户端连接
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error: %v", err)
		}

		// 创建goroutine处理客户端连接
		go handleConn(conn)
	}
}
