package zhanglin

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type ffn struct {
	fn func(int, int) (int, int)
	a  int
}

/*ffn的方法定义*/
func (ffnn *ffn) add(a int) (int, int) {
	return a, a + 1
}

/*这是一个正常的函数，非方法
func add(ffnn ffn)(int,int){
	return 1,a+1
}
*/
func fa1(a int, b int) (int, int) {
	return a, b
}

type FLOAT64 float64

func (fl64 FLOAT64) say(s string) string {
	return s
}
func (fl64 *FLOAT64) say2(s string) string {
	return s
}

func (fl *ffn) say2(s string) int {
	fl.a = 1000
	return fl.a
}

/*指针函数*/
func AA(fl *FLOAT64) FLOAT64 {
	return *fl
}

/*type类似C语言的typedef，使用type可以简化类型定义*/
type INT_ARR []int

type I interface {
	add(a int, b int) interface{}
	sub(c int, d int) int
}
type Test struct {
	a int
	b int
}

/*实现方法的时候，一般传入类的指针*/
func (t *Test) add(a int, b int) interface{} { /*返回值为空接口，可以返回多种类型，类似*/
	if nil == t { /*处理入参为空指针的情况*/
		fmt.Println("nil")
		return "error:nil"
	}
	return a + b + t.a + t.b
}
func (t *Test) sub(a int, b int) int {
	if nil == t {
		return -1
	}
	return a - b + t.a + t.b
}

/*类型判断*/
func X(a int) interface{} {
	if a = a + 10; a > 20 {
		return a
	} else {
		return "a<=20"
	}
}

/*类型选择，type为关键字*/
func S(a interface{}) {
	switch t := a.(type) {
	case int:
		fmt.Println(t)
	case string:
		fmt.Println(t + "hello")
	default:
		fmt.Println("unkonwn type")
	}
}

func CH(s string, c chan string) {
	c <- s + "1" /*存放管道的时候需要在一个goroutine中进行，否则会多线程乱序*/
	c <- s + "2"
	c <- s + "3"
	c <- s + "4"
	close(c)
}

/*返回值类型为chan string*/
func CH1(s string) chan string {
	c := make(chan string)
	go func() {
		for i := 1; i < 10; i++ {
			c <- s + strconv.Itoa(i)
		}
		close(c)
	}()
	return c
}

/*sync.Mutex使用*/
type SF struct {
	v   map[string]int
	mux sync.Mutex
}

func (sf *SF) inc() {
	(*sf).mux.Lock()
	(*sf).v["hello"]++
	(*sf).mux.Unlock()
}

//只能向chan里写数据
func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
}

//只能取channel中的数据
func recv(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
}

