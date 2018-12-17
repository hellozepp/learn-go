package mypkg

//import "fmt"  //1

import stdio "fmt" //2 起别名

import . "fmt" //3 导入到本包，最大问题是容易产生二异性，链接时连不过去
import (
	//stdio "fmt"
	"mypkg/a" //此处执行init
	"strconv"
) //4

func Mypkg() {
	Println("hello world!")             //1
	stdio.Println("stdio hello world!") //2
	//Println(". hello world!")           //3
	var istr int = 123456 //4
	Println(strconv.Itoa(istr))

	Println("========================1=======================")
	//========================================================================

	a.Testinita() //并非使用的时候才init
}
