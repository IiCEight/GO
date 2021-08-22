package main

import (
	"fmt"
	"net"
)

func main() {
	listen, _ := net.Listen("tcp", ":8080")
	for {
		conn, _ := listen.Accept()
		buf := make([]byte, 1024)
		l, _ := conn.Read(buf)
		fmt.Println(string(buf[:l]))
		// conn.Write([]byte("Hello World"))
		conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain;charset=UTF-8\r\n\r\nHello World"))
		fmt.Println("Accepted!")
		conn.Close()
	}
}
