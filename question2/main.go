package main

/*

The variable value is shared across all iterations of the loop. By the time the Goroutines run, the loop may have advanced and value will always reference the last value (in this case, 80). That's why all results are computed with 80, and you get 160 (because 2 * 80 = 160), which appears multiple times.


-> When you start a Goroutine inside a loop, the Goroutine does not immediately execute; it starts executing when the scheduler decides.

-> By that time, the loop has already moved on to the next iteration, so all Goroutines end up using the same reference to the loop variable value.



*/

import (
	"fmt"
	"sync"
	"time"
)

func process(v int) int {

	time.Sleep(1500 * time.Millisecond)
	return 2 * v

}

func main() {
	var wg sync.WaitGroup
	for _, value := range []int{9, 35, 27, 56, 88, 80} {

		wg.Add(1)

		go func() {
			defer wg.Done()
			fmt.Println(process(value))
		}()

	}

	wg.Wait()

}
