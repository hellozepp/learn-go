package myfunc

import "fmt"

/**
defer的用法遵循3个原则
1 在defer表达式被运算的同时，defer函数的参数也会被运算。立即运算
2 defer函数在一个函数return之后遵循后进先出的调用原则
3 defer函数可能会读取并赋值给所在函数的返回值
*/

func Mydefer() {

	fmt.Println(defertest1())
	fmt.Println("========================1=======================")
	defertest()
	fmt.Println("========================2=======================")
	fmt.Println(c())
	fmt.Println("========================3=======================")
	fmt.Println(f())
	fmt.Println("f1:", f1())
}
func defertest1() int { //defer不会调用，仅仅会压入本地方法栈中，放在：｛所有栈变量之后，ret之前｝先进后出
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
	return 1
}

func defertest() {
	for i := 0; i < 2; i++ {
		defer fmt.Println(i)
	}

	defer fmt.Println("======")

	for i := 0; i < 2; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
	//分析执行顺序:
	//1.栈中压入fmt.Println(0)
	//2.栈中压入fmt.Println(1)
	//3.栈中压入fmt.Println("======")
	//4.栈中压入func() { fmt.Println(i) }()//此时i==0
	//5.栈中压入func() { fmt.Println(i) }()//此时i==1
	//6.i==2退出循环，此时i==2
	//7.函数结束，开始弹出栈
	//8.执行5，打印i值，此时i==2
	//9.执行4，打印i值，此时i==2
	//10.执行3，打印======
	//11.执行2，打印1
	//12.执行1，打印0
}
func c() (i int) {
	defer func() {
		fmt.Println("i:", i) ////100
	}()
	return 100
}
func f() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}
func f1() (r int) {
	defer func() { //该返回值,在return之后
		r = r + 5
	}()
	return 1
}
func f2() (r int) {
	r = 1               //给返回值赋值
	defer func(r int) { //这里改的r是传值传进去的r，不会改变要返回的那个r值
		r = r + 5
		fmt.Println(r)
	}(r)
	return //空的return
}
