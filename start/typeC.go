package main

import (
	"fmt"
)

// 将NewInt定义为int类型
type newInt int

type intAlias = int

func main()  {
	// 将a声明为NewInt类型
	var a newInt

	// 查看a的类型名
	fmt.Printf("a type : %T\n", a)

	// 将a2声明为IntAlias类型
	var a2 intAlias
	fmt.Printf("a2 type : %T\n", a2)
}
