package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/prometheus/alertmanager/api/v2/models"
)

// Index 基本展示功能。w为响应给客户端的信息。r为客户端发起的请求信息。
func AlertmanagerV1(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("请求体为：%v\n", string(requestBody))

	var postableAlerts models.PostableAlerts

	err := json.Unmarshal(requestBody, &postableAlerts)
	if err != nil {
		fmt.Printf("请检查 Body，格式不正确或数据类型不对\n")
		return
	}

	fmt.Println("当前告警数量：", len(postableAlerts))

	for index, postableAlert := range postableAlerts {
		fmt.Printf("已推送 %v 号推送告警： %v \n", index, postableAlert.Labels["alertname"])
		fmt.Printf("开始时间：%v\n结束时间：%v\n", postableAlert.StartsAt, postableAlert.EndsAt)
	}
}

// Index 基本展示功能。w为响应给客户端的信息。r为客户端发起的请求信息。
func AlertmanagerV2(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("请求体为：%v\n", string(requestBody))

	var postableAlerts models.PostableAlerts

	err := json.Unmarshal(requestBody, &postableAlerts)
	if err != nil {
		fmt.Printf("请检查 Body，格式不正确或数据类型不对\n")
		return
	}

	fmt.Println("当前告警数量：", len(postableAlerts))

	for index, postableAlert := range postableAlerts {
		fmt.Printf("已推送 %v 号推送告警： %v \n", index, postableAlert.Labels["alertname"])
		fmt.Printf("开始时间：%v\n结束时间：%v\n", postableAlert.StartsAt, postableAlert.EndsAt)
	}
}
