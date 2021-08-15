package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
)

func checkerr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func up(res http.ResponseWriter, r *http.Request) {

	page, err := template.ParseFiles("D:\\Project\\Visual Studio Code\\Go\\src\\Fourth Week\\load.html")

	checkerr(err)

	if r.Method == "POST" {
		fmt.Println("Accepted!")
		r.ParseMultipartForm(1 << 50)
		f, h, err := r.FormFile("myfile")
		checkerr(err)
		// fmt.Println(string(h))
		nf, err := os.Create("D:\\Project\\Visual Studio Code\\Go\\src\\Fourth Week/" + h.Filename)
		checkerr(err)
		n, err := io.Copy(nf, f)
		checkerr(err)
		fmt.Println(n)
	}

	page.Execute(res, nil)

}

func down(res http.ResponseWriter, r *http.Request) {

	page, err := template.ParseFiles("D:\\Project\\Visual Studio Code\\Go\\src\\Fourth Week\\load.html")
	checkerr(err)
	if r.Method == "POST" {
		r.ParseMultipartForm(1 << 50)
		file := r.Form["down"]
		path := "D:\\Project\\Visual Studio Code\\Go\\src\\Fourth Week"
		http.ServeFile(res, r, path+file[0])
	}
	page.Execute(res, nil)

}

func main() {

	http.HandleFunc("/up", up)
	http.HandleFunc("/down", down)
	http.ListenAndServe(":8080", nil)
}