func main() {
	/*方法，结构体作为方法接收者，为结构体增加方法，可以看做增加结构体特有的功能；不能为内建类型声明方法*/
	fmn := &ffn{
		fa1,
		200,
	}
	fmt.Println(fmn.add(3))
	/*只能为在同一包内定义的类型的接收者声明方法， 而不能为其它包内定义的类型（包括 int 之类的内建类型）的接收者声明方法;
	  以指针为接收者的方法被调用时，接收者既能为值又能为指针，如下fl64t.say2其实被解析为(&fl64t).say2
	  以值为接收者的方法被调用时，接收者既能为值又能为指针，(&fl64t).say被解析为(*(&fl64t)).say
	  总结：方法接收者可以为值或是指针，go会自动解析。一般使用指针接收者，避免每次调用时复制该值
	*/
	fl64t := FLOAT64(1.1)
	fmt.Println(fl64t.say("hello"))
	fmt.Println(fl64t.say2("hello"))
	fmt.Println((&fl64t).say("hello"))
	fmt.Println((&fl64t).say2("hello"))
	/*指针作为方法接收者可以修改接收者指向的值*/
	fmn.say2("hi")
	fmt.Println(fmn) /*fmn.a由200变为1000*/

	/*指针函数需要传入一个指针类型*/
	fl1 := FLOAT64(1.10)
	fmt.Println(AA(&fl1))

	test := Test{5, 3}
	/*因为方法的接收体为指针，接口赋值为结构体的地址。如果方法的接收体全部为结构体，则接口赋值可以为指针也可以为结构体，
	有了结构体类型作为入参的函数，go语言会自动创建一个相应的指针作为入参的函数，但是只有指针入参的函数，不会创建相应的结构体类型的函数
	方法的接收体最好为指针*/
	var interf I = &test
	fmt.Println(interf.add(2, 1))
	fmt.Println(interf.sub(2, 1))
	var test1 *Test /*未初始化的指针为nil，未初始化的结构体的成员则会赋默认值0 或 nil。nil接口既不保存值也不保存具体类型*/
	var interf1 I = test1
	fmt.Println(interf1.add(2, 1))
	fmt.Println(interf1.sub(2, 1))

	/*空接口可以保存任意类型的值，因为每个类型至少实现了0个方法，用来处理未知类型的值*/
	var intef interface{}
	intef = 1
	fmt.Printf("%T\n", intef)
	intef = "hello"
	fmt.Printf("%T\n", intef)

	/*类型断言提供了访问接口值底层具体指的方式，类型断言写法为 x.(T)，x为interface{}类型，T就是要断言的类型*/
	if res, ok := X(5).(string); ok {
		fmt.Println(res)
	} else {
		fmt.Println("ok:", res)
	}
	type GG struct {
		a int
	}
	var g interface{}
	g = GG{3}
	gg, ggg := g.(GG)
	fmt.Println(gg, ggg)
	/*类型选择*/
	S(19)
	S("hi ")

	/*信道，默认情况下，发送和接收端在另一端准备好之前都会阻塞，这使得go可以在没有显示锁情况下同步；
	启动的所有goroutine里的非缓冲信道一定要一个线里存数据，一个线里取数据，要成对才行*/
	c := make(chan string)
	go CH("hello-liu", c)
	x, y, z, k := <-c, <-c, <-c, <-c //在main goroutine中期望从管道中获得一个数据，不带缓存时这个数据必须是其他goroutine中放入管道的,否则会导致发端阻塞main goroutine
	fmt.Println(x, y, z, k)
	/*如下会阻塞go func，但因为没有阻塞所有goroutine，所以不会发生错误，运行后可以看到因为没有读操作，该goroutine被阻塞了*/
	v := make(chan int)
	go func() {
		fmt.Println("yyyyyyyyyyyyyyyyyyyyyyy1")
		v <- 100
		fmt.Println("yyyyyyyyyyyyyyyyyyyyyyy2")
	}()
	/*由于没有写，下述goroutine也被阻塞了*/
	vv := make(chan int)
	go func() {
		fmt.Println("yyyyyyyyyyyyyyyyyyyyyyy3")
		fmt.Println(<-vv)
		fmt.Println("yyyyyyyyyyyyyyyyyyyyyyy4")
	}()
	/*带缓存的信道，仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。第二个参数是指缓存的个数*/
	cc := make(chan string, 2) /*如果不带缓存，下一行放入数据的时候会导致主goroutine阻塞，报错fatal error: all goroutines are asleep - deadlock!*/
	cc <- "j000000000000000000o"
	cc <- "j000000000000000000o"
	close(cc)               /*注意：此函数需要发送者关闭channel,表示没有需要发送的值,信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不再有值需要发送的时候才有必要关闭，例如终止一个 range 循环*/
	fmt.Println(<-cc, <-cc) /*带缓存且信道内有数据时，自发自收不会导致阻塞*/
	//fmt.Println(<-cc)/*如果缓存区内没有数据，此时接收方会阻塞，进而产生错误*/

	/*使用只读只写channel*/
	c3 := make(chan int)
	go send(c3)
	go recv(c3)
	time.Sleep(2 * time.Second)
	/*定义只读和只写的channel意义不大，一般用于在参数传递中*/
	q := make(chan int)
	q1 := chan<- int(q)
	fmt.Printf("%T\n", q1)

	/*使用range来不断从信道接收值，直到其关闭*/
	ccc := make(chan string)
	go CH("haha", ccc)
	for i := range ccc { /*如果CH中没有close，会导致此处循环无法停止，它需要一个关闭标记*/
		fmt.Println(i)
	}
	/*接收者可以通过为接收表达式分配第二个参数来测试信道是否关闭*/
	rrr, ttt := <-ccc
	fmt.Println(rrr, ttt)

	/*select机制*/
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		ch1 <- "1"
		ch1 <- "11"
		ch1 <- "111"
		ch1 <- "1111"
		ch1 <- "11111"
		ch2 <- "quit!!!" /*使用ch2传递一个结束信号*/
	}()

	for { /*使用for循环可以保证它的case之一能够在执行前发生阻塞,没有for循环知会进行select判定，类似switch*/
		select {
		case c := <-ch1: /*如果无需取信道中数据，则可以使用case <-ch1:*/
			fmt.Println("ch1", c)
			break
		case qu := <-ch2:
			fmt.Println(qu)
			goto AA /*使用break只能跳出select*/
		default: /*为了自尝试发送或接收时不发生阻塞，可以使用default分支*/
			continue
		}
	}

AA:
	ch := CH1("i am here")
	xx, yy, zz := <-ch, <-ch, <-ch
	fmt.Println(xx, yy, zz)
	fmt.Println("00000000000000000000")
	for i := range ch {
		fmt.Println(i)
	}

	vv1 := make(chan int)
	close(vv1) /*读取关闭的channel是不会阻塞的，如果没有这行，下面一行会报错fatal error: all goroutines are asleep - deadlock!*/
	xxx := <-vv1
	fmt.Println(xxx)

	/*sync.Mutex,多goroutine下的资源访问*/
	sf := SF{v: make(map[string]int)}
	for i := 1; i < 100; i++ {
		go sf.inc()
	}
	time.Sleep(2)
	fmt.Println("sf.v=", sf.v["hello"])

}
