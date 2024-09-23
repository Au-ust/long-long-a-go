package main

import (
	"fmt"
	"time"
)

/*
协程是由用户来调度的，可以创建成千上万个
在函数或方法前加上关键字：go，你会同时运行一个新的goroutine
*/
func main() {
	//go PrintfNum()
	/*for i := 0; i < 100; i++ {
		fmt.Printf("主Goroutine%d:Hello Goroutine!\n", i)
	}
	time.Sleep(1 * time.Second) //主Goroutine结束后无论子Goroutine是否结束，子进程都结束
	//所以让主进程睡1s，等子进程
	*/
	//go alphabtes() //go关键字
	//time.Sleep(3 * time.Second)
	//fmt.Println("over!") //不管运行多少次都是一样的结果，因为我们睡眠时间调整的是固定的
	//按理来说我们的并发是随机的，但是我们使用了sleep，A睡觉的时候B就可以执行
	//不同的goroutine会在对应醒的时刻执行

	//临界资源：可以被多个进程/协程共享的资源
	a := 3
	go func() {
		a = 1
		fmt.Println("子a=", a)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("父a=", a)

}
func PrintfNum() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("|———子GoroutineNum:%d\n", i)
	}
} //了解函数和goroutine的差别
/*
		函数只有执行完后才诞生并返回返回值
		但是goroutine会在调用的那一刻立即返回，他的任何返回值都会被忽略
	主goroutine会给每个goroutine申请空间（32位为250MB,64位为1GB）
	超出最大空间会恐慌（栈溢出）
	然后主goroutine进行初始化：
		1.创建一个特殊的defer语句，在主goroutine退出时做善后处理，因为主goroutine也可能异常退出
		2.启动专门用来在后台清扫内存垃圾的goroutine，并设置GC可用标识

*/
//如果有多个goroutine呢？
func alphabtes() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("|———子GoroutineAlp:%c\n", i)
	}
}
