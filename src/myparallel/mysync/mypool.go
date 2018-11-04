package mysync

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

func MyPooltest() {
	//禁用GC,并保证在main函数结束恢复GC

	defer debug.SetGCPercent(debug.SetGCPercent(-1))
	var i int32 = 0
	newFunc := func() interface{} {
		return atomic.AddInt32(&i, 1)
	}
	pool := sync.Pool{New: newFunc}

	v1 := pool.Get()
	fmt.Println(v1)

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
