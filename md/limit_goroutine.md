# limit goroutine

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int)  {
	time.Sleep(time.Millisecond*100)
	fmt.Printf("worker num: %d\n", i)
}

func main() {
	maxGogoutineNum := 10
	wg := sync.WaitGroup{}
	ch := make(chan struct{}, maxGogoutineNum)

	for i := 0; i < 100; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer func() {
				time.Sleep(time.Second*1)
				<- ch
			}()
			worker(i)
		}(i)
	}
	wg.Wait()
}
```