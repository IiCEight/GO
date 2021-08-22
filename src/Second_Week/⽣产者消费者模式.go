package main

import (
	"fmt"
	"math/rand"
	"time"
)

const mod int = 1e9 + 7

func produce(queue chan int) {
	for {
		queue <- rand.Int()
		time.Sleep(time.Second)
	}

	close(queue)
}

func receive(queue chan int, exit chan bool) {
	for v := range queue {
		fmt.Println("Receive successfully: ", (v % mod))
	}
	exit <- true
	close(exit)
}

func main() {
	queue := make(chan int, 10)
	exit := make(chan bool, 1)
	go produce(queue)
	go receive(queue, exit)
	<-exit

}
