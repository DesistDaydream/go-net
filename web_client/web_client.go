package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Message 是一条消息应该具有的基本属性
type Message struct {
	Name string `json:"name"`
	Body string `json:"body"`
	Time string `json:"time"`
}

// NewMessage 实例化 Message
func NewMessage() *Message {
	return &Message{}
}

// Client1 直接使用 http.Get() 来连接服务端
func Client1() {
	// net/http 标准库中还可以实现作为客户端发送 http 请求
	// Get() 向指定的服务器发送一个 HTTP GET 请求，并返回一个 Response
	resp, err := http.Get("http://172.38.40.250:8080/index")
	if err != nil {
		panic(err)
	}
	// 关闭连接
	defer resp.Body.Close()

	// 输出服务端响应的的状态码
	fmt.Println("Response status:", resp.Status)

	// 输出 Response 中 Body 的前 5 行内容
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

// Client2 先构建一个 Request，再根据这个 Request 发起请求
func Client2() {
	// 构建 Request
	req, _ := http.NewRequest("GET", "http://172.38.40.250:8080/index", nil)
	req.Header.Set("Content-type", "application/json;charset=utf-8")
	// 根据新构建的 req 来发起请求，并获取响应信息
	resp, _ := (&http.Client{}).Do(req)
	// 处理响应，并输出 Response Body
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

// GetJSON 获取 Json 数据并处理
func GetJSON() {
	// 第一种请求
	// 模拟从外部读取 json 格式文件，将 json 与 struct 绑定，然后再发送请求
	m := NewMessage()
	// 这里假定 struct 的值时从外部文件获取的
	m.Name = "DesistDaydream"
	m.Body = "你好"
	m.Time = time.Now().Format("2006-01-02 15:04:05")
	// 将 struct 转换为 json
	jsonData, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	// 构建 Request
	req, _ := http.NewRequest("POST", "http://172.38.40.250:8080/json", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-type", "application/json;charset=utf-8")
	// 处理响应信息并输出
	resp, _ := (&http.Client{}).Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	// 第二种请求
	// 手动指定 json 数据，并发起请求
	// 下面的代码等效于使用 crul 命令发起请求
	// curl -XPOST http://172.38.40.250:8080/json -d '{"name":"lichenhao"}'
	// 创建一个 json 数据
	jsonReqBody := []byte(`{"name":"lichenhao"}`)
	// 构建 Request
	req, _ = http.NewRequest("POST", "http://172.38.40.250:8080/json", bytes.NewBuffer(jsonReqBody))
	req.Header.Set("Content-type", "application/json")
	// 处理响应信息并输出 Response Body
	resp, _ = (&http.Client{}).Do(req)
	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	Client1()
	Client2()
	GetJSON()
}
