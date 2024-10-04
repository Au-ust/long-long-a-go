package main

import (
	"bufio"
	"fmt"
	"os"
)

// 这个定义为全局变量，但是不能用简短定义定义全局变量
func main() {
	//变量不用会报错
	/*声明，定义
	1.var,变量名，数据类型
	变量名=赋值
	var 变量名 数据类型 =赋值
	2.类型推断
	var 变量名=赋值
	var sum=10
	3。简短声明
	sum :=100
	冒号左边至少有一个未赋值的新变量
	eg:var a,b
	a,b,sum:=1,2,3,√
	var a,b
	a,b:=1,2,×
	访问，赋值，取值
	1，根据变量名访问2，地址
	*/
	//1，定义变量赋值
	var num1 int
	num1 = 30
	fmt.Printf("num1=%d\n", num1)
	/*静态语言：强类型语言，go,java,cpp都为静态语言
	动态语言：JavaScript，python，php
	*/
	//2，类型推断
	var name = "go"
	fmt.Printf("类型是%T，name=%s\n", name, name)
	//简短定义，省略var
	sum := 100
	fmt.Println(sum)
	//多个变量同时定义
	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Println(a, b, c)
	var m, n int = 100, 100
	fmt.Println(m, n)
	var n1, f1, s1 = 20, 13.14, "august"
	fmt.Println(n1, f1, s1)
	var (
		StudentName = "august"
		age         = 19
		sex         = "女"
	)
	fmt.Printf("StudentName=%s\nage=%d\nsex=%s\n", StudentName, age, sex)
	fmt.Println(StudentName, age, sex)
	//两种都可以
	var s2 []int           //[]切片类型，包含三个字段：地址，长度，容量
	fmt.Println(s2 == nil) //s2的地址为nil,称为nil切片
	/*
		字符串类型：
		概念：多个byte的集合
		语法：使用双引号“”或者``
	*/
	//区别‘A’和“A”
	v1 := "A"
	v2 := 'A'
	fmt.Printf("%T,%s\n", v1, v1)
	fmt.Printf("%T\n", v2) //‘A’为int8，"A"`A`是string
	//编码表：acsii,gbk,unicod(统一全世界）go语言为UTF-8，统一ascii
	/*转义字符‘\'\
	转义有其他用处的符号
	改变符号的意思
	\t制表符（空格）
	*/

	fmt.Println("\"helloworld\"")
	fmt.Println(`"helloworld"`) //变成字符串打印

	/*强制类型转换（像c)
	go语言为静态语言，定义，赋值，运算保持一致
	格式：类型（变量）
	eg:var a int8
	int16(a)
	*/
	var x int8
	x = 10
	var y int16
	y = int16(x)
	fmt.Printf("x=%d, y=%d\n", x, y)
	/*
				转化为二进制，按位操作
				按位&：对应位置都为1才为1，有一个0就为0
				按位|：对应位置都为0才为0，有一个1就为1
				异或^：
		1.a^b按位运算，相同为1，不同为0
		2.^a按位取反
		1->0,0->1
		go语言特有:位清空&^
				a&^b对于b上的数值，如果为0，则取a对应位上的数值
				如果为1，则结果位就取0
	*/
	res0 := 8
	res1 := res0 << 3
	fmt.Printf("res1=%d\n", res1)
	res2 := 8
	res3 := res2 >> 2
	fmt.Printf("res3=%d\n", res3)
	/*Print=打印
	Printf=格式化打印，可以打印多种数据类型
	Println=打印之后换行
	*/
	/*
		var u int
		var v float64
		fmt.Println("请输入一个整数，一个浮点类型数")
		fmt.Scanln(&u, &v) //输入一个值以后空格，继续敲另一个值
		fmt.Printf("a==%d,b==%f", u, v)
		fmt.Scanf("%d,%f", &u, &v) //Scanf要输入第一个数字后输入“，”
		fmt.Printf("a==%d,b==%f", u, v)
	*/
	//bufio包
	fmt.Println("输入一个字符串")
	reader := bufio.NewReader(os.Stdin) //从键盘输入，标准输入
	ss, _ := reader.ReadString('\n')    //判断读取结尾
	fmt.Println("输入的数据", ss)
}
