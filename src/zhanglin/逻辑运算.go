package zhanglin

import (
	"fmt"
	"runtime"
	"time"
)

func LogicTest() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d , ", i)
	}
	fmt.Println()
	if i := 1 == 1; i {
		fmt.Println(i)
	}

	switch os := runtime.GOOS; os { //支持多语句
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	case "win" + "dows":
		fmt.Println("windows.")
		fallthrough //支持向下穿越
	default: // freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
	//外置条件
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	//defer 语句会延迟函数的执行直到上层函数返回。
	defer call1()
	time.Sleep(time.Second * 1)
	fmt.Println("call finished...")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

}
func call1() {
	fmt.Println("start...")
	fmt.Println("end...")
}
