package main

import (
	"fmt"
	"time"
)


func main()  {
	now := time.Now()
	fmt.Printf("current time:%v\n", now)
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

	//求差

}