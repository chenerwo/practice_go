package main

import (
	"fmt"
)

func main()  {
	//var mulS [][]int
	//mulS = [][]int{{10}, {10, 20}}

	//or
	mul1 := [][]int{{100}, {110, 120}}

	mul1[0] = append(mul1[0], 200)
	fmt.Println(mul1)


}
