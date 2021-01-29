package main

import (
	"fmt"
	"runtime"
)

func main() {
	//cn := runtime.GOMAXPROCS(1)
	cn := runtime.NumGoroutine()
	fmt.Println(cn)
}
