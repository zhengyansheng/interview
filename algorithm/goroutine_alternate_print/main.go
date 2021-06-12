package main

import (
	"fmt"
	"sync"
)

/*
	使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
	12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
*/

// 生成数组 元素是字母
func genW() (arr []string) {
	for i := 'a'; i <= 'z'; i++ {
		arr = append(arr, fmt.Sprintf("%c", i))
	}
	return
}

func main() {
	var wg sync.WaitGroup

	wArr := genW()
	fmt.Println(len(wArr))
	stopNChan := make(chan struct{}, 1)
	stopWChan := make(chan struct{}, 1)

	wg.Add(2)

	go func() {
		defer wg.Done()
		i := 0
		for i < 28 {
			<-stopWChan
			i++
			fmt.Printf("%d", i)
			i++
			fmt.Printf("%d", i)
			stopNChan <- struct{}{}
		}
		close(stopNChan)
	}()

	go func() {
		defer wg.Done()
		i := 0
		for i < 25 {
			stopWChan <- struct{}{}
			<-stopNChan
			i++
			fmt.Printf("%s", wArr[i-1])
			i++
			fmt.Printf("%s", wArr[i-1])
		}
		close(stopWChan)
	}()

	wg.Wait()
}
