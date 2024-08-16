package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
			结构体：不同类型数据集合
		结构体成员是由一系列的成员变量构成，这些成员变量也被称为“字段”
		放弃了大量面向对象的特性，首字母大写可共用，小写为私用
	*/
	//定义结构体，方法1
	var p1 Person
	fmt.Println(p1)
	p1.name = "XUPT"
	p1.sex = "无性别"
	p1.age = 74
	fmt.Println(p1) //{XUPT 74 无性别}
	//定义结构体，方法2
	p2 := Person{}
	p2.age = 244
	p2.name = "艾美莉卡"
	p2.sex = "沃尔玛购物袋"
	fmt.Println(p2) //{艾美莉卡 244 沃尔玛购物袋}
	//方法3
	p3 := Person{name: "荷叶饭",
		sex: "女",
		age: 18, //加逗号
	}
	fmt.Println(p3) //{荷叶饭 18 女}
	//方法4
	p4 := Person{
		"米库",
		16,
		"女",
	}
	fmt.Println(p4) //{米库 16 女}
	//1.结构体为值类型，默认深拷贝,往往用指针操纵
	fmt.Printf("%T\n", &p4) //*main.Person
	fmt.Printf("%T\n", p1)  //main.Person
	//2.定义结构体指针
	var pp1 *Person
	pp1 = &p1
	fmt.Println(pp1)
	fmt.Printf("%p,%T\n", pp1, pp1) //0xc0001040c0,*main.Person
	fmt.Println(*pp1)               //{XUPT 74 无性别}
	//使用内置函数new(),go中专门用来创建某种类型的指针的函数
	pp2 := new(Person)
	fmt.Printf("%p,%T\n", pp2, pp2) //0xc0000202d0,*main.Person,指针类型
	//(*pp2).name="太宰治”也可以
	pp2.name = "zyzy"
	pp2.age = 22
	pp2.sex = "男"
	fmt.Println(pp2) //&{zyzy 22 男}
	//匿名结构体
	s1 := Student{name: "无敌原神大王", age: 18}
	fmt.Println(s1.name, s1.age)
	func() {
		fmt.Println("hello world")
	}()
	s2 := struct {
		name string
		age  int
	}{
		name: "无敌高考大王",
		age:  18,
	}
	fmt.Println(s2.name, s2.age)
	w2 := Worker{
		"无敌崩坏星穹铁道大王",
		18,
	}
	fmt.Println(w2.string, w2.int) //打印匿名字段结构体
	bb1 := Book{}
	bb1.bookName = "1984"
	bb1.price = 45.8
	ss1 := Studentbook{}
	ss1.name = "奥威尔"
	ss1.age = 24
	ss1.book = bb1
	fmt.Println(bb1)
	fmt.Println(ss1)
	ss2 := Studentbook{
		name: "code",
		age:  18,
		book: Book{
			"西西弗神话",
			24.9,
		},
	}
	fmt.Println(ss2)
	bb3 := Book{
		bookName: "三体",
		price:    14.9,
	}
	ss3 := Studentbook2{
		name: "刘慈欣",
		age:  18,
		book: &bb3,
	}
	fmt.Println(bb3)                               //{三体 14.9}
	fmt.Println(ss3)                               //{刘慈欣 18 0xc000008060}
	fmt.Println(ss3.book.bookName, ss3.book.price) //三体 14.9
	/*
			type关键字：用于类型定义和类型别名
		1.类型定义：type 类型名 Type
		2.类型别名：type 类型名=Type
	*/
	var i1 myint        //int
	var i2 = 100        //int
	fmt.Println(i1, i2) //0 100
	//i2!=i1,i1,i2是两种语法类型
	var str1 mystr
	str1 = "嘻嘻"
	fmt.Printf("%T\n", str1) //main.mystr
	res1 := fun11()
	fmt.Println(res1(10, 20)) //拼接10和20
	type myint2 = int         //给int起了别名，就像define
}

// 定义一个新的函数类型
type myfun func(int, int) string

func fun11() myfun {
	fun := func(a, b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
	return fun
}

type myint int
type mystr string
type Person struct {
	name string
	age  int
	sex  string
}
type Student struct {
	name string
	age  int
}
type Worker struct {
	/*name string
	age  int*/
	string
	int //匿名字段
}

// 嵌套结构体
type Book struct {
	bookName string
	price    float64
}
type Studentbook struct {
	name string
	age  int
	book Book
}
type Studentbook2 struct {
	name string
	age  int
	book *Book //Book的指针
}
