package mynet

import (
	"io"
	"log"
	"net/http"
	"time"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

func MyHttp31() {
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler31{},
		ReadTimeout: 5 * time.Second,
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/hello"] = sayhello31

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

type myHandler31 struct {
}

func (this *myHandler31) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if h, ok := mux[req.URL.String()]; ok {
		h(resp, req)
		return
	}
}

func sayhello31(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "Hello world! 31")
}
