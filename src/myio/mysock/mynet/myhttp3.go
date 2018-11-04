package mynet

import (
	"io"
	"log"
	"net/http"
	"time"
)

func MyHttp3() {
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler3{},
		ReadTimeout: 5 * time.Second,
	}
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

type myHandler3 struct {
}

func (this *myHandler3) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "Hello world!")
}
