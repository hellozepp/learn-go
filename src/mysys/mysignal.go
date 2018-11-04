package mysys

import "fmt"
import "time"
import "os"
import "os/signal"
import "syscall"

func sysSignalHandleDemo() {
	ss := signalSetNew()
	handler := func(s os.Signal, arg interface{}) {
		fmt.Printf("handle signal: %v\n", s)
	}

	ss.register(syscall.SIGINT, handler)  //kill -2
	ss.register(syscall.SIGUSR1, handler) //kill -10
	ss.register(syscall.SIGUSR2, handler) //kill -12

	for {
		c := make(chan os.Signal)
		var sigs []os.Signal
		for sig := range ss.m {
			sigs = append(sigs, sig)
		}
		signal.Notify(c)
		sig := <-c

		err := ss.handle(sig, nil)
		if err != nil {
			fmt.Printf("unknown signal received: %v\n", sig)
			os.Exit(1)
		}
	}
}

func MySignal() {
	go sysSignalHandleDemo()
	time.Sleep(time.Hour) // make the main goroutine wait!
}

type signalHandler func(s os.Signal, arg interface{})

type signalSet struct {
	m map[os.Signal]signalHandler
}

func signalSetNew() *signalSet {
	ss := new(signalSet)
	ss.m = make(map[os.Signal]signalHandler)
	return ss
}

func (set *signalSet) register(s os.Signal, handler signalHandler) {
	if _, found := set.m[s]; !found {
		set.m[s] = handler
	}
}

func (set *signalSet) handle(sig os.Signal, arg interface{}) (err error) {
	if _, found := set.m[sig]; found {
		set.m[sig](sig, arg)
		return nil
	} else {
		return fmt.Errorf("No handler available for signal %v", sig)
	}

	panic("won't reach here")
}
