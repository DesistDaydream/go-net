package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

// 定义容器列表响应的结构体
type Container struct {
	ID      string   `json:"Id"`
	Names   []string `json:"Names"`
	Image   string   `json:"Image"`
	Status  string   `json:"Status"`
	Command string   `json:"Command"`
	Created int64    `json:"Created"`
}

// 使用 SOCK_STREAM 类型创建 Unix socket 连接
func SOCK_STREAM() {
	// 创建一个自定义的 HTTP 客户端，使用 Unix socket 连接
	httpClient := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/var/run/docker.sock")
			},
		},
		Timeout: time.Second * 30,
	}

	// 示例1: 获取 Docker 版本信息
	fmt.Println("1. 获取 Docker 版本信息:")
	versionResp, err := httpClient.Get("http://localhost/version")
	if err != nil {
		fmt.Printf("无法获取 Docker 版本: %v\n", err)
		os.Exit(1)
	}
	defer versionResp.Body.Close()

	var version map[string]interface{}
	if err := json.NewDecoder(versionResp.Body).Decode(&version); err != nil {
		fmt.Printf("无法解析版本响应: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Docker 版本: %s\n", version["Version"])
	fmt.Printf("API 版本: %s\n", version["ApiVersion"])
	fmt.Printf("Go 版本: %s\n", version["GoVersion"])
	fmt.Println()

	// 示例2: 列出运行中的容器
	fmt.Println("2. 列出运行中的容器:")
	containersResp, err := httpClient.Get("http://localhost/containers/json")
	if err != nil {
		fmt.Printf("无法列出容器: %v\n", err)
		os.Exit(1)
	}
	defer containersResp.Body.Close()

	var containers []Container
	if err := json.NewDecoder(containersResp.Body).Decode(&containers); err != nil {
		fmt.Printf("无法解析容器响应: %v\n", err)
		os.Exit(1)
	}

	if len(containers) == 0 {
		fmt.Println("当前没有运行中的容器")
	} else {
		fmt.Println("运行中的容器:")
		for _, container := range containers {
			fmt.Printf("ID: %s\n", container.ID[:12])
			fmt.Printf("名称: %s\n", container.Names[0][1:]) // 移除开头的 '/'
			fmt.Printf("镜像: %s\n", container.Image)
			fmt.Printf("状态: %s\n", container.Status)
			fmt.Println("---")
		}
	}
	fmt.Println()

	// 示例3: 获取系统信息
	fmt.Println("3. 获取 Docker 系统信息:")
	infoResp, err := httpClient.Get("http://localhost/info")
	if err != nil {
		fmt.Printf("无法获取系统信息: %v\n", err)
		os.Exit(1)
	}
	defer infoResp.Body.Close()

	body, err := io.ReadAll(infoResp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		os.Exit(1)
	}

	var info map[string]interface{}
	if err := json.Unmarshal(body, &info); err != nil {
		fmt.Printf("无法解析系统信息: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("容器数量: %v\n", info["Containers"])
	fmt.Printf("镜像数量: %v\n", info["Images"])
	fmt.Printf("Docker Root Dir: %v\n", info["DockerRootDir"])
	fmt.Printf("操作系统: %v\n", info["OperatingSystem"])
	fmt.Printf("架构: %v\n", info["Architecture"])
	fmt.Printf("内核版本: %v\n", info["KernelVersion"])
}

// 使用 SOCK_DGRAM 类型创建 Unix socket 连接
func SOCK_DGRAM() {}

// 使用 SOCK_SEQPACKET 类型创建 Unix socket 连接。
// 参考 github.com/DesistDaydream/go-dpdk 项目
func SOCK_SEQPACKET() {}
