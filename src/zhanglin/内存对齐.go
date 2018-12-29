package zhanglin

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Data struct {
	b byte  //1
	a int32 //4
	//x int16
}
type Data1 struct {
	b byte  //1
	x int16 //2
}
type Data2 struct {
	b byte //1
	x int  //int默认8byte对齐.
}

func call2() {
	var d Data
	t := reflect.TypeOf(d)
	fmt.Println(t.Size(), t.Align())
}
func call3() {
	var d Data1
	t := reflect.TypeOf(d)
	fmt.Println(t.Size(), t.Align())
}
func call4() {
	var d Data2
	t := reflect.TypeOf(d)
	fmt.Println(t.Size(), t.Align())
}
func AlignTest() {
	ShouldEqual := true
	So(unsafe.Sizeof(true), ShouldEqual, 1)
	So(unsafe.Sizeof(int8(0)), ShouldEqual, 1)
	So(unsafe.Sizeof(int16(0)), ShouldEqual, 2)
	So(unsafe.Sizeof(int32(0)), ShouldEqual, 4)
	So(unsafe.Sizeof(int64(0)), ShouldEqual, 8)
	So(unsafe.Sizeof(int(0)), ShouldEqual, 8)
	So(unsafe.Sizeof(float32(0)), ShouldEqual, 4)
	So(unsafe.Sizeof(float64(0)), ShouldEqual, 8)
	So(unsafe.Sizeof(""), ShouldEqual, 16)
	So(unsafe.Sizeof("hello world"), ShouldEqual, 16)
	So(unsafe.Sizeof([]int{}), ShouldEqual, 24)
	So(unsafe.Sizeof([]int{1, 2, 3}), ShouldEqual, 24)
	So(unsafe.Sizeof([3]int{1, 2, 3}), ShouldEqual, 24)
	So(unsafe.Sizeof(map[string]string{}), ShouldEqual, 8)
	So(unsafe.Sizeof(map[string]string{"1": "one", "2": "two"}), ShouldEqual, 8)
	So(unsafe.Sizeof(struct{}{}), ShouldEqual, 0)
	So(unsafe.Sizeof(struct{ i8 int8 }{}), ShouldEqual, 1)
	// |x---|xxxx|xx--|
	So(unsafe.Sizeof(struct {
		i8  int8 //bit
		i32 int32
		i16 int16
	}{}), ShouldEqual, 12)

	// |x-xx|xxxx|
	So(unsafe.Sizeof(struct {
		i8  int8
		i16 int16
		i32 int32
	}{}), ShouldEqual, 8)
	// |x-xx|
	So(unsafe.Sizeof(struct {
		i8  int8
		i16 int16
	}{}), ShouldEqual, 4)
	// |x---|xxxx|xx--|----|xxxx|xxxx| int64 只能出现在8的倍数的地址处，因此第一个结构体中，有连续的4个字节是空的
	So(unsafe.Sizeof(struct {
		i8  int8
		i32 int32
		i16 int16
		i64 int64
	}{}), ShouldEqual, 24)

	// |x-xx|xxxx|xxxx|xxxx|
	So(unsafe.Sizeof(struct {
		i8  int8
		i16 int16
		i32 int32
		i64 int64
	}{}), ShouldEqual, 16)
	call2()
	call3()
	call4()
}
func So(sizeof uintptr, b bool, bytes int) {
	fmt.Println(sizeof == uintptr(bytes) && b)
}
