package myruntime

import (
	"fmt"
	"runtime"
	"time"
)

//由于多个协程同时往stdout写入， 纵然前面执行再快，最后还是要单线程输出到屏幕
var quit chan int = make(chan int)

func loop() {
	for i := 0; i < 10000; i++ {
		fmt.Printf("%d\n", i)
	}

	quit <- 0
}

func NumCpuTest() {
	fmt.Println(runtime.NumCPU()) // 默认CPU核心数
	time.Sleep(2 * time.Second)
	a := 5000
	t0 := time.Now() // 起点时间
	//runtime.GOMAXPROCS(1) // 设置执行使用的核心数

	for i := 1; i <= a; i++ {
		go loop()
	}

	for i := 0; i < a; i++ {
		<-quit
	}

	endTime := time.Since(t0) // 耗时
	fmt.Println("运行时间", endTime)

}
