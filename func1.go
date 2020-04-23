package main

import "fmt"

//func main()  {
//	var a int = 100
//	var b int = 200
//	var ret int
//
//	ret = max(a, b)
//	fmt.Printf("最大值是：%d\n", ret)
//}
//
//func max(num1, num2 int) int {
//	var result int
//
//	if (num1 > num2) {
//		result = num1
//	} else {
//		result = num2
//	}
//	return result
//}

//func main()  {
//	a, b := swap("Google", "Runoob")
//	fmt.Println(a, b);
//}
//
//func swap(x, y string) (string, string) {
//	return y, x
//}

func main(){
	var a int = 0
	fmt.Println("for start")
	for a:=0; a < 10; a++ {
		fmt.Println(a)
	}
	fmt.Println("for end")

	fmt.Println(a)
}

func main(){
	var a int = 0
	fmt.Println("for start")
	for a = 0; a < 10; a++ {
		fmt.Println(a)
	}
	fmt.Println("for end")

	fmt.Println(a)
}