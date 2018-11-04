package mynet

import (
	"io"
	"log"
	"net/http"
)

func MyHttp1() {
	http.HandleFunc("/", sayhello1)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sayhello1(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "Hello world!")
}
