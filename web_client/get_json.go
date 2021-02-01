package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Message 用来表示是一条消息应该具有的基本属性，用于绑定 JSON 各字段
type Message struct {
	Name string `json:"name"`
	Body string `json:"body"`
	Time string `json:"time"`
}

// NewMessage 实例化 Message
func NewMessage() *Message {
	return &Message{}
}

// GetJSON 将 JSON 格式的数据作为请求体，然后获取 JSON 格式的响应体数据并处理
func GetJSON() {
	// 第一种请求、模拟从外部读取 json 格式文件，将 json 与 struct 绑定，然后再发送请求
	m := NewMessage()
	// 这里模拟 struct 的值时从外部文件获取的
	m.Name = "DesistDaydream"
	m.Body = "你好"
	m.Time = time.Now().Format("2006-01-02 15:04:05")
	// 将 结构体m 中的数据，以 JSON 格式写入到 jsonData 变量中
	jsonData, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("查看JSON格式的数据是什么样的：%v\n", string(jsonData))
	// 构建 Request
	req, _ := http.NewRequest("POST", "http://localhost:8080/json", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-type", "application/json;charset=utf-8")
	// 发起请求并获取 Redponse
	if resp, err = (&http.Client{}).Do(req); err != nil {
		panic(err)
	}
	// 关闭连接
	defer resp.Body.Close()
	// 处理 Response 并输出 Body 内容
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("第一种响应体为：%v\n", string(body))

	// 第二种请求、手动指定 json 数据，并发起请求
	// 下面的代码等效于使用 crul 命令发起请求
	// curl -XPOST http://localhost:8080/json -d '{"name":"lichenhao"}'
	// 创建一个 json 数据
	jsonReqBody := []byte(`{"name":"lichenhao"}`)
	// 构建 Request
	req, _ = http.NewRequest("POST", "http://localhost:8080/json", bytes.NewBuffer(jsonReqBody))
	req.Header.Set("Content-type", "application/json")
	// 发起请求并获取 Redponse
	if resp, err = (&http.Client{}).Do(req); err != nil {
		panic(err)
	}
	// 关闭连接
	defer resp.Body.Close()
	// 处理 Response 并输出 Body 内容
	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Printf("响应体为：%v\n", string(body))
}
