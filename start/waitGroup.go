package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		fmt.Println("group：", i)
		wg.Done()
	}

	wg.Wait()
}
