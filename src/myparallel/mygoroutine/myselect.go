package mygoroutine

import (
	"fmt"
	"strconv"
	"time"
)

func Myselect() {
	selectTest()
	fmt.Println()
	c1 := make(chan int)
	go func(c1 chan int) {
		for v := range c1 {
			fmt.Println(v)
		}
	}(c1)

	for i := 0; i < 5; i++ { //有点订阅发布的意思，如果之前没有订阅就会死锁
		select {
		case c1 <- 0:
		case c1 <- 1:
		}
		fmt.Println("case done!", time.Stamp)
	}

	fmt.Println("========================1=======================")
	//========================================================================
	c11 := make(chan int)
	c12 := make(chan int)
	go func() {
		c11 <- 1
	}()
	go func() {
		c12 <- 2
	}()
	for i := 0; i < 2; i++ {
		select {
		case j := <-c11:
			fmt.Println("c11: " + strconv.Itoa(j))
		case j := <-c12:
			fmt.Println("c12: " + strconv.Itoa(j))
		}
	}

	fmt.Println("========================2=======================")
	//========================================================================
	a := make(chan string)
	select {
	case str := <-a:
		fmt.Println(str)
	default:
		fmt.Println("非阻塞") //否则就是死锁错误
	}
	select {
	case a <- "xxx":
		fmt.Println("发送了xxx")
	default:
		fmt.Println("非阻塞") //否则就是死锁错误
	}

}

/*
每个case都必须是一个通信
所有channel表达式都会被求值
所有被发送的表达式都会被求值
如果任意某个通信可以进行，它就执行；其他被忽略。
如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
否则：
如果有default子句，则执行该语句。
如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。

*/
func selectTest() {
	var c1, c2 chan int
	var c3 = make(chan int, 1)
	c3 <- 3
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
