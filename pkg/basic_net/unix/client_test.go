package main

import "testing"

func TestSOCK_STREAM(t *testing.T) {
	// go test -timeout 1m -v -run TestSOCK_STREAM pkg/basic_net/unix/*.go
	SOCK_STREAM()
}
