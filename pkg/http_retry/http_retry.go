package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func retryDo() {
	originalBody := []byte("abcdefghigklmnopqrst")
	reader := strings.NewReader(string(originalBody))
	req, _ := http.NewRequest("POST", "http://localhost:8090/", reader)
	client := http.Client{
		Timeout: time.Second * 5,
	}

	for {
		_, err := client.Do(req)
		if err != nil {
			fmt.Printf("error sending the first time: %v\n", err)
		}
		time.Sleep(1 * time.Second)
	}
}
func main() {
	go func() {
		http.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(time.Second * 10)
			body, _ := ioutil.ReadAll(r.Body)
			fmt.Printf("received body with length %v containing: %v\n", len(body), string(body))
			w.WriteHeader(http.StatusOK)
		}))
		http.ListenAndServe(":8090", nil)
	}()
	fmt.Print("Try with bare strings.Reader\n")
	retryDo()
}
