package main

import "fmt"

func lengthOfLongestSubstring(s string) int {

	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)

	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		fmt.Println("round ", i, "======")
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		//fmt.Println("m,s[rk+1],m[s[rk+1]:\n", m, s[rk+1], m[s[rk+1]])
		fmt.Println("s[rk+1]", s[rk+1])
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			fmt.Println("右移指针:", m[s[rk+1]])
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = maxim(ans, rk-i+1)
	}
	return ans
}

func maxim(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	//str := "englishisagretlopnguage"
	str := "pwwkew"
	l := lengthOfLongestSubstring(str)
	fmt.Println(l)
}
