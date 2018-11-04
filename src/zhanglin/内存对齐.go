package zhanglin

import (
	"fmt"
	"reflect"
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
func AlignTest() {
	call2()
	call3()
}
