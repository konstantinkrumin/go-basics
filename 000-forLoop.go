package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("Go!")
			wg.Done()
		}()
	}

	wg.Wait()
}
