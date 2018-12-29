package mytime

import (
	"fmt"
	"sync"
	"time"
)

func MyTicker() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println(err)
			}
		}()
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
		fmt.Println("Ticker stopped") //不执行
		//wg.Done()
	}()
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	//wg.Wait()
	time.Sleep(time.Millisecond * 1600)

}
