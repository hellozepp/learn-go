package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("hello")
	fmt.Println("hello")
	time.Sleep(time.Second)
}
