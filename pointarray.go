package main

import "fmt"

const MAX int = 3

func main1()  {
	a := []int{10,100,200}
	var i int
	for i = 0; i < MAX; i++  {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
}

func main2()  {
	a := []int{10,100,200}
	var i int
	var ptr [MAX]*int

	for i=0; i<MAX; i++ {
		ptr[i] = &a[i]
	}

	for i=0; i<MAX; i++ {
		fmt.Printf("ptr[%d] = %d\n", i, ptr[i])
	}
}

func main()  {
	var a int
	var ptr *int
	var pptr **int
	a = 3000

	ptr = &a
	pptr = &ptr

	fmt.Printf("a的值%d\n", a)
	fmt.Printf("ptr的值%x\n", ptr)
	fmt.Printf("pptr的值%d\n", **pptr)
}
