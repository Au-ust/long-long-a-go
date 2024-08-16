package main

import "fmt"

func main() {
	/*
						面向对象：oop
							go语言的结构体嵌套：
				    1.模拟继承性：is-a,一个类继承一个类
					子类可以直接访问父类的属性和方法
					子类可以新增自己的属性和方法
					子类可以重写父类的方法（override，就是将父类已有的方法，重新实现）
							type A struct{
							field
							}
							type B struct{
							A//匿名字段
							}
					2.模拟聚合关系 has-a，一个类拥有一个类
							type C struct{
								field
							}
							type D struct{
								c C//聚合关系
							}
					3.模拟多态
					一个对象受到类型的限制，可以定义为父类类型，也可以定义为子类，类似于界门纲目科属种的关系
					类型不同，能够访问的字段（功能）不同
					go语言通过接口模拟多态，就是一个接口的实现
			 			1.看成实现本身的类型，能够访问实现类中的属性和方法
						2.看成是对应接口的类型，只能访问接口中定义的方法
			接口用法：
		1.一个函数如果接受接口类型作为残念书，实际上也可以传入该接口任意实现类对象作为参数
		例如在usb中可以直接传入f1或m1,因为f1和m1是usb接口的实现类型对象
	*/
	//1.创建父类对象
	p1 := Personoop{"张三", 18}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.age)
	//2.创建子类对象
	s1 := Studentoop{Personoop{"李四", 18}, "XUPT"}
	fmt.Println(s1)
	s2 := Studentoop{Personoop: Personoop{"王五", 18}, school: "XUPT"}
	fmt.Println(s2)
	var s3 Studentoop
	s3.Personoop.name = "赵六"
	s3.Personoop.age = 18
	s3.school = "XUPT"
	fmt.Println(s3) //{{赵六 18} XUPT}
	//Personoop在Studentoopz中为匿名结构体字段，被称为提升字段
	//因此可以如下命名
	s3.name = "ruby"
	s3.age = 16
	fmt.Println(s3) //{{ruby 16} XUPT}
}

// 1.定义父类
type Personoop struct {
	name string
	age  int
}

// 2.定义子类
type Studentoop struct {
	Personoop        //结构体嵌套模拟继承结构,提升字段,可以直接进行访问
	school    string //子类的新增属性

}
