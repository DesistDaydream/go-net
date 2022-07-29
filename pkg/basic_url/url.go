package main

import (
	"log"
	"net/url"
)

func main() {
	// 解析 URL
	u, err := url.Parse("https://admin:password@example.org/path?foo=bar")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(u.User.Password()) ///path?foo=bar
}
