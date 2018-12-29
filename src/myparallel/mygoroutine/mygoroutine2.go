package mygoroutine

import (
	"fmt"
	"time"
)

func Mygoroutine2() {
	//crw := make(chan int) //无缓冲的chan是阻塞的
	crw := make(chan int, 10000)
	go send1(crw)
	go recv2(crw)
	//go send3(crw)
	//go recv4(crw)
	time.Sleep(6 * time.Second)

}

func send1(c chan<- int) {
	fmt.Println("send send...")
	for i := 0; i < 10000; i++ {
		c <- i
	}
	fmt.Println("send is ok!")
}

func recv2(c <-chan int) {
	fmt.Println("start receive ...")
	for i := range c {
		fmt.Println(i)
		if i == 10 {
			break
		}
	}
	_, ok := <-c
	fmt.Println("receive is ok ! c:", ok)
}
func send3(c chan<- int) {
	fmt.Println("send send...")
	for i := 0; i < 10000; i++ {
		c <- i
	}
	fmt.Println("send is ok!")
	close(c) //当一个chanel被关闭后，再取出不会阻塞，而是返回零值
}

func recv4(c <-chan int) {
	fmt.Println("start receive ...")
	for i := range c {
		if i < 10 {
			fmt.Println(i)
		}
	}
	fmt.Println("receive is ok !")
}
