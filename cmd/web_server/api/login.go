package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/DesistDaydream/go-net/cmd/web_server/database"
	"github.com/sirupsen/logrus"
)

// 响应结构体
type LoginResponse struct {
	Code  int    `json:"code"`
	Token string `json:"token"`
	Msg   string `json:"msg"`
}

// 生成 Token
func GenerateToken(username string) string {
	return fmt.Sprintf("%s:%s", username, "123456")
}

// Login 登录相关的表单处理功能
func Login(w http.ResponseWriter, r *http.Request) {
	// 允许跨域访问
	w.Header().Set("Access-Control-Allow-Origin", "*")

	logrus.Printf("当前客户端的请求 %v 页面的 Method 为：%v\n", r.RequestURI, r.Method)
	// 不同的 Request Method 进行不同处理
	switch r.Method {
	case "GET":
		// 解析一个html文件,并将该页面作为响应，交给客户端
		t, _ := template.ParseFiles("./web/templates/login.html")
		t.Execute(w, nil)
	default:
		username := r.FormValue("username")
		password := r.FormValue("password")
		logrus.Println("用户输入的信息为: ", r.Form)

		// 根据请求体中的用户名和密码，查询数据库，判断是否存在该用户
		if _, err := database.VerifyUser(username, password); err != nil {
			logrus.Printf("%v 登录失败，原因：%v", username, err)
			resp := LoginResponse{
				Code:  0,
				Token: "",
				// 返回登录失败信息
				Msg: "登录失败，用户名或密码错误",
			}
			// 将结构体转换为 JSON 格式的字符串
			jsonStr, _ := json.Marshal(resp)

			// 响应 401 状态码 或者 响应 JSON 格式的字符串
			// w.WriteHeader(http.StatusUnauthorized)
			w.Write(jsonStr)
		} else {
			logrus.Printf("%v 登录成功", username)
			// 生成 Token
			token := GenerateToken(username)
			// 响应结构体
			resp := LoginResponse{
				Code:  1,
				Token: token,
				Msg:   "登录成功",
			}
			// 将结构体转换为 JSON 格式的字符串
			jsonStr, _ := json.Marshal(resp)

			// 将 JSON 格式的字符串响应给客户端
			w.Write(jsonStr)
		}
	}
}
