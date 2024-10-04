package main

import "fmt"

// 全局变量不可以简短定义
var num = 100

func main() {
	getSum(1, 2, 3, 4, 5)
	//go语言支持多返回值，不是所有语言都支持
	//1.
	res1, res2 := rectang(5, 3)
	fmt.Println(res1, res2)
	//2.
	res11, res22 := rectang2(4, 6)
	fmt.Println(res11, res22)
	res111, _ := rectang3(5, 7) //"_"下划线充当空白符，表示不要某个参数
	fmt.Println(res111)         //打印res222会报错，因为用"_"代替
	//定义了返回值，就必须返回
	//如果函数定义了返回值，函数中又有分支，每个分支下都应该有return语句被执行
	//return中断整个函数，break中断一个分支
	ans := summ(100)
	fmt.Println(ans)

	/*
		defer函数
		使对应的语句暂停执行，放到最后执行
	*/
	fun1(1111111)
	defer fun1(2222222)
	fun1(3333333)
	defer fun1(4444444)
	/*
			不加defer应为
			1111111
			2222222
			3333333
			4444444
			5555555

			加了之后为
			1111111
			3333333
			4444444
			2222222

			因为defer置后了两个函数的调用，采用栈的结构打印
			先被defer的语句后执行，后被defer的语句先执行
		defer函数:
		在return之前，所有的defer函数会执行完毕

		当执行外围函数中的return语句时，只有其中所有的延迟函数都执行完毕后
		外围函数才会真正返回。
		当外围函数中的代码引发运行恐慌时，只有其中所有的延迟函数都执行完毕后，该运行时恐慌才会真正被扩展至调用函数。
	*/

	//函数的本质
	fmt.Printf("%T\n", fun1) //func(int)
	fmt.Println(fun1)        //0xe4b3c0,函数名对应的函数体地址
	c := fun1                //c被赋值为赋值为fun1的地址
	fmt.Println(c)           //0x8cb400，所以“c"可以被当作函数调用
	c(6)
	fmt.Println(c) //6
	/*
				匿名函数：没有名字，直接在匿名函数后面加上“（)"代表调用
			 定义一个匿名函数后匿名函数直接被调用，一般只能调用一次
			也可以把这个匿名函数赋值给一个变量，这样就可以多次调用
		go语言支持函数式编程：
		1.将匿名函数作为另一个函数的参数，回调参数
		2.将匿名函数作为另一个函数的返回值，可以形成闭包结构
	*/
	func() {
		fmt.Println("我是一个匿名函数")
	}() //直接在匿名函数后面加上“（)"代表调用
	//即使写一个新的一模一样结构的函数，也不是同一个函数，会开辟新的一块内存
	b := func() {
		fmt.Println("我也是一个匿名函数")
	}
	b() //执行b
	func(a, b int) {
		fmt.Println(a + b)
	}(5, 4)
	ans1 := func(a, b int) int {
		return a + b
	}(18, 9) //将18和9的执行结果赋值给ans1
	fmt.Println(ans1) //27
	ans2 := func(a, b int) int {
		return a + b
	} //将匿名函数的地址给ans2
	fmt.Println(ans2) //0x2cb7a0
	/*
			go语言函数为一种数据类型，支持高阶函数：函数把另一个函数作为参数
		fun1(),fun2()
		将fun1函数作为了fun2()这个函数的参数
		fun2：高阶函数
		fun1:回调函数
	*/
	fmt.Printf("%T\n", add)  //func(int, int) int
	fmt.Printf("%T\n", oper) //func(int, int, func(int, int) int) int
	ans11 := add(12, 4)
	fmt.Println(ans11) //16
	ans22 := oper(10, 9, add)
	fmt.Println(ans22) //10 9 0x28b800,fun(10,9)==19
	ans33 := oper(10, 9, sub)
	fmt.Println(ans33) //10-9==1,执行的是函数中的函数，sub
	//闭包
	ans111 := increment() //res111:=fun=内部的匿名函数
	v1 := ans111()
	fmt.Println(v1) //1
	v2 := ans111()
	fmt.Println(v2) //2，第二次调用i不会被销毁
	/*
			闭包（closure):
				一个外层函数中，有内层函数，该内层函数中，会操作外层函数的局部变量（外层函数中的参数，或者外层函数中直接定义的变量），
			并且该外层函数的返回值就是该内层函数，这个内层函数和外层函数的局部变量，统称为闭包结构

					闭包导致的变量逃逸：
							当外层函数的局部变量被内层函数使用时，外层函数的返回值为该内层函数，且该内层函数还会用到该局部变量
							则外层函数的局部变量不被销毁
						eg：
						func increment() func() int { //外层函数
							i := 0 //1.定义了一个局部变量
							//2。定义了一个匿名函数
							fun1 := func() int { //内层函数
								i++
								return i
							}
							return fun1 //3.返回该匿名函数
						}
		局部变量的生命周期会发生变化，正常的局部变量会随着函数的调用创建，函数的结束销毁
		但是闭包结构中的外层函数的局部变量并不会随着外层函数的结束而销毁，因为内层函数还要继续使用
		每新调用一次外层函数，就会有一个新的i与之对应，但旧的仍然继承
	*/

}
func getSum(nums ...int) { //可变参数：确定类型不确定数量
	sum := 0
	for i := 0; i < len(nums); i++ { //可变参数要放在所有参数之后，一个函数里只能有一个可变参数
		sum += nums[i]
		//这里定义的是局部变量
	}
	fmt.Println(sum)
}
func rectang(len, wid float64) (float64, float64) {
	perimeter := 2 * (len + wid)
	area := len * wid
	return perimeter, area
}
func rectang2(len, wid float64) (per float64, are float64) {
	per = 2 * (len + wid)
	are = len * wid //不需要加冒号，函数第一行已经定义过了
	return          //不需要写返回值，函数定义的时候已经标明了
}
func rectang3(len, wid float64) (per1 float64, are1 float64) {
	per1 = 2 * (len + wid)
	are1 = len * wid //不需要加冒号，函数第一行已经定义过了
	return           //不需要写返回值，函数定义的时候已经标明了
}

// 引用传递传递地址，值传递传递复制好的数值
func summ(a int) int {
	if a == 1 {
		return 1
	}
	return summ(a-1) + a

} //简简单单写个递归
func fun1(a int) {
	fmt.Println(a)
}

// 函数的类型：func(参数列表的类型）（返回值列表的类型）
func add(a, b int) int {
	return a + b
}
func oper(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b, fun)
	res := fun(a, b)
	return res
}
func sub(a, b int) int {
	return a - b
}

// 闭包
func increment() func() int { //外层函数
	i := 0 //1.定义了一个局部变量
	//2。定义了一个匿名函数
	fun1 := func() int { //内层函数
		i++
		return i
	}
	return fun1 //3.返回该匿名函数
}
