package mymethod

import (
	"fmt"
	"runtime"
	"time"
)

type Data struct {
	Name string
	d    [1024 * 100]byte
}

func close1(d *Data) {
	fmt.Println("被回收了！！", d.Name)
}
func test() {
	var a Data
	a.Name = "xxx"
	//设置析构函数
	runtime.SetFinalizer(&a, close1)
}
func Finallizefunc() {
	for {
		test()
		time.Sleep(time.Millisecond)
	}
}
