package myhttp2

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

const (
	_HTTP2URLBase = "https://localhost:9000"
	_CertFile     = "/disk/mygopath/src/iotestgo/myio/mysock/mynet/myhttp2/cert.pem"
	_KeyFile      = "/disk/mygopath/src/iotestgo/myio/mysock/mynet/myhttp2/key.pem"
)

type handlerFunc func(w http.ResponseWriter, r *http.Request)

func MyHttp2Server() {
	httpMux, http2Mux := getHttpMux()
	go httpSrv(httpMux)
	httpsSrv(http2Mux)
}

// Mux定义 -- 设置HTTP1.1访问转向HTTP2
func getHttpMux() (httpMux, http2Mux *http.ServeMux) {

	httpMux = http.NewServeMux()
	http2Mux = http.NewServeMux()

	x := make(map[string]handlerFunc, 0)
	x["/"] = Home
	x["/v1"] = Hello1

	for k, v := range x {
		redirectURL := http.RedirectHandler(_HTTP2URLBase+k, 307)
		httpMux.Handle(k, redirectURL)
		http2Mux.HandleFunc(k, v)
	}

	return
}

//HTTP服务
func httpSrv(mux *http.ServeMux) {
	log.Fatal(http.ListenAndServe(":9001", mux))
}

//HTTP2服务
func httpsSrv(mux *http.ServeMux) {
	srv := &http.Server{
		Addr:         ":9000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      mux,
	}
	http2.VerboseLogs = true
	http2.ConfigureServer(srv, &http2.Server{})

	log.Fatal(srv.ListenAndServeTLS(_CertFile, _KeyFile))
}

//Handler函数
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "RequestURI: %s\n", r.RequestURI)
	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	fmt.Fprintf(w, "Home")
}

func Hello1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "RequestURI: %s\n", r.RequestURI)
	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	fmt.Fprintf(w, "Hello V1")
}
