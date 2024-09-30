package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticktes = 10
var mutex sync.Mutex  //定义一个互斥锁的对象
var wg sync.WaitGroup //同步等待组对象

func main() {
	wg.Add(4)
	fmt.Println("陶吉吉演唱会，开始！")
	go TicktesSaler("1号")
	go TicktesSaler("2号")
	go TicktesSaler("3号")
	go TicktesSaler("4号")
	wg.Wait()
	//time.Sleep(1 * time.Second)
	fmt.Println("All Done")
}
func TicktesSaler(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		mutex.Lock() //从这行开始上锁，只有一个goroutine可以执行
		if ticktes > 0 {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			ticktes--
			fmt.Printf("%sTicktesSold卖出，余量：%d\n", name, ticktes) //票数出现负数是因为1号判定完票数>0后进入if分支先睡觉
			//睡觉的时候被别的goroutine抢占了资源，等睡醒的时候票已经为0了，但是因为已经进入>0的分支，所以只能继续执行，变成负数

		} else {
			mutex.Unlock() //解锁，输出售罄需要解锁
			fmt.Printf("%s做不了自己了\n", name)

			break
		}
		mutex.Unlock() //这样就不会有负数啦
	}
}
