package main

import "fmt"

type Aa struct {
	code string
	msg  []int
}

func main() {
	var vv interface{}
	rate := 3
	if rate == 0 {
		vv = rate
	} else {
		vv = float64(rate) / 100.00
	}
	aa := fmt.Sprintf(`{"ceshi1":%v, "tax":%v}"`, vv, vv)
	fmt.Println(aa)

	/*
		aa := Aa{
			code: "1",
			msg:  []int{1, 2, 33},
		}
		abc := fmt.Sprintf("%v, %v", aa.code, aa.msg)
		fmt.Println(abc)
		fmt.Println("=========")

		var cc float64 = 2.33
		dd := fmt.Sprintf("%v", cc)
		fmt.Println(dd)

		var a float64 = 996.38
		var b float64 = 889.32

		var perInvoiceBaseMoney int32 = 100000

		var total int32 = 62105

		invoiceWxNum := math.Ceil(float64(total) / float64(perInvoiceBaseMoney))

		surplus := float32(3) / float32(100)

		fmt.Println("invoiceWxNum:", invoiceWxNum)

		fmt.Println("surplus:", surplus)

		c := a + b
		num := math.Ceil(c / 1000)
		fmt.Println(a + b)
		fmt.Println(num)

		str := "select a from b where,"
		str += "a='b'"
		//str = strings.Trim(str, ",")
		fmt.Println(str)
	*/
}
