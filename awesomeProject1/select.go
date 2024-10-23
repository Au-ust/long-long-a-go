package main

import "fmt"

func main() {
	/*
		select语句和switch很像，都有case和default,区别是select会随机执行一个可执行的case，直到没有case可以执行的时候执行defalut，什么都没有的时候，阻塞
	*/
	//直接这么写会死锁，因为只有读取没有写入

	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- 10
	}() //加入写入ch1
	go func() {
		ch2 <- 1999
	}()
	select {
	case num1 := <-ch1:
		fmt.Println("ch1获取的数据为", num1)
	case num2, ok := <-ch2:
		if ok {
			fmt.Println("ch2获取的数据为", num2)
		} else {
			fmt.Println("ch2 closed")
		} //无论几次执行都只执行第一个case,因为第二个case里没写数据，只读会阻塞，要想读case2需要主动往里写
	default:
		fmt.Println("this is default")
	}
	fmt.Println("main end")
}
