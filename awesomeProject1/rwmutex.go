package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex

var wg2 *sync.WaitGroup //包内重复定义了

func main() {
	rwMutex = new(sync.RWMutex)
	wg2 = new(sync.WaitGroup)
	wg2.Add(2)
	go rData(1)
	go rData(2)
	go wData(3)
	wg2.Wait()
	fmt.Println("done")
}
func wData(n int) {
	defer wg2.Done()
	fmt.Printf("goroutineNum==%d,开始写：write start\n", n)
	rwMutex.RLock()
	time.Sleep(5 * time.Second)
	fmt.Printf("goroutineNum==%d正在写入数据。。。\n", n)
	rwMutex.RUnlock() //写操作解锁
	fmt.Println("goroutineNum==", n, "写操作结束")
}
func rData(n int) {
	defer wg2.Done()
	fmt.Printf("goroutineNum==%d,开始读：read start\n", n)
	rwMutex.RLock() //读操作上锁
	fmt.Printf("goroutineNum==%d正在读取数据。。。\n", n)
	time.Sleep(5 * time.Second)
	rwMutex.RUnlock() //读操作解锁
	fmt.Println("goroutineNum==", n, "读操作结束")
} //可以同时有很多读，但是写只能同时有一个
