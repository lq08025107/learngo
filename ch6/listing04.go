package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(2)
	wg.Add(2)

	fmt.Println("Create GoRoutines")
	go PrintPrime("A")
	go PrintPrime("B")

	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("Terminating Program")
}

func PrintPrime(prefix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
