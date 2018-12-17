package zhanglin

import (
	"fmt"
	"math"
)

/*函数支持返回多个值*/
func add(x int, y int) (int, int) {
	return x + y, 0
}

/*没有参数的return返回已命名的返回值，下例中已命名返回值即(x int,y int) ，适用于短函数中*/
func split(sum int) (x int, y int) {
	x = sum * 4 / 9
	y = sum - 1
	return
}

/*var可以用来定义全局变量，也可以用在函数中，类似x:=8这种表达方式只能用于函数中。函数外的每个语句都必须以关键字开始(var,func等)*/
var (
	k  int
	kk bool = false
)

/*go语言结构体*/
type Vertex struct {
	X int
	Y int
}

func ff(fn func(int, int) int, b int) int {
	return fn(b, 2*b)
}

func fa(a int, b int) (int, int) {
	return a, b
}

type ffn1 struct {
	fn func(int, int) (int, int)
	a  int
}

/*返回值为函数的函数，一般用于闭包*/
func AAA() func(int, int) (int, string) {
	return func(i int, i2 int) (int, string) {
		return i + i2, "ok"
	}
}

func main() {
	var ii bool = true
	fmt.Println(split(17))
	fmt.Printf("%t\n", ii)

	/*类型推导，由右值推导未定义变量的类型*/
	j := 6.1

	/*表达式T(v)将值v转换为类型T，Go的类型转换需要显示转换，如下也可以为jj:=uint32(j),var jj uint32 = uint32(j)*/
	var jj = uint32(j)
	fmt.Printf("type(j)=%T,type(jj)=%T\n", j, jj)

	/*常量的声明与变量类似，只不过使用const关键字，常量不能使用 := 语法声明，常量的类型是在编译的时候确定的
	如下面例子中的var a float32 = float32(Pi)这一行，如果写成var a int = int(Pi)会报错，但如果Pi是变量的话则没有问题，常量的类型转换
	保证浮点数转浮点数，整数转整数
	无类型的常量不仅可以提供更高的运算精度，而且可以直接用于更多的表达式而不需要显式的类型转换。
	*/
	const Pi = 111.11
	var a float32 = float32(Pi)
	var aa float64 = Pi
	fmt.Printf("%f %f %T\n", a, aa, Pi)

	var in int = 9
	var fl float32 = float32(in)
	fmt.Println(fl)

	var fl1 = 3.1111
	var in1 int = int(fl1)
	fmt.Println(in1)
	/*常量可以用于枚举,每一行的iota都会相比上一行+1，运算方法保持一致*/
	const (
		a1 = iota * 3
		b
		c
		d
		e
	)
	fmt.Println(a1, b, c, d, e)

	/*这种for类型的跟C语言用法一样，只是没有小括号*/
	for j := 0; j < 10; j++ { //写法1
		fmt.Println(j)
		break
	}
	/*这种for的用法可以看做是C语言的while，也可以谢伟for f:=10;f<20;{}*/
	f := 10
	for f < 20 { //写法2
		fmt.Println(11)
		break
	}
	/*这种for的用法可以看做是C语言的while(1)*/
	for { //写法3
		fmt.Println(22)
		break
	}

	/*if语句可以在条件表达式前执行一个简单的语句,该语句声明的变量作用域尽在if之内*/
	v := 6.3
	if v = math.Abs(55); v < 10 { //写法1
		fmt.Println("v<6")
	} else {
		fmt.Println("a>=10")
	}
	if v := 7; v == 19 { //写法2
		fmt.Println(55)
	}
	if v < 10 { //写法3
		fmt.Println(11)
	}
	/*switch是一连串if-else语句的简便方法,go的switch只运行选定的case*/
	switch v := 6; v { //写法1
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("default")
	}

	switch v { //写法2
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("default")
	}

	switch { //写法3，这种形式能将一长串if then else写的更加清晰
	case v < 10:
		fmt.Println(1)
	case v == 10:
		fmt.Println(2)
	default:
		fmt.Println(3)
	}

	/* defer panic recover参见：http://www.cnblogs.com/charlieroro/p/8617056.html */

	/*go的指针用法与C类似，但不同点在于GO没有指针运算*/
	var p *int
	var iii int = 8
	p = &iii
	fmt.Println(*p)

	/*go语言结构体操作*/
	fmt.Println(Vertex{1, 2})
	fmt.Println(Vertex{X: 22, Y: 33}) /*结构体可以指定成员赋值*/
	/*使用点号来访问*/
	var vertex Vertex = Vertex{3, 4}
	fmt.Println(vertex.X, vertex.Y)
	/*结构体指针操作,可以使用(*p_vertex).X格式来访问，也可以使用隐式间接访问p_vertex.X*/
	var p_vertex *Vertex = &vertex
	p_vertex.X = 88
	p_vertex.Y = 99
	fmt.Println(*p_vertex)
	/*结构体赋值类似C语言的memcpy模式，赋值后的两个结构体变量为独立的值*/
	var vertex1 Vertex = vertex
	vertex.X = 199
	fmt.Println(vertex, vertex1)

	/*数组初始化的时候可以忽略[]中的数值，由go判断*/
	arr := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	/*数组切片,它会选择一个半开区间，包括第一个元素，但排除最后一个元素，用法与python一样*/
	var arr2 = arr[:3]
	fmt.Println(arr2)
	/*切片并不存储任何数据， 它只是描述了底层数组中的一段,更改切片的元素会修改其底层数组中对应的元素*/
	name := []string{
		"john",
		"paul",
		"george",
		"ringo",
	}
	var name1 = name[0:2]
	name2 := name[1:3]
	fmt.Println(name, name1, name2)
	name1[0] = "XXX"
	fmt.Println(name, name1, name2)

	/*使用已定义的结构体创建切片，[n]T表示数组，[]T表示切片,数组存储数据，切片只是描述*/
	ffff := []Vertex{{1, 2}, {3, 4}}
	fmt.Println("@@@@@@", ffff)
	/*使用结构体语法创建切片*/
	fff := []struct {
		x int
		y int
	}{
		{1, 2}, {3, 4},
	}
	fmt.Println(fff)
	/*切片的零值是nil*/
	var ss []int
	fmt.Println(ss, len(ss), cap(ss))
	/*使用make创建切片,切片的startindex会影响切片的容量，但endindex则不会影响切片的容量*/
	sl := make([]int, 5, 5)
	fmt.Println(cap(sl[:2]), cap(sl[2:]))
	/*多维切片*/
	sl1 := [][]int{
		[]int{1, 2, 3},
		[]int{3, 4, 5, 4},
	}
	fmt.Println(sl1)
	/*使用append向切片追加元素*/
	var inn []int = make([]int, 5, 5)
	fmt.Println(inn, cap(inn))
	inn = append(inn, 4, 5, 6)
	fmt.Println(inn, cap(inn))
	/*for循环的range可以遍历切片或映射*/
	for k, v := range inn {
		fmt.Println(k, v)
	}
	/*可以使用下标 _ 来忽略返回值*/
	for _, v := range inn {
		fmt.Println(v)
	}

	/*映射,使用make创建一个map*/
	type VV struct {
		a, b, c int
	}
	mapp := make(map[string]VV) //映射定义1
	mapp["hello"] = VV{3, 4, 5}
	fmt.Println(mapp)
	/*映射的语法与结构体相似，不过必须有键名*/
	var m = map[string]VV{ //映射定义2
		"one": VV{1, 2, 3},
		"two": VV{2, 3, 4},
	}
	mm := map[string]VV{ //映射定义3
		"one": VV{1, 2, 3},
		"two": VV{2, 3, 4},
	}
	/*若顶级类型只是一个类型名，可以省略*/
	mmm := map[string]VV{ //映射定义3
		"one": {1, 2, 3},
		"two": {2, 3, 4},
	}
	fmt.Println(m, mm, mmm)
	/*插入元素*/
	mmm["three"] = VV{3, 4, 5}
	fmt.Println(mmm)
	/*删除元素*/
	delete(mmm, "two")
	fmt.Println(mmm)
	/*检测某个键是否存在,如下ok==true时表示存在，ok==false时表示不存在*/
	if elem, ok := mmm["three"]; ok {
		fmt.Println(elem, ok)
	}

	/*函数也是值，可以像其他值一样传递,函数作为函数入参*/
	fmt.Println(ff(func(i int, i2 int) int {
		return i + i2
	}, 10))
	/*函数作为结构体参数*/
	ffnn := ffn1{
		fa,
		100,
	}
	fmt.Println(ffnn.fn(1, 2))
}
