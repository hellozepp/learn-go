package mygoroutine

import (
	"fmt"
	"time"
)

//无缓冲的channel : chan同步的,不及读取channel，程序就已经阻塞了
//有缓冲的channel : 阻塞直到数据被拷贝到缓冲区
//channel的机制是先进先出
func TestGoroutine() {
	main2()
}
func main1() {

	var i = 3
	go func(a int) {
		fmt.Println(a)
		fmt.Println("1")
	}(i)
	fmt.Println("2")
	time.Sleep(1 * time.Second) //fix:等子线程
}
func main2() { //fatal error: all goroutines are asleep - deadlock!
	//ch := make(chan int)
	ch := make(chan int, 1) //fix:使用有缓冲chan fix2:把ch<-1这一行代码放到子线程代码的后面
	ch <- 1
	go func() {
		<-ch
		fmt.Println("1")
		time.Sleep(2 * time.Second)
		fmt.Println("end sleep")
	}()
	//ch <- 1
	fmt.Println("2")
}
