package main

import (
	"fmt"
	"sync"
	"time"
)

func process(n int) int {
	time.Sleep(2 * time.Second)
	return 2 * n
}

func main() {

	var wg sync.WaitGroup

	for _, value := range []int{9, 35, 27, 56, 88, 80} {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(process(n))
		}(value)
	}
	wg.Wait()

}
