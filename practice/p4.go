package main

import (
	"fmt"
)

//题目：输入某年某月某日，判断这一天是这一年的第几天？

var (
	input_Data int
	result int
	year_data int
	month_data int
	day_data int
	sum int
)

func get_Input(input_Type string) (int, error) {
	switch input_Type {

	case "year":
		fmt.Println("输入年份:")
		fmt.Scanln(&input_Data)
		if input_Data < 1 {
			fmt.Println("年份输入错误")
			break
		}
		result = input_Data
	case "month":
		months := []int{0,31,59,90,120,151,181,212,243,273,304,334}
		fmt.Println("输入月份:")
		fmt.Scanln(&input_Data)
		if (input_Data > 12) || (input_Data < 0) {
			fmt.Println("月份输入错误")
			break
		}
		result = months[input_Data - 1]
	case "day":
		fmt.Println("输入日期:")
		fmt.Scanln(&input_Data)
		if (input_Data < 0) && (input_Data > 31) {
			fmt.Println("日期输入错误")
			break
		}
		result = input_Data
	default:
		return 0, fmt.Errorf("输入参数非法:%s", input_Type)
	}
	return result, nil
}

func main() {
	year_data, _ = get_Input("year")
	month_data, _ = get_Input("month")
	day_data, _ = get_Input("day")

	sum += day_data + month_data
	leap := 0
	if (year_data % 400 == 0) || ((year_data %4 == 0) && (year_data % 100 != 0)) {
		leap = 1
	}
	if (leap == 1) && (month_data > 2) {
		sum += 1
	}
	fmt.Printf("该日为第%d天\n", sum)
}
