package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Test is
func Test(w http.ResponseWriter, req *http.Request) {
	// 允许跨域访问
	w.Header().Set("Access-Control-Allow-Origin", "*")

	logrus.WithField("method", req.Method).Infof("请求方法")
	logrus.WithField("url", req.URL).Info("URL")
	for k, v := range req.Header {
		logrus.WithFields(logrus.Fields{
			k: v,
		}).Infof("请求头")
	}
	// logrus.WithField("header", req.Header).Infof("请求头")
	body, _ := io.ReadAll(req.Body)
	logrus.WithField("body", string(body)).Infof("请求体")
	fmt.Fprintf(w, "测试页面!")
}
