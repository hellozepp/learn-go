package myexception

import (
	"fmt"
	"time"
)

func Panic() {
	fmt.Println("----------------start----------------")
	f()
	fmt.Println("----------------end----------------")
}

func f() {
	defer func() {
		//必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("defer start")
		if err := recover(); err != nil {
			fmt.Println(err) //这里的err其实就是panic传入的内容，"this is bug"

		}
		fmt.Println("defer end")

	}()
	for {
		fmt.Println("func begin")
		a := []string{"a", "b"}
		//panic("this is bug")
		fmt.Println("func end") // 不会运行的.
		fmt.Println(a[3])       // 越界访问，肯定出现异常
		time.Sleep(time.Second * 1)
		//panic("bug")            // 上面已经出现异常了,所以肯定走不到这里了。

	}
}
