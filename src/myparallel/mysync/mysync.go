package mysync

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroupTest() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Second)
			fmt.Println("ok", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("WaitGroup is ok !")
}
