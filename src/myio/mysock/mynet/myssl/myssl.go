package myssl

import (
	"io"
	"log"
	"net/http"
)

//openssl genrsa -out key.pem 2048
//openssl req -new -x509 -key key.pem -out cert.pem -days 3650

func MySsl() {
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServeTLS(":8080", "/disk/mygopath/src/iotestgo/myio/mysock/mynet/myssl/cert.pem", "/disk/mygopath/src/iotestgo/myio/mysock/mynet/myssl/key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS:", err.Error())
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world!")
}
