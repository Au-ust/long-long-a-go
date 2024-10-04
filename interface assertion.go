package main

import (
	"fmt"
	"math"
)

func main() {
	var t1 Triangle = Triangle{3, 4, 5}
	fmt.Println(t1.perimeter()) //12
	fmt.Println(t1.area())      //6
	var c1 Circle = Circle{5}
	fmt.Println(c1.perimeter()) //31.41592653589793
	fmt.Println(c1.area())      //78.53981633974483
	var s1 Shape                //Shape类型，但是没有确定为Triangle类型还是Circle类型
	s1 = t1                     //定义为Triangle类型
	fmt.Println(s1.perimeter())
	fmt.Println(s1.area())
	//fmt.Println(s1.a,s1.b,s1.c),打印不了，因为s1是Shape接口类型
	testShape(t1) //周长 12 面积 6
	testShape(c1) //周长 31.41592653589793 面积 78.53981633974483
	getType(t1)   //是三角形,三边为 3 4 5
	getType(c1)   //是圆形半径为 5
	getType(s1)   //是三角形,三边为 3 4 5
	var t2 *Triangle = &Triangle{3, 4, 5}
	getType(t2) //ins:*main.Triangle 0xc000058030
	//s:*main.Triangle 0xc0000260a0,地址不一样，值传递
	getType1(t2) //三角形结构体指针 3 4 5
}
func getType(s Shape) {
	//断言，判断接口对象是不是对应的实际类型
	//1.if分支
	if ins, ok := s.(Triangle); ok {
		fmt.Println("是三角形,三边为", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(Circle); ok {
		fmt.Println("是圆形半径为", ins.radius)
	} else if ins, ok := s.(*Triangle); ok {
		fmt.Printf("ins:%T %p\n", ins, &ins)
		fmt.Printf("s:%T %p\n", s, &s)
	} else {
		fmt.Println("我也不知道")
	}
}
func getType1(s Shape) {
	//2.switch分支
	switch ins := s.(type) {
	case Triangle:
		fmt.Println("是三角形,三边为", ins.a, ins.b, ins.c)
	case Circle:
		fmt.Println("是圆形半径为", ins.radius)
	case *Triangle:
		fmt.Println("三角形结构体指针", ins.a, ins.b, ins.c)
	}
}
func testShape(s Shape) {
	fmt.Println("周长", s.perimeter(), "面积", s.area())
}

// 1.定义一个接口
type Shape interface {
	perimeter() float64
	area() float64
}

// 2.定义实现类
type Triangle struct {
	a, b, c float64
}

func (t Triangle) perimeter() float64 {
	return t.a + t.b + t.c
}
func (t Triangle) area() float64 {
	p := t.perimeter() / 2
	s := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	return s
}

type Circle struct {
	radius float64
}

func (c Circle) perimeter() float64 {
	return math.Pi * c.radius * 2
}
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
