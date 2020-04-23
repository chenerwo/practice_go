package main

import (
	"fmt"
)

func main()  {
	var chanSendOnly chan<- int

	ch := make(chan int)
	chanSendOnly = ch

	var chRecvOnly <-chan int
	chRecvOnly = ch


}
