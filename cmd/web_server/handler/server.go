package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func Server(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("当前客户端的请求 %v 页面的 Method 为：%v\n", r.RequestURI, r.Method)

	// 根据不同请求方法，执行不同的行为
	switch r.Method {
	case "GET":
		logrus.Infof("对方发起了 GET 请求")
	default:
		// 如果请求方法不是 GET，则响应错误信息
		logrus.Errorf("%v 请求方法不允许", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
