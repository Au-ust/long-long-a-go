package main

import (
	"fmt"
)

func main() {
	/*
					接口：interface
					当某个类型为这个接口中所有的方法提供了方法的实现，即为接口
					go语言中，接口类型的实现关系，是非侵入式：只实现，不用在乎是谁实现的
					java中要显示定义
				eg:
				class Mouse implements USB{}

			一个对象受到类型的限制，可以定义为父类类型，也可以定义为子类，类似于界门纲目科属种的关系
								类型不同，能够访问的字段（功能）不同
								go语言通过接口模拟多态，就是一个接口的实现
						 			1.看成实现本身的类型，能够访问实现类中的属性和方法
									2.看成是对应接口的类型，只能访问接口中定义的方法
						接口用法：
					1.一个函数如果接受接口类型作为残念书，实际上也可以传入该接口任意实现类对象作为参数
					例如在usb中可以直接传入f1或m1,因为f1和m1是usb接口的实现类型对象
					2.定义一个类型为接口类型，实际上可以赋值为任意实现类的对象
					eg:var arr[3]USB(接口类型的数组)
		鸭子类型：符合要求就代表你属于该类型
	*/
	//1.创建mouse类型
	m1 := Mouse{"罗技小红"}
	fmt.Println(m1)
	//2.创建FlashDisk
	f1 := FlashDisk{"闪速64G"}
	fmt.Println(f1)
	testIterface(m1)
	testIterface(f1)
	//Mouse.start
	//Mouse.end
	//FlashDisk.start
	//FlashDisk.end
	var usb USB
	usb = m1
	usb.start()
	usb.end()
	usb = f1
	usb.start()
	usb.end()
	//Mouse.start
	//Mouse.end
	//FlashDisk.start
	//FlashDisk.end
	var arr [3]USB
	arr[0] = m1
	arr[1] = f1
	fmt.Println(arr) //[{罗技小红} {闪速64G} <nil>]
	/*空接口
	例如fmt包下的函数
	type Formatter interface {
		Format(f State, verb rune)
	}
	*/
	var a1 A = Dog{"黑犬"}
	var a2 A = People{"纸包鱼", 30}
	fmt.Println(a1, a2) //{黑犬} {纸包鱼 30}
	test1(a1)           //{黑犬}
	test1(a2)           //{纸包鱼 30}
	test2(a1)           //空接口 {黑犬}
	test2(a2)           //空接口 {纸包鱼 30}
	//空接口代表任意类型的使用
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, a1, a2, 100, "abc") //[{黑犬} {纸包鱼 30} 100 abc]
	fmt.Println(slice1)
	//接口嵌套->多继承
	var dog Dog = Dog{"青色"}
	dog.test11()
	dog.test22()
	dog.test33()
	fmt.Println("------分隔符------")
	var aa1 AA = dog //AA接口类型
	aa1.test11()
	fmt.Println("------分隔符------")
	var bb1 BB = dog //BB接口类型
	bb1.test22()
	fmt.Println("------分隔符------")
	var cc1 CC = dog //CC接口类型
	cc1.test11()
	cc1.test22()
	cc1.test33()
}

type AA interface {
	test11()
}
type BB interface {
	test22()
}
type CC interface {
	AA
	BB
	test33()
}

func (d Dog) test11() {
	fmt.Println("test11...")
}
func (d Dog) test22() {
	fmt.Println("test22...")
} //如果test11和test22都能实现，则Dog同时是AA和BB的实现
func (d Dog) test33() {
	fmt.Println("test33...")
}

func test1(a A) {
	fmt.Println(a)
}

// 空接口代表了任意类型
func test2(a interface{}) {
	fmt.Println("空接口", a)
}

// 1.定义接口
type USB interface {
	start()
	end()
}

// 2.实现类
type Mouse struct {
	name string
}
type FlashDisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Println("Mouse.start")
}
func (m Mouse) end() {
	fmt.Println("Mouse.end")
}
func (f FlashDisk) start() {
	fmt.Println("FlashDisk.start")
}
func (f FlashDisk) end() {
	fmt.Println("FlashDisk.end")
}

// 3.测试方法
func testIterface(usb USB) { //usb=m1 usb=f1
	usb.start()
	usb.end()
}

// 空接口
type A interface {
}
type Dog struct {
	color string
}
type People struct {
	name string
	age  int
}
