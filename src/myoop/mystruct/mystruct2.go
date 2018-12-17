package mystruct

import (
	"fmt"
	"unsafe"
)

func Mystruct2() {
	type person struct {
		id   int
		name string
	}
	var p = new(person)
	p.id = 99
	p.name = "xxx"
	var i = uintptr(unsafe.Pointer(p))
	var j = (*int)(unsafe.Pointer(uintptr(i)))
	fmt.Println(*j)

}
