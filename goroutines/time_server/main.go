package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

// To Run this,
// "go run main.go"
// open new terminal and execute "nc localhost 9000"
func main() {
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go func() {
			for {
				_, err := io.WriteString(conn, time.Now().Format("15:05:05\n"))
				if err != nil {
					return
				}
				time.Sleep(1 * time.Second)
			}
		}()
	}
}
