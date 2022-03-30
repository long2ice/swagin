package main

import (
	"net"
	"os/exec"
	"testing"
	"time"
)

func TestIntegration(t *testing.T) {
	go main()
	WaitServer()
	out, err := exec.Command("venom", "run", "test/*.yml").CombinedOutput()
	if err != nil {
		t.Fatalf("running venom: %s", string(out))
	}
}

func WaitServer() {
	timeout := 10 * time.Millisecond
	for {
		conn, err := net.DialTimeout("tcp", "127.0.0.1:8080", timeout)
		if err == nil {
			conn.Close()
			return
		}
		time.Sleep(timeout)
	}
}
