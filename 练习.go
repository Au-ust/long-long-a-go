package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		for i := 23; i <= 58; i++ {
			fmt.Println(i)
		}
	*/
	/*
		var num int
		num = 0
		for i := 1; i <= 100; i++ {
			num += i
		}
		fmt.Println(num)
	*/
	/*
		for i := 1; i <= 100; i++ {
			if i%3 == 0 && i%5 != 0 {
				fmt.Println(i)
			}
		}
	*/
	/*
		for i := 1; i <= 9; i++ {

			for j := 1; j <= i; j++ {
				fmt.Printf("%d * %d = %d ", i, j, j*i)
			}
			fmt.Println("\n")
		}*/
	//随机数
	/*
		num1 := rand.Int()
		fmt.Println(num1)
		for i := 0; i <= 10; i++ {
			num := rand.Intn(10)
			fmt.Println(num)
		}
		rand.Seed(2)
		num2 := rand.Intn(10)
		fmt.Println(num2)
	*/
	t1 := time.Now()
	fmt.Println(t1)
	fmt.Printf("%T\n", t1)
	//1.设置种子数，设置为时间戳
	rand.Seed(time.Now().UnixNano())
	//2.使用生成随机数的函数
	fmt.Println(rand.Intn(100))
	num3 := rand.Intn(46) + 3 //3->46之间
	fmt.Println(num3)

	var pp int
	fmt.Scanf("%d", &pp)
	if pp <= 5 {
		fmt.Printf("%d", pp+2)
	} else {
		fmt.Printf("%d", pp+2-7)
	}
}
