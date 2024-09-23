package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// returns true if number is prime
func isPrime(v int64) bool {
	sq := int64(math.Sqrt(float64(v))) + 1
	var i int64
	for i = 2; i < sq; i++ {
		if v%i == 0 {
			return false
		}
	}
	return true
}

// get a random prime number between 1 and maxP
func getPrime(maxP int64) int64 {
	var i int64
	for i = 0; i < maxP; i++ {
		n := rand.Int63n(maxP)
		if isPrime(n) {
			return n
		}
	}
	return 1 // just in case
}

func main() {
	var primes []int64              // slice of prime numbers
	const maxPrime int64 = 10000000 // max value for primes
	start := time.Now()
	for i := 0; i < 10000; i++ {
		p := getPrime(maxPrime) // add a new prime
		primes = append(primes, p)
	}
	fmt.Println(primes)

	end := time.Now()
	fmt.Println("End of program.", end.Sub(start))
}

/*
package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

// returns true if number is prime
func isPrime(v int64) bool {
	sq := int64(math.Sqrt(float64(v))) + 1
	var i int64
	for i = 2; i < sq; i++ {
		if v%i == 0 {
			return false
		}
	}
	return true
}

// get a random prime number between 1 and maxP
func getPrime(maxP int64) int64 {
	for {
		n := rand.Int63n(maxP)
		if isPrime(n) {
			return n
		}
	}
}

// primeGenerator generates prime numbers and sends them to the channel
func primeGenerator(maxP int64, primeChan chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		primeChan <- getPrime(maxP)
	}
}

func main() {
	const maxPrime int64 = 10000000 // max value for primes
	const numPrimes = 10000         // number of primes to generate
	const numWorkers = 5            // number of concurrent workers

	var primes []int64              // slice of prime numbers
	primeChan := make(chan int64)   // channel for prime numbers
	var wg sync.WaitGroup           // wait group to ensure all workers finish

	start := time.Now()

	// Start 5 goroutines to generate primes concurrently
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go primeGenerator(maxPrime, primeChan, &wg)
	}

	// Collect 10,000 prime numbers from the channel
	for i := 0; i < numPrimes; i++ {
		primes = append(primes, <-primeChan)
	}

	// Stop workers after receiving the desired number of primes
	close(primeChan)
	wg.Wait() // Wait for all goroutines to finish

	end := time.Now()
	fmt.Println("End of program.", end.Sub(start))
}

*/
