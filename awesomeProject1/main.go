package main

import (
	_ "awesomeProject1/utils" //前面加”_"可以只使用导入包的初始化函数
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s1 := time.Now().Format("2006年1月2日15:04:05") //模板，将现在的时间套进模板里
	s2 := "2005年10月11日13:04:05"
	t2, err := time.Parse("2006年01月02日15:04:05", s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t2) //2005-10-11 13:04:05 +0000 UTC,string->time

	fmt.Println(s1) // 2024年6月13日20:22:45
	//时间戳：1974.1.1距离指定日期的时间差值
	t3 := time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC)
	timeStamp := t3.Unix()
	fmt.Println(timeStamp) //3600
	t4 := time.Now()
	timeStamp2 := t4.Unix()
	fmt.Println(timeStamp2) //1718433515
	//时间间隔
	t5 := t4.Add(time.Minute)
	fmt.Println(t4)                     //2024-06-15 15:13:23.2724067 +0800 CST m=+0.009575901,t5比t4多了一分钟
	fmt.Println(t5)                     //2024-06-15 15:14:23.2724067 +0800 CST m=+60.009575901
	fmt.Println(t5.Add(24 * time.Hour)) //2024-06-16 15:14:23.2724067 +0800 CST m=+86460.009575901
	//睡眠
	time.Sleep(3 * time.Second) //让当前程序进入睡眠状态
	fmt.Println("sleep over")
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10) + 1
	fmt.Println(randNum)
	time.Sleep(time.Duration(randNum) * time.Second) //转换成duration类型
	fmt.Println("sleep over")

}
