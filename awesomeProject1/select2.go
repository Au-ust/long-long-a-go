package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()
	select {
	case <-ch1:
		fmt.Println("ch1已执行")
	case <-ch2:
		fmt.Println("ch2已执行")
	case t := <-time.After(2 * time.Second):
		fmt.Println("超时了，当前时间是:", t) //超时了，当前时间是: 2024-10-22 15:58:58.7965958 +0800 CST m=+2.002385501
	default:
		fmt.Println("this is a default") //执行default，因为另外的case阻塞或者暂时阻塞
	}
}
