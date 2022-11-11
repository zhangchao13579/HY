package main

import (
	"fmt"
	"time"
)

// 打印奇数的函数
func f1(ch chan int) {
	for i := 1; i <= 100; i++ {
		ch <- i //将数据发送到ch
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

// 打印偶数的函数
func f2(ch chan int) {
	for i := 1; i <= 100; i++ {
		<-ch //从ch中取值
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	ch := make(chan int) //确保同步，采用无缓冲的channel
	go f1(ch)
	go f2(ch)
	time.Sleep(3 * time.Second) //确定时间间隔
}
