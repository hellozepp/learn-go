package mynet

import (
	"io"
	"log"
	"net/http"
)

func MyHttp21() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler21{})
	mux.HandleFunc("/hello", sayhello21)
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}
}

type myHandler21 struct {
}

func (this *myHandler21) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "Hello world!")
}

func sayhello21(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "Hello world! 21")
}
