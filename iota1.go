package main

import "fmt"

type Weekday int

func main()  {

	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		ii
	)
	fmt.Println(a,b,c,d,e,f,g,h,ii)

	const (
		i = 1 << iota
		j = 3 << iota
		k
		l
	)
	fmt.Println(i,j,k,l)
	//res 1,6,12,24

	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
}