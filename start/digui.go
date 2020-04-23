package main

import (
	"fmt"
)

func main()  {
	var count int = 6
	jie := jiecheng(count)
	fmt.Println(jie)
}

func jiecheng(count int) (res int) {
	if count > 0 {
		res = count * jiecheng(count-1)
		return
	}
	res = 1
	return
}