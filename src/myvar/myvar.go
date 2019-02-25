package myvar

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

var (
	a1, _, c1 = 1, 2, 3
)

func Myvar() {

	var i1 int32 = 123
	fmt.Println(i1)

	var i2 = 123
	fmt.Println(i2)

	i3 := 123
	fmt.Println(i3)

	fmt.Println("========================1=======================")
	//========================================================================

	var a, _, c int = 1, 2, 3
	fmt.Println(a, c)
	fmt.Println(a1, c1)
	fmt.Println("========================2=======================")
	//========================================================================
	var j = 0x16
	fmt.Println("hello " + strconv.Itoa(j))
	j1, _ := strconv.Atoi("1.1") // Array to Integer
	fmt.Println(j1)
	fmt.Println(int(j))

	fmt.Println("========================3=======================")
	//========================================================================
	type size int //类型别名
	var s size = 12
	fmt.Println(s)

	fmt.Println("========================4=======================")
	//========================================================================
	var f, g = 1, "abc"
	fmt.Println(f, g)

	fmt.Println("========================5=======================")
	//========================================================================
	const z string = "abc" //不能类型推断
	//zz:="abc"
	fmt.Println(z)
	const (
		z1 = "abc"
		z2 = len(z1) //等同于写死３
		z3 = 123
		//z4=len(zz) //由于zz不是常量，编译时无法取得立即数，所以无法在编译的时候计算len，因此无法为常量赋值
		//z ="" //不能重新定义
	)
	fmt.Println(z1, z2, z3) //abc 3 123
	const (
		k1 = iota //希腊字母。golang中定义常量经常用的iota关键字 默认递增值0开始 go语言中的枚举
		k2
		k3 = 100
		k4
		k5 = iota //如果中断iota，必须显式恢复 4
		k6        //5
	)
	const (
		k11 = iota //括号外不连续
		k21
		k31
		k41
	)
	fmt.Println(k1, k2, k3, k4, k5, k6) //0 1 100 100 4 5
	fmt.Println(k11, k21, k31, k41)     //0 1 2 3
	// 1.iota常量自动生成器,每隔一行,自动累加1
	// 2.iota给常量赋值用
	const (
		aaa = iota //0
		b   = iota //1
		cc  = iota //2
	)
	fmt.Printf("a = %d, b = %d, c = %d\n", aaa, b, cc)

	// 3.iota遇到const,重置为0
	const d = iota
	fmt.Printf("d = %d\n", d)
	const dd = iota //不连续 0
	fmt.Printf("dd = %d\n", dd)

	// 4 可以只写一个iota
	const (
		a11 = iota //0
		b1         //1
		c1         //2
	)
	fmt.Printf("a1 = %d, b1 = %d, c1 = %d\n", a11, b1, c1)

	// 5 如果是同一行,值都一样
	const (
		i           = iota
		j11, j2, j3 = iota, iota, iota
		k           = iota
	)

	fmt.Printf("i = %d, j1 = %d, j2 = %d, j3 = %d, k = %d\n", i, j11, j2, j3, k)

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays // 这个常量没有导出
	)
	fmt.Println("========================6=======================")
	//========================================================================

	/*

				二进制   十进制

		            01100   12

		            10110   22

				-----------------------

				二元位运算符  二进制   十进制    逻辑

					&       00100   4        相同位的两个数字­都为1，则为1；若有一个不为1，则为0。

					|       11110   30       相同位只要一个为1即为1。否则为0

					^       11010   26       相同位不同为1则为1，否则为0。

					&^      01000   8        如果第二个数是1，则强制把第一个数改成0，否则不变。

	*/
	fmt.Println(12 & 22)
	fmt.Println(12 | 22)
	fmt.Println(12 ^ 22)
	fmt.Println(12 &^ 22)

	fmt.Println("========================7=======================")
	//========================================================================
	f1, _ := strconv.ParseFloat("1.234", 64) //1.234
	fmt.Println(f1)

	ii, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(ii)

	ddd, _ := strconv.ParseInt("0x1c8", 0, 64) //456
	fmt.Println(ddd)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	kk, _ := strconv.Atoi("135")
	fmt.Println(kk)

	_, e := strconv.Atoi("wat") //strconv.Atoi: parsing "wat": invalid syntax
	fmt.Println(e)
	fmt.Println("========================8=======================")
	//========================================================================

	var aint int = 111
	var aa *int = &aint
	fmt.Println(aa, *aa, *&aa)

	var a1 int16 = 0x0102
	var i22 = (*int8)(unsafe.Pointer(&a1))
	fmt.Printf("%x \n", *i22) //2
	var i222 = (*int)(unsafe.Pointer(&a1))
	fmt.Printf("0x%4x \n", *i222) //258
	var p1 = (int64)(uintptr(unsafe.Pointer(i22)))
	fmt.Println(*(*int8)(unsafe.Pointer(uintptr(p1 + 1)))) //1

	fmt.Println("========================9=======================")
	//========================================================================
	type myint int //自定义类型
	var m1 myint = 100
	fmt.Println(m1, reflect.TypeOf(m1), reflect.TypeOf(m1).Kind())

	type myint2 = int //类型别名
	var m2 myint2 = 100
	fmt.Println(m2, reflect.TypeOf(m2), reflect.TypeOf(m2).Kind())
}
