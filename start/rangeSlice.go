package main

import (
	"fmt"
)

func main()  {
	var aS = []int{10, 20, 30, 40, 50}
	for index, value := range aS {
		fmt.Println(index, value)
	}
}