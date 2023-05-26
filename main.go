package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("hello world")
	
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("loop:", i)
		}
	}()
	wg.Wait()
}
