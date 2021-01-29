package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	//生成10个0-99之间的随机数
	for i := 0; i < 3; i++ {
		fmt.Println(rand.Intn(8998) + 1001)
	}
	//for i := 0; i < 3; i++ {
	//	a := strconv.Itoa(rand.Intn(8998) + 1001)
	//	fmt.Println(a)
	//}
}
