package mynet

import (
	"io"
	"log"
	"net/http"
)

func MyHttp2() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct {
}

func (this *myHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "Hello world!")
}
