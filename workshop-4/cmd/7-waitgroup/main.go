package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			time.Sleep(time.Second)
			fmt.Println("Done", i)
			wg.Done()
		}()
	}

	wg.Wait()
}
