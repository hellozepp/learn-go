package a

import (
	"fmt"
	"mypkg/b"
)

//导入包的时候会执行init
func init() {
	fmt.Println("a package init")
}

func Testinita() {
	b.Testinitb()
}
