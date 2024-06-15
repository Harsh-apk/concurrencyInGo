package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	x := []int{1, 2, 3, 4, 5}

	for _, index := range x {
		wg.Add(1)
		go myFunction(index)

	}
	//waits until all the goroutines return
	wg.Wait()

}

func myFunction(num int) {
	defer wg.Done()
	fmt.Println(num)
}
