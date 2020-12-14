package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// printRequest 在服务器上输出客户端发起的 Request 信息
func printRequest(r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
}

// processingFormData 处理 Request 中的表单数据
func processingFormData(w http.ResponseWriter, r *http.Request) {
	// 解析客户端的 Request 数据。由于是要识别客户端输入的信息，则必须要解析，否则请求中的 body 无法识别。
	r.ParseForm()
	// 请求的是登陆数据，那么执行登陆的逻辑判断
	fmt.Println("username:", r.Form["username"])
	fmt.Println("password:", r.Form["password"])
	// 验证用户的输入是否合法
	// 用户名不能为空。获取 username 长度，如果为 0 则返回错误信息
	if len(r.Form["username"][0]) == 0 {
		fmt.Fprint(w, "错误！用户名不能为空")
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	// 解析客户端发送的 request，并输出 request 中的信息
	// printRequest(r)
	fmt.Fprintf(w, "Hello DesistDaydream!") // 这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login 请求方法为:", r.Method)
	switch r.Method {
	case "GET":
		// 使用 template 库来解析 go 的 html 模板文件
		t, _ := template.ParseFiles("./form/login.gtpl")
		// template 库中的 Execute() 将已经解析的模板文件内容写入到 w 中
		t.Execute(w, nil)
	default:
		processingFormData(w, r)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("upload 请求方法为:", r.Method)
	switch r.Method {
	case "GET":
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("form/upload.html")
		t.Execute(w, token)
	default:
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	// 设置访问的路由与处理器
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)

	fmt.Println("go web 启动，监听在 8080")
	// 设置监听的端口
	log.Fatal(http.ListenAndServe(":8080", nil))
}
