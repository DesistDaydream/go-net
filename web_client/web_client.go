package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// User is
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Example is
func Example() {
	// net/http 标准库中还可以实现作为客户端发送 http 请求
	// Get() 向指定的服务器发送一个 HTTP GET 请求，并返回一个 Response
	resp, err := http.Get("http://172.38.40.250:8080")
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

// GetJSON 获取 Json 数据并处理
func GetJSON() {
	var user User
	user.Name = "aaa"
	user.Age = 99
	if bs, err := json.Marshal(user); err == nil {
		//        fmt.Println(string(bs))
		req := bytes.NewBuffer([]byte(bs))
		tmp := `{"name":"junneyang", "age": 88}`
		req = bytes.NewBuffer([]byte(tmp))

		bodyType := "application/json;charset=utf-8"
		resp, _ := http.Post("http://172.38.40.250:8080/json/", bodyType, req)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	} else {
		fmt.Println(err)
	}

	client := &http.Client{}

	// 等效于：curl
	req := `{"name":"junneyang", "age": 88}`
	reqNew := bytes.NewBuffer([]byte(req))
	request, _ := http.NewRequest("POST", "http://172.38.40.250:8080/json/", reqNew)
	request.Header.Set("Content-type", "application/json")
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}
}

func main() {
	// Example()
	GetJSON()
}
