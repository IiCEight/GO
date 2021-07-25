package main

import "fmt"

const N = 25 + 5

var a [N]int

func main() {
	for i := 1; i < 10; i++ {
		a[i] = i
	}
	for i := 1; i < 10; i++ {
		fmt.Print(a[i], ' ')
	}

	// fmt.Println(a)
}
