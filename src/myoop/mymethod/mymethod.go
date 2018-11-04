package mymethod

import (
	"fmt"
	"strconv"
)

type size int

func (sz *size) dis() {
	i := int(*sz)
	*sz = 100
	fmt.Println("num is " + strconv.Itoa(i))
}
func (sz size) dis1() {
	i := int(sz)
	sz = 100
	fmt.Println("num is " + strconv.Itoa(i))
}

func Mymethod() {
	var p = person{}
	p.setidandname(1, "aaa")
	(&p).setidandname(1, "aaa")
	fmt.Println(p.getinfo(), p.name)
	fmt.Println((&p).getinfo(), p.name)

	fmt.Println("===========")

	var s = stu{}
	(*stu).setidageandname(&s, 2, 22, "xxx")
	fmt.Println(stu.getinfo(s))
	fmt.Println(s.getinfo())
	fmt.Println(s.person.getinfo())

	fmt.Println("========================1=======================")
	//========================================================================
	var i size
	i = 99
	i.dis()
	fmt.Println(i)

	i = 99
	i.dis1()
	fmt.Println(i)
}

type person struct {
	id   int
	name string
}

type stu struct {
	person
	age int
}

func (p person) getinfo() string {
	str := strconv.Itoa(p.id) + " " + p.name
	p.name = "ok"
	return str
}

func (p *person) setidandname(id int, name string) int {
	p.id = id
	p.name = name
	return 1
}

func (s stu) getinfo() string {
	str := strconv.Itoa(s.id) + " " + s.name + " " + strconv.Itoa(s.age)
	s.name = "ok"
	return str
}

func (s *stu) setidageandname(id, age int, name string) int {
	s.id = id
	s.age = age
	s.name = name
	return 1
}
