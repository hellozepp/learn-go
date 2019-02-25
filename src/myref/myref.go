package myref

import (
	"fmt"
	"reflect"
	"strconv"
)

type person struct {
	int
	name string
}

func (p person) Dis() { //*student不属于student 仅属于*student对象
	fmt.Println(p)
}

type student struct {
	person
	age int
	Sex int //可以反射赋值
}

func (s student) Sayhello(toname string) (string, int) { //如果方法不是大写第一个字母的公开方法，是无法被反射获取的
	return s.name + " say hello to " + toname, 1
}

func (s *student) Dis() { //*student不属于student 仅属于*student对象
	fmt.Println(s)
}

type MyStruct struct {
	name string
}

func (this *MyStruct) GetName() string {
	return this.name
}
func Myref() {
	s := student{person: person{int: 1, name: "aaa"}, age: 22}
	t := reflect.TypeOf(s)
	fmt.Println(t) //获取字段属性以及字段相关的东西
	t11 := reflect.TypeOf(&s)
	fmt.Println(t11)
	k := reflect.TypeOf(s).Kind()
	k1 := reflect.TypeOf(&s).Kind()
	v := reflect.ValueOf(s)

	fmt.Println(k) //类型
	fmt.Println(k1)
	fmt.Println(v) //获取字段的值以及值相关的东西
	fmt.Println(reflect.TypeOf(v))
	fmt.Println("========================1=======================")
	//========================================================================
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i)
		val := v.Field(i)
		fmt.Println(key.Name, key.Type, "|", val) //val 可读可写 因为是公有,key只读
	}

	fmt.Println("======")

	fmt.Println(t.Kind())

	if k := t.Kind(); k == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			key := t.Field(i)
			val := v.Field(i)
			fmt.Println(key.Name, key.Type, "|", val)
		}
	}

	fmt.Println("========================2=======================")
	//========================================================================
	t1 := reflect.TypeOf(&s)
	v1 := reflect.ValueOf(&s)
	fmt.Println(t1)
	fmt.Println(v1)

	fmt.Println("======")

	if k := t1.Kind(); k == reflect.Ptr {
		tt := t1.Elem() //指针通过elem获取内容
		vv := v1.Elem()
		for i := 0; i < tt.NumField(); i++ {
			key := tt.Field(i)
			val := vv.Field(i)
			fmt.Println(key.Name, key.Type, "|", val)
		}
	}
	fmt.Println("======")
	for i := 0; i < t1.NumMethod(); i++ { //可以获取所有方法，但是不能用指针获取属性
		m := t1.Method(i)
		fmt.Println(m.Name, m.Type)
	}

	fmt.Println("========================3=======================")
	//========================================================================
	fmt.Println(t.FieldByName("person"))
	fmt.Println(t.FieldByIndex([]int{0}))
	fmt.Println(t.FieldByIndex([]int{0, 0}), t.FieldByIndex([]int{0, 1}))
	m2, _ := t.MethodByName("Sayhello")
	fmt.Println(m2)
	fmt.Println("===========mystruct==========")
	a := new(MyStruct)
	a.name = "hehe"
	fmt.Println(reflect.ValueOf(a))
	//b1 := reflect.ValueOf(a).MethodByName("GetName").Call([]reflect.Value{})
	b1 := reflect.ValueOf(a).MethodByName("GetName").Call(nil)
	ref := reflect.ValueOf(a).Elem()
	typp := reflect.TypeOf(a).Elem()
	fmt.Println(b1[0]) //反射调用
	fmt.Println(ref.Kind())
	for i := 0; i < ref.NumField(); i++ {
		key := ref.Field(i)
		fmt.Println(typp.Field(i).Name, key)
	}

	val := reflect.ValueOf(a).Elem().FieldByName("name")
	fmt.Println("name---->", val)
	fmt.Println(reflect.ValueOf(a).Elem().FieldByName("name").CanSet())
	fmt.Println(reflect.ValueOf(&(a.name)).Elem().CanSet())

	var c string = "yejianfeng"
	p := reflect.ValueOf(&c)
	fmt.Println(p.CanSet())        //false
	fmt.Println(p.Elem().CanSet()) //true
	p.Elem().SetString("newName")
	fmt.Println(c)
	fmt.Println("========================4=======================")
	//========================================================================
	x := 123
	vx := reflect.ValueOf(&x)
	vx.Elem().SetInt(333)
	fmt.Println(x)

	fmt.Println("===1===")

	s11 := student{person: person{int: 1, name: "aaa"}, age: 22, Sex: 1}
	fmt.Println(s11)
	v11 := reflect.ValueOf(&s11)
	fmt.Println(v11.Elem().FieldByName("age").CanSet()) //小写是只读
	fmt.Println(v11.Elem().FieldByName("Sex").CanSet()) //大写才是读写

	if v11.Elem().FieldByName("age").CanSet() {
		v11.Elem().FieldByName("age").SetInt(99)
	} else {
		fmt.Println("age cant set")
	}
	if v11.Elem().FieldByName("Sex").CanSet() {
		v11.Elem().FieldByName("Sex").SetInt(2)
	}
	fmt.Println(v11)
	fmt.Println(s11)

	fmt.Println("===2===")

	fv := reflect.ValueOf(prints) //获取一个普通函数的句柄

	params := make([]reflect.Value, 1) //参数
	params[0] = reflect.ValueOf(20)    //参数设置为20
	rs := fv.Call(params)              //rs作为结果接受函数的返回值
	fmt.Println(rs[0])
	fmt.Println("result: " + rs[0].Interface().(string)) //当然也可以直接是rs[0].Interface()

	fmt.Println("===3===")

	params1 := make([]reflect.Value, 1) //参数
	params1[0] = reflect.ValueOf("ppp") //参数设置为ppp
	b := reflect.ValueOf(s11).MethodByName("Sayhello").Call(params1)
	fmt.Println(b[0], "|", b[1])

	b2 := reflect.ValueOf(&s11).MethodByName("Dis").Call([]reflect.Value{})
	fmt.Println(b2)

	fmt.Println("========================5=======================")
	// object TypeOf
	a2 := new(person)
	a2.name = "xxx"
	typ := reflect.TypeOf(a2)
	fmt.Println(typ)        // output: "*myref.person"
	fmt.Println(typ.Elem()) // output: "myref.person"

	// reflect.Type Base struct
	fmt.Println(typ.NumMethod())                   // 1
	fmt.Println(typ.Method(0))                     // {GetName  func(*myref.person) <func(*myref.person) Value> 0}
	fmt.Println(typ.Name())                        // ""
	fmt.Println(typ.PkgPath())                     // ""
	fmt.Println(typ.Size())                        // 8
	fmt.Println(typ.String())                      // *myref.person
	fmt.Println(typ.Elem().String())               // myref.person
	fmt.Println(typ.Elem().FieldByIndex([]int{0})) // {int go.builtin int  0 [0] true}
	fmt.Println(typ.Elem().FieldByName("name"))    // {name iotestgo/myref string  8 [1] false} true

	fmt.Println(typ.Kind() == reflect.Ptr)           // true
	fmt.Println(typ.Elem().Kind() == reflect.Struct) // true

	fmt.Println("======")

	fmt.Println(reflect.TypeOf(12.12).Bits()) // 64, 因为是float64

	fmt.Println("======")

	cha := make(chan int)
	fmt.Println(reflect.TypeOf(cha).ChanDir()) // chan

	var fun func(x int, y ...float64) string
	var fun2 func(x int, y float64) string
	fmt.Println(reflect.TypeOf(fun).IsVariadic())  // true
	fmt.Println(reflect.TypeOf(fun2).IsVariadic()) // false
	fmt.Println(reflect.TypeOf(fun).In(0))         // int
	fmt.Println(reflect.TypeOf(fun).In(1))         // []float64
	fmt.Println(reflect.TypeOf(fun).NumIn())       // 2
	fmt.Println(reflect.TypeOf(fun).NumOut())      // 1
	fmt.Println(reflect.TypeOf(fun).Out(0))        // string

	fmt.Println("======")

	mp := make(map[string]int)
	mp["test1"] = 1
	fmt.Println(reflect.TypeOf(mp).Key(), reflect.TypeOf(mp).Elem())                 //string int
	fmt.Println(reflect.ValueOf(mp).Type(), reflect.ValueOf(mp).MapKeys()[0].Type()) //map[string]int string

	arr := [2]string{"test", "xxx"}
	fmt.Println(reflect.TypeOf(arr).Len()) // 2

	fmt.Println("========================6=======================")
	//========================================================================

	type Obj1 struct {
		Name string
	}

	var objmap map[string]reflect.Type = make(map[string]reflect.Type)
	objmap["Obj1"] = reflect.TypeOf(Obj1{})
	objmap["int"] = reflect.TypeOf(1)

	i1 := reflect.New(objmap["Obj1"]) //相当于java的Class.forName()
	i1.Elem().FieldByName("Name").SetString("xxx")
	fmt.Println(i1, i1.Elem().Addr(), i1.Elem())

	i2 := reflect.New(objmap["int"]) //相当于java的Class.forName()
	i2.Elem().SetInt(222)
	fmt.Println(i2, i2.Elem().Addr(), i2.Elem())

}
func prints(i int) string {
	fmt.Println("i =", i)
	return strconv.Itoa(i)
}
