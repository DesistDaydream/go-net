package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Download 下载文件
func Download(w http.ResponseWriter, r *http.Request) {
	filename := "./templates/file"
	file, err := os.Open(filename)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("打开文件错误：", err)
		_, _ = io.WriteString(w, "打开文件错误")
		return
	}
	defer file.Close()

	_, err = io.Copy(w, file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "拷贝文件错误")
		return
	}
}
