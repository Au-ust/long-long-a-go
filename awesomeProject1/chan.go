package main

import (
	"fmt"
	"strconv"
)

func main() {
	var c chan bool
	c = make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子goroutine的i=", i)
		}
		c <- true //把true写入c
	}()
	data := <-c
	fmt.Println("main->", data)
	//fmt.Println("main->", <-c)接收两次数据会死锁，因为子goroutine只向c写入了一次数据，永远满足不了再读，所以死锁
	fmt.Println("main over")
	/*
		有读就得有写，有写就得有读，不然就会死锁
		就好像互斥和读写锁，有锁就有开，有开就有锁
	*/
	go func() {
		data = <-c
	}()
	c <- false
	fmt.Println("data=", data)
	//c <- truefatal //error: all goroutines are asleep - deadlock!
	//嘻嘻嘻
	/*
			通道是有读就有写的，那么我们怎么知道 通道里面有数据需要我们读/写呢？
			我们之前学map的时候，学过ok-idom值：
		/*如何才能知道自己对应键位存的是"0"值，还是空呢？

			在map的返回中除了value值，还有"ok-idiom",代表

		v4, ok := map1[4]
		if ok == true {
			fmt.Println(v4)
		}
		channel也可以哦
	*/
	var chan1 chan int
	chan1 = make(chan int)
	/*
		go sendData(chan1)
		//while循环一直判定chan1里到底有没有对象可读
		for {
			v, ok := <-chan1
			if !ok {
				fmt.Println("读完了,ok==", ok)
				break
			}
			fmt.Println("chan1读取的数据是=", v, "ok==", ok)
		}

	*/
	//我们也可以通过range来访问通道，这样就不需要ok值了
	go sendData(chan1)
	for v := range chan1 {
		fmt.Println("v==", v)
	}
	fmt.Println("main over")
	/*
		普通的通道没有缓冲区，或者称为缓冲区为0，有缓冲区的通道就是缓冲通道
		普通的通道是一次发送， 在接受前都是阻塞的；一次接收，在发送前都是阻塞的
		而缓冲通道有一个缓冲区，在缓冲区满了之后才会阻塞
		非缓冲：ch:=make(chan type)
		缓冲：ch:=make(chan type,capcity)
	*/
	chan2 := make(chan int) //非缓冲,len= 0 cap= 0
	fmt.Println("len=", len(chan2), "cap=", cap(chan2))
	//chan2 <- 100                                        //阻塞式，需要有goroutine读，否则死锁
	chan3 := make(chan int, 5)                          //非阻塞式
	fmt.Println("len=", len(chan3), "cap=", cap(chan3)) //len= 0 cap= 5
	chan3 <- 100
	fmt.Println("len=", len(chan3), "cap=", cap(chan3)) //len= 1 cap= 5
	chan3 <- 200
	chan3 <- 300
	chan3 <- 400
	chan3 <- 500
	fmt.Println("len=", len(chan3), "cap=", cap(chan3)) //len= 5 cap= 5,此时已经满了，进入阻塞
	//chan3 <- 600                                        //死锁啦
	close(chan3) //满了就关闭
	for v3 := range chan3 {
		fmt.Println("chan3==", v3) //像一个队列，先进先出
	}
	fmt.Println("--------------------我是一条分界线----------------")
	//写一个缓存channel
	chan4 := make(chan string, 5)
	done4 := make(chan bool)
	go sendData4(chan4, done4)
	<-done4
	for {
		v4, ok := <-chan4
		if !ok {
			fmt.Println("读完了,ok==", ok)
			break
		}
		fmt.Println("读取的数据是", v4)
	}
	fmt.Println("main over")
}
func sendData(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Println("输入成功")
	}
	close(c)
}
func sendData4(c chan string, done chan bool) {
	for i := 0; i < 10; i++ {
		c <- "数据" + strconv.Itoa(i)
		fmt.Println("这是第", i, "个数据")
	}
	close(c)     // 关闭通道
	done <- true // 通知主线程任务完成
}
