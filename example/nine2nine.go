package main

import "fmt"

func main() {
	var a string
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			//fmt.Println(i, "*", j, "=", i*j)
			a += fmt.Sprintf("%d * %d = %d ", i, j, i*j)
			if i == j {
				a += "\n"
			}
		}
	}
	fmt.Println(a)
}
