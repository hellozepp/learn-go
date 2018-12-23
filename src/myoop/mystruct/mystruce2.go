package mystruct

import (
	"fmt"
	"unsafe"
)

func Mystruct22() {
	type person struct {
		id   int
		name string
	}
	var p = new(person)
	p.id = 99
	p.name = "xxx"
	var i = uintptr(unsafe.Pointer(p))
	fmt.Println(unsafe.Pointer(p))          //0xc0000465e0
	fmt.Println(i)                          //824634009056
	fmt.Println(uintptr(unsafe.Pointer(p))) //824634009056
	fmt.Println(&i)                         //0xc00004c0e0
	var j = (*int)(unsafe.Pointer(i))

	fmt.Println(*j) //99
}
