package main

import (
	"fmt"
)

//将一个列表的数据复制到另一个列表中

func printSlice(s []int)  {
	fmt.Printf("%v，长度%d,cap=%d\n", s, len(s),cap(s))
}

func main() {
	sliceA := []int{1,2,3,4,5,6,7,8}
	var sliceB = make([]int, len(sliceA))
	copy(sliceB, sliceA)
	printSlice(sliceB)
}
