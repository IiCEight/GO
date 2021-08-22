package main

import (
	"fmt"
	"net"
)

func proc(conn net.Conn) {
	buf := make([]byte, 1024*16)

	n, err := conn.Read(buf)
	checkerr(err)

	fmt.Println(string(buf[:n]))
	conn.Write([]byte("Accepted!!!"))
	conn.Close()
}

func checkerr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	lis, _ := net.Listen("tcp", "127.0.0.1:8080")

	defer lis.Close()

	for {
		conn, err := lis.Accept()
		checkerr(err)
		go proc(conn)
	}

	return
}
