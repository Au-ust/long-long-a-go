package main

import (
	"fmt"
	"runtime"
)

func init() {
	//逻辑CPU的数量可以设置，设置go下的最大CPU数量：[1，256]
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(n)                               //返回值是上一次设置的数量，12
	fmt.Println("旧的逻辑CPU的数量：", runtime.NumCPU()) //12
	//我们需要在程序运行起来之前旧设置好我们的逻辑CPU的数量，所以可以放到init函数里
}

func main() {
	/*
		包含了运行时与系统交互的操作，例如控制goroutine的函数
	*/
	//获取goroot目录
	fmt.Println("GOROOT——>", runtime.GOROOT()) //goroot的目录位置
	//获取操作系统
	fmt.Println("os/platform:", runtime.GOOS)
	//获取逻辑cpu的数量
	fmt.Println("逻辑CPU的数量:", runtime.NumCPU()) //12
	//逻辑CPU的数量可以设置，设置go下的最大CPU数量：[1，256]
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(n)                               //返回值是上一次设置的数量，12
	fmt.Println("旧的逻辑CPU的数量：", runtime.NumCPU()) //12
	//我们需要在程序运行起来之前旧设置好我们的逻辑CPU的数量，所以可以放到init函数里
	//go1.8以前，这个逻辑cpu是要自己设置的，1.8以后可以不自己设置，默认在多核
	/*
		之前我们让time.sleep来延缓主线程的执行，现在可以使用runtime.Gosched
	*/
	/*go func() {
		for i := 0; i <= 5; i++ {
			fmt.Println("goroutine:", i)
			runtime.Goexit() //停止该goroutine，但是函数里的defer还要执行
		}
		defer fmt.Println("i am defer")
	}()
	*/
	go func() {
		fmt.Println("begin:")
		fun()               //调用fun函数
		fmt.Println("end.") //fun函数里有Goexit（）所以没有打印这句
	}()
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println("main:", i)
	}
}
func fun() {
	defer fmt.Println("i am defer")
	for i := 0; i <= 5; i++ {
		fmt.Println("goroutine:", i)
	}
	runtime.Goexit() //停止该goroutine，但是函数里的defer还要执行

}
