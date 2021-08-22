package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func checkerr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func index(resp http.ResponseWriter, r *http.Request) {
	f, err := os.Open("D:\\Project\\Visual Studio Code\\web\\Mywebsite\\index.html")
	checkerr(err)
	io.Copy(resp, f)
	f.Close()
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
