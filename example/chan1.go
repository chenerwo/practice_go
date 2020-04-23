package main

import (
	"fmt"
	"math/rand"
	"time"
)

//数据生产者
func productor(header string, channel chan<- string)  {
	//无限循环，不停地生产数据
	for {
		// 将随机数和字符串格式化为字符串发送给通道
		channel <- fmt.Sprintf("%s:%v", header,rand.Int31())
		//等待1秒
		time.Sleep(time.Second)
	}
}

func consumer(channel <-chan string)  {
	//不停的获取数据
	for {
		message := <-channel
		fmt.Println(message)
	}
}

func main()  {
	channel := make(chan string)

	go productor("cat", channel)
	go productor("dog", channel)
	consumer(channel)
}

