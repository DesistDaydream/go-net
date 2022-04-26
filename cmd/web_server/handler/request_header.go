package handler

import (
	"fmt"
	"net/http"
)

// RequestHeader 输出请求头的信息
func RequestHeader(w http.ResponseWriter, req *http.Request) {
	// 读取 HTTP Request 中的 Header 中的所有内容
	// 并将这些请求头信息，写入到 Response 中，并响应给客户端
	for name, headers := range req.Header {
		// name 与 headers 是 请求头中的键/值对。
		// 每个 headers 都是数组，再通过 range 循环数组中的元素
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
