package main

import (
	"fmt"
	"net"
)

func checkerr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	buf := make([]byte, 1024*16)

	_, err := conn.Write([]byte("Hello!"))
	checkerr(err)
	n, err := conn.Read(buf)
	fmt.Println(string(buf[:n]))

	defer conn.Close()
}
