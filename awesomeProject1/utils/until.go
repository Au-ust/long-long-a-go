package utils

import "fmt"

func Count() {
	fmt.Println("我是utils包下的util中的Count()函数")
}
func init() {
	fmt.Println("我是另一个init()函数")
}
func init() {
	fmt.Println("我是一个init()函数，负责main前的初始化")
}

/*
1.init()用于初始化信息，main()作为程序的入口
2.init()定义时不能有参数，不能有返回值，由go程序自己调用
3.init()函数可以定义在任意的包中，可以有多个。main()函数只能在main()包下，并且只能有一个
2.
init函数调用有顺序，随包在工程下的目录的顺序改变而改变
若包和包有依赖关系，eg:A->B->C->D,则沿D->C->B->A的顺序执行
不存在依赖则按照导入main()包的顺序执行init()
避免出现循环import，eg:A->B->C->A
初始化只执行一次

*/
