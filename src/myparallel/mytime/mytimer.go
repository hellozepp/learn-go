package mytime

import (
	"fmt"
	"time"
)

func MyTimer() {
	timer1 := time.NewTimer(time.Second * 1)
	<-timer1.C
	fmt.Println("Timer 1 expired")
	timer2 := time.NewTimer(time.Second)
	if timer1.Stop() {
		fmt.Println("Timer 2 stopped")
	}
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
