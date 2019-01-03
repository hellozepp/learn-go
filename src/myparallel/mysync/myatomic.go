package mysync

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func MyAtomic() {

	// 定义一个整数
	var ops int64 = 0

	// 使用50个线程给ops累加数值
	for i := 0; i < 50; i++ {
		go func() {
			for {
				if atomic.LoadInt64(&ops) >= 10000 {
					fmt.Println("finished...")
					break
				}
				// 每次加1
				atomic.AddInt64(&ops, 1)

				// 这个函数用于时间片切换
				//可以理解为高级版的time.Sleep()
				//避免前面的for循环将CPU时间片都卡在一个线程里，使得其它线程没有执行机会
				runtime.Gosched()
			}
		}()
	}

	//停一秒，上面50个线程有1秒的执行时间
	time.Sleep(time.Second)

	// 获取结果
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", int(opsFinal))
	fmt.Println("sub:", atomic.AddInt64(&ops, -3)) //有符号减法

	//无符号
	var ui uint32 = 0
	fmt.Println("ui:", ui)
	var NN int32 = -3
	fmt.Println("int32 NN:", NN)
	atomic.AddUint32(&ui, ^uint32(-NN-1))
	fmt.Println("ui:", ui)
	fmt.Println("ui:", int32(ui))

}
