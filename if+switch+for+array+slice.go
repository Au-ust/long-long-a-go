package main

import (
	"fmt"
)

func main() {
	/*条件循环语句：if
	语法格式：
	if 条件表达式{
		//
	}
	*/
	num := 10
	if num > 10 {
		fmt.Println("大于10")
	} else if num < 10 { //else必须在“}”的正右边
		fmt.Println("小于10")
	} else {
		fmt.Println("等于10")
	}
	/*
			if语句的其他写法：
		if 初始化；条件{
		}
	*/
	if res := 9; res > 0 { //if外面访问不了if里面的res，只作用于当前的if
		fmt.Println("res>0")
	}
	/*
				switch 变量名{
				case 数值：分支1
				case 数值：分支2
				......
				default:

			}
		还可以与bool类型结合
		switch(不加变量){
		case true:
		case false:
	*/
	a := 8
	switch a {
	case 1:
		fmt.Println("春天")
	case 2:
		fmt.Println("夏天")
	case 3:
		fmt.Println("秋天")
	case 4:
		fmt.Println("冬天")
	default:
		fmt.Println("不存在")
	}
	/*
		var score int
		fmt.Scan(&score)
		switch {
		case score >= 0 && score < 60:
			fmt.Println("不及格")
		case score >= 60 && score < 70:

			fmt.Println("及格")
		case score >= 70 && score < 80:
			fmt.Println("中等")
			//fallthrough,满足条件后向下穿透，无条件执行下一个case的语句,放在case的最后一行
		case score >= 80 && score < 90:

			fmt.Println("良好")
		case score >= 90 && score <= 100:
			fmt.Println("优秀")
		default:
			fmt.Println("成绩有误")
			}*/
	for i := 1; i <= 5; i++ {
		fmt.Println("helloworld")
	}
	/*
		标准写法
		for 表达式1；表达式2；表达式3{
		}
		其他写法：
			1.省略表达式1和3
			for (;)表达式2(;)（while）分号可不写
			2.省略所有表达式
			for(
			}
			(while(1))死循环
	*/
	//for range 遍历数组
	arr1 := [5]int{1, 2, 3, 4, 5}
	for index, value := range arr1 {
		fmt.Printf("下标是%d,数值是%d", index, value)
	}
	/*数组的数据类型：[长度]类型
	为值类型
	例如array，int，float，string
	可以值传递，可值传递的数据类型为存储数值本身
	传递的数值为备份（副本）

	传递地址，为引用类型
	例如slice(切片），map..
	*/
	/*
				slice:像vector,动态数组
				左闭右开
				定义：var 名字 []
			make函数可以创建引用类型的数据
			make(type,len,cap)类型，长度，容量
		切片一旦扩容，指向一个新的数组地址，容量成倍增长
	*/
	var s1 []int
	fmt.Println(s1)
	s2 := []int{1, 2, 3}
	fmt.Println(s2)
	fmt.Printf("%T", s2)
	s3 := make([]int, 3, 8)
	fmt.Printf("%d,%d\n", cap(s3), len(s3))
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}   //设定初始数组
	ss1 := array[1:9]                                 //1-9，len:8,cap:9
	ss2 := array[:4]                                  //0-4,len:4,10
	fmt.Printf("%p,%d,%d\n", ss1, len(ss1), cap(ss1)) //检查容量
	fmt.Printf("%p,%d,%d\n", ss2, len(ss2), cap(ss2))
	//更改数组内容
	array[4] = 100
	fmt.Println(ss1)
	fmt.Println(ss2) //操作了底层数组
	//更改切片内容
	ss2 = append(ss2, 50, 50, 50, 50) //在是、ss2的后面改变，底层数组也被改变
	fmt.Println(ss2)
	fmt.Println(array) //底层数组也被更改
	//切片是引用类型,数组是值类型
	str := []int{1, 2, 3}
	fmt.Println(str)
	str2 := str
	fmt.Printf("%p\n", str)  //切片指向的数组的地址
	fmt.Printf("%p\n", &str) //切片本身的地址
	fmt.Printf("%p\n", str2)
	fmt.Printf("%p\n", &str2)
	/*
				深拷贝：拷贝值
				浅拷贝：拷贝地址
			内置函数copy,对切片进行深拷贝
		copy(目标切片，原切片[]type)
	*/
	str3 := make([]int, 3)
	copy(str3, str)
	fmt.Println(str3)
	copy(str3, str[1:2])
	fmt.Println(str3)
}
