package main

import (
	"fmt"
)

func main() {
	/*
		双向通道：我们之前学的就是双向通道，又能获取又能发送
		单向通道：只能发送或只能获取
	*/
	ch1 := make(chan string)
	done1 := make(chan bool)
	go sendData5(ch1, done1)
	data := <-ch1
	fmt.Println("子goroutine传来：", data)
	ch1 <- "我是main"
	<-done1 //通知main goroutine
	fmt.Println("main over")
	//写个单向
	ch2 := make(chan int) //双向
	//ch3 := make(chan<- int) //单向，只能写，不能读
	//ch4 := make(<-chan int) //单向，只能读，不能写
	//ch4 <- 100,报错，只能读，不能写,cannot send to receive-only channel ch4 (variable of type <-chan int)
	//data := <-ch3,报错，ch3只能写
	go fun1(ch2)
	//fun1(ch3)，也报错，因为读了没写
	data2 := <-ch2
	//data3 := <-ch3 //报错，因为只写
	fmt.Println(data2)
	//fmt.Println(data3)
	/*妈呀，单向通道的声明简直就像我的人生一样没有意义
	也就是说，双向通道也可以实现单向通道的能力，但是单向通道这个东西出现的意义，仅仅只是为了限制某些函数的参数类型
	例如你写了一个函数，这个函数专门来负责一个通道的写，那么只要你把函数的参数写成单向通道的形式，就可以不让读的通道参与这个函数
	如果你又写了另一个负责读的函数，那么负责写的管道就进不去
	也就是说单向通道只是为了区分一个双向管道的不同方面的使用的！
	无语
	*/
}
func sendData5(c chan string, done chan bool) {
	c <- "我们虚假套也是好起来了"
	data := <-c
	fmt.Println("main goroutine传来：", data)
	done <- true
}
func fun1(c chan<- int) { //可以调用双向通道和只写通道
	c <- 100
	fmt.Println("fun1() over")
}
