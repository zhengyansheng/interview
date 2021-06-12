package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex
	var chInt = make(chan int)
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(chInt chan int, lock sync.Mutex) {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			fmt.Println(<-chInt)
		}(chInt, lock)
		chInt <- i
	}
	wg.Wait()
}
