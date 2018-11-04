package myintf

import (
	"fmt"
	"strconv"
)

func Myintf() {
	s := new(student)
	s.id = 1
	s.name = "aaa"
	s.age = 22
	str := s.dorun(100)
	printinfo(str, s)

	fmt.Println("========================1=======================")
	//========================================================================
	s2 := student{}
	s2.id = 2
	s2.name = "xxx"
	s2.age = 33
	dis(&s2)

}

type irun interface {
	dorun(i int) string
}

type ijump interface {
	dojump(i int) string
}

type person struct {
	id   int
	name string
}

type student struct {
	person
	age int
}

func (s *student) dorun(i int) string {
	return strconv.Itoa(s.id) + " " + s.name + " is runing at " + strconv.Itoa(i) + " speed...... his age is " + strconv.Itoa(s.age)
}

func dis(obj irun) {
	fmt.Println(obj.dorun(200))
	//(*student)(obj).dojump(200) 无法这样强转，因为类型不同
	fmt.Println(obj.(*student).dojump(100)) //可以通过　.()　运算符拿到原始内存的结构指针
	fmt.Println(obj, obj.(*student), *obj.(*student))
	fmt.Println("==================")
	fmt.Println((*obj.(*student)).dojump(200))
	fmt.Println((interface{})(obj).(ijump).dojump(800))
	fmt.Println((interface{})(obj).(irun).dorun(900))

	fmt.Println((interface{})(obj).(*student).person.dojump(1000)) //通过转成子类然后获取父类
	dis2(&(obj.(*student).person))
}

func (p *person) dojump(i int) string {
	return strconv.Itoa(p.id) + " " + p.name + " is jumping at " + strconv.Itoa(i) + " speed......"
}
func dis2(obj ijump) {
	fmt.Println(obj.dojump(200))
}

func printinfo(s string, obj interface{}) {
	if v, ok := obj.(student); ok {
		fmt.Println("hello " + v.name)
		fmt.Println(s)
	} else {
		fmt.Println("呵呵")
	}

	switch value := obj.(type) {
	case person:
		fmt.Println(value.name + " is person")
	case student:
		fmt.Println(value.name + " is student")
	default:
		fmt.Println("呵呵")
	}
}
