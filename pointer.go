package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a)
	var p *int
	p = &a                             //p等于a的地址，*p等于p的指的地址的值
	fmt.Println(*p)                    //10
	fmt.Println(p)                     //0xc00000a0c8
	fmt.Printf("%T,%p,%p\n", p, p, &p) //*int,0xc00000a0c8,0xc000058030,p的类型，p指的地址，p自己的地址
	var p1 **int
	p1 = &p          //p1是p的指针，存的是
	fmt.Println(*p1) //0xc00000a0c8
	fmt.Println(p1)  //0xc000058030
	/*
		数组指针：存储的是指针，指向一个数组
		指针数组：首先是一个数组，存储的类型是指针
	*/
	//1.创建普通数组
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1)
	//2.创建数组指针
	var p2 *[4]int
	p2 = &arr1
	fmt.Println(p2)                       //&[1 2 3 4]
	fmt.Printf("%T,%p,%p\n", p2, p2, &p2) //*[4]int,0xc000014200,0xc000058038,类型，数组arr1的地址，p2指针自己的地址
	//3.根据数组指针，操作数组
	(*p2)[1] = 100  //&[1 100 3 4]
	p2[0] = 200     //简化写法
	fmt.Println(p2) //&[200,2,3,4]
	//4.指针数组
	b := 1
	c := 2
	d := 3
	e := 4
	arr2 := [4]int{b, c, d, e}
	arr3 := [4]*int{&b, &c, &d, &e}
	fmt.Println(arr2) //[1 2 3 4]
	fmt.Println(arr3) //[0xc00000a148 0xc00000a150 0xc00000a158 0xc00000a160]
	*arr3[0] = 100
	fmt.Println(b) //100
	/*
		函数指针：指向函数的指针
		function默认为一个指针，不加*
	*/
	var f func()
	f = fun5 //函数是一个指针
	f()      //f()==fun1()
	arr4 := fun2()
	fmt.Printf("arr4:类型%T,地址%p,数值%v\n", arr4, &arr4, arr4) //arr4:类型[4]int,地址0xc0000142a0,数值[1 2 3 4]
	//函数内的arr被销毁
	arr5 := fun3()
	fmt.Printf("arr5:类型%T,地址%p,存的地址：%p,数值%v\n", arr5, &arr5, arr5, arr5) //arr5:类型*[4]int,地址0xc000058040,存的地址：0xc000014300,数值&[5 6 7 8]
	//函数内的arr没有被销毁，因为arr5==fun3()返回值的地址
	/*
		指针作为参数：
		参数传递：值传递，引用传递
	*/
	g := 10
	fmt.Println("fun6()函数调用前，g:", g) //fun6()函数调用前，g: 10
	fun6(g)
	fmt.Println("fun6()函数调用后g:", g) //fun6()函数调用后g: 10
	fun7(&g)
	fmt.Println("fun7()函数调用后，g:", g) //fun7()函数调用后，g: 200
	arrexe := [4]int{1, 2, 3, 4}
	funexe1(&arrexe)
	fmt.Println(arrexe)
	funexe2(arr1)

}
func fun3() *[4]int {
	arr := [4]int{5, 6, 7, 8}
	return &arr
}
func fun2() [4]int {
	arr := [4]int{1, 2, 3, 4}
	return arr
}
func fun5() {
	fmt.Println("fun1()")
}
func fun6(num int) { //值传递
	num = 100
	fmt.Println("num在函数中被修改")
}
func fun7(p1 *int) { //引用传递
	fmt.Println("fun7()函数中，p1：", *p1) //fun7()函数中，p1： 10
	*p1 = 200                         //p1在函数结束后会被销毁，但是因为p1指的地址为输入进来的实参的值，所以被传入的实参（g)的值被指针改变，销毁指针p1的内存地址，不销毁实参的地址
	fmt.Println("fun7()函数中，p1：", *p1) //fun7()函数中，p1： 200
}
func funexe1(p1 *[4]int) {
	fmt.Println("函数中的p1", p1)  //&[1 2 3 4]
	fmt.Println("函数中的p1", *p1) //[1 2 3 4]
	p1[0] = 200
}
func funexe2(arr [4]int) {
	fmt.Println("函数中的arr", &arr) //&[200 2 3 4]
	fmt.Println("函数中的arr", arr)  //[200 2 3 4]
	arr[1] = 99
}
