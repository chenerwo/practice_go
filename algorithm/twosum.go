package main

import "fmt"

var nums []int = []int{2, 7, 11, 15}
var target int = 9

func twoSum(nums []int, target int) []int {
	n := make(map[int]int, len(nums))
	for k, v := range nums {
		s := target - v
		if m, ok := n[s]; ok {
			return []int{m, k}
		}
		n[v] = k
	}
	return []int{}
}

func main() {
	res := twoSum(nums, target)
	fmt.Println(res)
}
