package main

import (
	"fmt"
	"sort"
)

var (
	data []int
	input_data int
)

//输入5个整数，请把这5个数分别由小到大和由小到大输出。

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Printf("输入第%d个参数:", int(i))
		fmt.Scanln(&input_data)
		data = append(data, input_data)
	}
	//fmt.Println("结果:", data)
	sort.Ints(data)
	fmt.Println("正序:", data)

	fmt.Println("反序:", sort.Reverse(data))
}
