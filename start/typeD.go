package main

import (
	"fmt"
	"reflect"
)

//在结构体成员嵌入时使用别名

type Brand struct {

}

// 为商标结构添加Show()方法
func (t Brand) show() {

}

// 为Brand定义一个别名FakeBrand
type FakeBrand = Brand

// 定义车辆结构
type Vehicle struct {
	// 嵌入两个结构
	FakeBrand
	Brand
}

func main()  {
	// 声明变量a为车辆类型
	var a Vehicle

	// 指定调用FakeBrand的Show
	a.FakeBrand.show()

	var b string = "5"

	bi := int(b)

	// 指定调用FakeBrand的Show
	ta := reflect.TypeOf(a)

	for i := 0; i < ta.NumField(); i++ {
		//a的成员信息
		f := ta.Field(i)

		fmt.Printf("FieldName: %v, FieldType: %v\n", f.Name, f.Type.Name())
	}
}

