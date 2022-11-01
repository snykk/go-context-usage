package main

import (
	"fmt"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	time.Sleep(5 * time.Second) // Simulasi proses lambat
	conn.Write([]byte("Hello from server"))
}

func main() {
	listener, _ := net.Listen("tcp", ":8080")
	defer listener.Close()

	fmt.Println("Server is running on port 8080...")
	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}
