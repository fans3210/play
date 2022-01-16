package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("hello world")

	wg := sync.WaitGroup{}
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("i = ", i)
			ch <- i
		}(i)

	}

	go func() {
		for c := range ch {
			fmt.Println("channel value: ", c)
		}
	}()

	wg.Wait()
	close(ch)

	fmt.Println("done")
}
