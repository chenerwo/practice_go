package main

import "fmt"

var a int8 = 3
var b int8

func main()  {

	switch a {
	case 1, 2, 3:
		b = 2
		fallthrough
	case 4:
		b = 4
		fallthrough
	case 5:
		b = 5
		fallthrough
	case 9:
		fmt.Println("aabbccd")
		fallthrough
	default:
		b = 0
	}
	fmt.Println(b)
}
