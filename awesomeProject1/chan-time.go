package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		定时器，让用户自己定义自己的超时逻辑，尤其是应对多个select处理多个channel的超时、单channel读写的超时等情形
		t:=time.NewTimer(d),d：定时时间，什么时候开始
		t:=time.AfterFunc(d,f)，f:出发动作
		c:=time.New.After(d)
	*/
	/*
		func NewTimer(d Duration)*Timer
		创建一个计时器，d时间后触发，里面是个结构体，结构体里面是个通道
	*/
	/*
		timer := time.NewTimer(3 * time.Second) //声明一个结构体做计时器，3秒后触发
		fmt.Printf("类型：%T\n", timer)            //类型：*time.Timer
		fmt.Println(time.Now())                 //现在的时间
		ch1 := timer.C
		fmt.Println(<-ch1) //3秒后的时间，所以要等三秒才出发打印
	*/
	//小示例
	/*
		timer2 := time.NewTimer(5 * time.Second)
		go func() {
			<-timer2.C
			fmt.Println("Timer2 over...start:")
		}()
		//time.Sleep(2 * time.Second)
		time.Sleep(5 * time.Second) //Timer2 over...start:
		flag := timer2.Stop()
		if flag {
			fmt.Println("Timer2 over")
		}

	*/
	/*
			解释一下：我们启动上面的goroutine的时候，会等待timer2
		这个通道发出信号，触发后（也就是5秒后），这个goroutine打印Timer2 over...start
		但是time.Sleep(2 * time.Second)会导致主goroutine等待2秒，期间子goroutine仍然在等待
		flag是用来停止计时器的，如果定时器没触发（也就是没到五秒）Stop返回true，并且截止定时器（不管触发没都截止）
		但是如果在截止前就触发定时器了，就返回false

		如果时间一样则都不打印，因为这样的意思就是flag会返回false，所以不执行if true下的代码
		子goroutine因为计时器已经触发，也不会打印子goroutine下的代码
	*/
	/*
			func After(d Duration)<-chan Time,d时间后从通道获取数据
			返回值是一个通道，存储的是d时间间隔后的当前时间
			func After(d Duration) <-chan Time {
			return NewTimer(d).C
		}

	*/
	ch := time.After(5 * time.Second)
	fmt.Printf("ch==%T\n", ch) //ch==<-chan time.Time
	fmt.Println(time.Now())    //打印当前时间
	time2 := <-ch
	fmt.Println(time2) //等五秒，五秒之后的时间
}
