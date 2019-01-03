package mysync

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

func MyPooltest() {
	// 建立对象
	var pool1 = &sync.Pool{New: func() interface{} { return "Hello,xiequan" }}
	// 准备放入的字符串
	val := "Hello,World!"
	pool1.Put(val)
	log.Println(pool1.Get())
	// 再取就没有了,会自动调用NEW
	log.Println(pool1.Get())

	//禁用GC,并保证在main函数结束恢复GC

	defer debug.SetGCPercent(debug.SetGCPercent(-1))
	var i int32 = 0
	newFunc := func() interface{} {
		return atomic.AddInt32(&i, 1)
	}
	pool := sync.Pool{New: newFunc}
	pool.Get()
	pool.Get()
	pool.Get()
	v1 := pool.Get()

	fmt.Println(reflect.TypeOf(v1).Kind(), v1)
	pool.Put(10)
	pool.Put(20)
	pool.Put(30)
	v2 := pool.Get()
	fmt.Println(v2)

	//垃圾回收对临时对象的影响
	debug.SetGCPercent(100)
	runtime.GC()

	v3 := pool.Get()

	fmt.Println(v3)

	pool.New = nil

	v4 := pool.Get()

	fmt.Println(v4)
}
