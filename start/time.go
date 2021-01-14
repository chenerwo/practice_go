package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("current time:%v\n", now)
	/*
		year := now.Year()
		month := now.Month()
		day := now.Day()
		hour := now.Hour()
		minute := now.Minute()
		second := now.Second()

		unix_time := now.Unix()


		fmt.Printf("YMD::%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour,minute, second)
		fmt.Printf("时间戳:%d\n", unix_time)

		hour_later := now.Add(time.Hour).Unix()

		fmt.Println(hour_later)
		local := now.Local()
		fmt.Println(local)

		local1 := now.Location()
		fmt.Println(local1)
		fmt.Println("当前时间", getNowIntDate())

		//求差
		rand.Seed(now.Unix())
		str := "WW" + strconv.Itoa(getNowIntDate()) + strconv.Itoa(rand.Intn(8998) + 1001)
		str1 := "WW" + strconv.Itoa(getNowIntDate()) + strconv.Itoa(rand.Intn(8998) + 1001)
		fmt.Println(str, str1)

		a := fmt.Sprintf("%d%02d%02d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		fmt.Println(a)
	*/

}

func getNowIntDate() (r int) {
	t := time.Now()
	a := fmt.Sprintf("%d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	r, err := strconv.Atoi(a)
	if err != nil {
		return 0
	}
	return
}

func getNowFormatDate() string {
	return ""
}
