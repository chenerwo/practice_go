package main

import "fmt"

func main()  {
	/**
	var a = "Runoob"
	fmt.Println(a);

	var b, c int = 1, 2
	fmt.Println(b,c);

	var d int32
	fmt.Println(d);

	var f bool
	fmt.Println(f)
	/**
	var i int
	var f float32
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)



	var ddf = true
	fmt.Println(ddf)

	v_name := "value"
	fmt.Println(v_name)

	var vname1,vname2,vname3 int8
	fmt.Println(vname1,vname2,vname3);
	 */
	var v_name1,v_name2,v_name3 = 15,"hee",1.5
	fmt.Println(v_name1, v_name2, v_name3);

	//var (
	//	vv1 int
	//	vv2 string
	//	vv3 float64
	//)

	const LENGTH int  = 10
	const WIDTH int = 5
	var area int
	const a,b,c  = 1, false,"str"

	area = LENGTH * WIDTH
	fmt.Printf("面积为：%d", area);
	println()
	println(a,b,c)

	const (
		UNKNOW = 0
		FEMALE = 1
		MALE = 2
	)
}
/**
var age int;

fruit = apples + oranges;

var b bool = true;
*/
/**
派生类型
a 指针 Pointer
b 数组
c 结构化 struct
d Channel
e 函数
f 切片切片
g 接口 interface
h Map
 */
