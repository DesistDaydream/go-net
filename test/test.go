package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	dialer := &net.Dialer{
		// 在这里设置 tcp 的 keepalive 和 timeout 时间
		Timeout:   60 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}

	transport := &http.Transport{
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
	}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				r, err := client.Get("http://172.19.42.248:9090")
				if err != nil {
					fmt.Println(err)
					return
				}
				_, err = io.ReadAll(r.Body)
				r.Body.Close()
				if err != nil {
					fmt.Println(err)
					return
				}
				time.Sleep(30 * time.Millisecond)
			}
		}()
	}
	wg.Wait()
}
