package main

import (
	"fmt"
	"net/http"
)

func mainPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello!!!"))
}
func apiPage(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Api Page!"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/", apiPage)
	mux.HandleFunc("/", mainPage)
	err := http.ListenAndServe(":8080", mux)
	fmt.Println("Hello")
	if err != nil {
		panic(err)
	}
}
