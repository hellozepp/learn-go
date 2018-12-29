package mysync

import (
	"fmt"
	"sync"
	"time"
)

func MyOnce() {

	o := &sync.Once{}

	go do(o)
	go do(o)
	go do(o)
	go do(o)
	go do(o)
	go do(o)

	time.Sleep(time.Second * 2)

}

func do(o *sync.Once) {

	fmt.Println("Start do")

	o.Do(func() { //只会执行一次
		fmt.Println("Doing something...")
	})

	func() {
		fmt.Println("not once")
	}()

	fmt.Println("Do end")

}
