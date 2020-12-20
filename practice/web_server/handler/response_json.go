package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Message 是一条消息应该具有的基本属性
type Message struct {
	Name string `json:"name"`
	Body string `json:"body"`
	Time int64  `json:"time"`
}

// NewMessage 实例化 Message
func NewMessage() *Message {
	return &Message{
		Name: "DesistDaydream",
		Body: "Hello World",
		Time: 0,
	}
}

// ResponseJSON 将会响应 JSON 格式数据
func ResponseJSON(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("当前客户端的请求 %v 页面的 Method 为：%v\n", r.RequestURI, r.Method)
	body, _ := ioutil.ReadAll(r.Body)
	// r.Body.Close()
	fmt.Printf("请求体为：%v\n", string(body))
	// fmt.Fprint(w, string(body))

	m := NewMessage()
	if bs, err := json.Marshal(m); err == nil {
		fmt.Println(string(bs))
	} else {
		fmt.Println(err)
	}

	if err := json.Unmarshal(body, &m); err == nil {
		fmt.Println(m)
		ret, _ := json.Marshal(m)
		fmt.Fprint(w, string(ret))
	} else {
		fmt.Println(err)
	}
}
