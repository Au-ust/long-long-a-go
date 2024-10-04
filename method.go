package main

import "fmt"

func main() {
	/*
				方法：method
			定义了接受者的函数，接受者可以是命名类型或者结构体类型中的一个值或者一个指针。
			所有给定类型的方法属于该类型的方法集
			方法名字可以重复

			语法：
			func(接受者)方法名(参数列表)(返回值列表){
			}
		method区别调用者
	*/
	w1 := Worker1{name: "张三", age: 18, sex: "男"}
	w1.work() //张三 在工作
	w2 := &Worker1{name: "李四", age: 18, sex: "女"}
	fmt.Printf("%T\n", w2) //*main.Worker1
	w2.work()              //李四 在工作
	w2.rest()              //李四 在休息
	c1 := Cat{name: "艾露猫", color: "浅棕色", sex: "男"}
	c1.printfInfo() //name: 艾露猫 color: 浅棕色 sex: 男
	c1.name = "呆猫"
	fmt.Printf("%T\n", c1) //main.Cat
	c1.printfInfo()        //name: 呆猫 color: 浅棕色 sex: 男
	w1.printfInfo()        //name: 张三 age: 18 sex 男
	//方法可以重名，因为接受者不同

	//创建Personmethod类型
	p1 := Personmethod{"王二狗", 18}
	fmt.Println(p1) //{王二狗 18}
	p1.eat()        //父类的方法，吃盖浇面
	s1 := Studentmethod{Personmethod{"ruby", 18}, "XUPT"}
	fmt.Println(s1) //{{ruby 18} XUPT}
	s1.study()      //子类的新增方法，学生学习
	s1.eat()        //子类重写的方法，吃炸鸡
}

type Worker1 struct {
	name string
	age  int
	sex  string
}
type Cat struct {
	name  string
	color string
	sex   string
}

func (w Worker1) work() {
	fmt.Println(w.name, "在工作")
}
func (p *Worker1) rest() {
	fmt.Println(p.name, "在休息")
}
func (p *Worker1) printfInfo() {
	fmt.Println("name:", p.name, "age:", p.age, "sex", p.sex)
}
func (p *Cat) printfInfo() {
	fmt.Println("name:", p.name, "color:", p.color, "sex:", p.sex)
}

// 继承中的方法
// 定义一个“父类”
type Personmethod struct {
	name string
	age  int
}

// 定义一个子类
type Studentmethod struct {
	Personmethod
	school string
}

// 定义一个父类方法
func (p Personmethod) eat() {
	fmt.Println("父类的方法，吃盖浇面")
}

// 定义一个子类方法
func (p Studentmethod) study() {
	fmt.Println("子类的新增方法，学生学习")
}
func (p Studentmethod) eat() {
	fmt.Println("子类重写的方法，吃炸鸡")
}
