package main

import (
	"fmt"
	"time"
)

func odd() {
	for i := 0; i <= 100; i += 10 {
		fmt.Println(i + 1)
		fmt.Println(i + 3)
		fmt.Println(i + 5)
		fmt.Println(i + 7)
		fmt.Println(i + 9)
		time.Sleep(time.Second)
	}
}

func even() {
	for i := 2; i <= 100; i += 10 {
		fmt.Println(i + 0)
		fmt.Println(i + 2)
		fmt.Println(i + 4)
		fmt.Println(i + 6)
		fmt.Println(i + 8)
		time.Sleep(time.Second)
	}
}

func main() {
	go odd()
	go even()

	for {

	}

}
