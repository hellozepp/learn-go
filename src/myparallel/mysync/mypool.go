package mysync

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"
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

type Dao struct {
	bp sync.Pool //不能自由控制Pool中元素的数量，放进Pool中的对象每次GC发生时都会被清理掉
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		bp: sync.Pool{
			New: func() interface{} {
				return &bytes.Buffer{}
			},
		},
	}
	return
}

var defaultSpliter string = ","
var defaultReplacer string = " "

var ErrType error

func (d *Dao) Infoc(args ...string) (value string, err error) {
	if len(args) == 0 {
		return
	}
	// fetch a buf from bufpool
	buf, ok := d.bp.Get().(*bytes.Buffer)
	if !ok {
		return "", ErrType
	}
	// append first arg
	if _, err := buf.WriteString(args[0]); err != nil {
		return "", err
	}
	for _, arg := range args[1:] {
		// append ,arg
		if _, err := buf.WriteString(defaultSpliter); err != nil {
			return "", err
		}
		if _, err := buf.WriteString(strings.Replace(arg, defaultSpliter, defaultReplacer, -1)); err != nil {
			return "", err
		}
	}
	value = buf.String()
	buf.Reset()
	d.bp.Put(buf)
	return
}
