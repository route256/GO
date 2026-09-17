package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	// defer cancel()

	sharedResource := map[int]int{}
	var mx sync.RWMutex

	go func() {
		mx.Lock()
		defer mx.Unlock()

		i := 0
		for {
			i++
			sharedResource[i] = i
		}
	}()
	go func() {
		i := 0
		for {
			i++
			mx.Lock()
			sharedResource[i] = i
			mx.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			time.Sleep(time.Millisecond * 2)
			i := i

			mx.RLock()
			fmt.Println(sharedResource[i])
			mx.RUnlock()
		}
	}()

	mx.RLock()
	fmt.Println(sharedResource)
	mx.RUnlock()
	time.Sleep(time.Minute)
}
