package main

import (
	"fmt"
)

func main()  {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//s1 := 0
	//i1 := 0
	//for ; i1 <= 12; {
	//	i1++
	//	s1 += i1
	//}
	//fmt.Println(s1)
	//
	////无限循环
	//za := 0
	//for {
	//	za++
	//}

	/**
	strings := []string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := []int8{1, 2, 3, 4, 5, 6}
	for k, v := range numbers {
		fmt.Println(k, v)
	}
	 */
	var a int = 10
	LOOP: for a < 20 {
		if a == 15 {
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d \n", a)
		a++
	}
}

