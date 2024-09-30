package main

import "fmt"

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
	//我们也可以通过range来访问通道，这样就不需要ok值了
	go sendData(chan1)
	for v := range chan1 {
		fmt.Println("v==", v)
	}
	fmt.Println("main over")
}
func sendData(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}
