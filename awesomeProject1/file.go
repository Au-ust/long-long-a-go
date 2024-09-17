package main

import (
	"fmt"
	"os"
)

func main() {
	/*
			FileInfo:文件信息
		interface
		Name(),文件名
	*/
	fileInfo, err := os.Stat("C:\\Users\\August\\Desktop\\6.15.txt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("fileInfo:%T\n", fileInfo) //fileInfo:*os.fileStat
	//获取文件名
	fmt.Println(fileInfo.Name()) //6.15.txt
	//空间
	fmt.Println(fileInfo.Size()) //0
	//是否为目录
	fmt.Println(fileInfo.IsDir()) //是不是目录，false,不是目录
	//修改时间
	fmt.Println(fileInfo.ModTime()) //2024-06-15 15:42:41.4009077 +0800 CST
	//权限
	fmt.Println(fileInfo.Mode()) //-rw-rw-rw-
	/*
		字符表示：
			r:可读
			w:可写
			x:可执行
			type owner group other
			    -rw   - rw  -  rw -
		八进制表示：
		r->004
		w->002
		x->001
		-->000
		eg:0755->可读可写可执行-可读可执行-可读可执行
	*/
	//i/o操作
	/*
		对于一个程序的输入输出就是i/o操作
		go语言实现io操作的接口由i/o包提供，这些接口包装了一些偏底层的操作
	*/
	//读取文件
	//1.打开文件
	filename := "C:\\Users\\August\\Desktop\\6.15.txt"
	file, err := os.Open(filename) //open:只读模式打开
	if err != nil {
		fmt.Println("err:", err)
	}
	//2.读取数据
	bs := make([]byte, 4, 4) //创建切片读取数据
	n, err := file.Read(bs)
	if err != nil {
		fmt.Println(err) //一开始打印EOF是因为我的文件为空
	}
	fmt.Println(n)
	fmt.Println(bs)
	fmt.Println(string(bs))
	//第二次读取，会接着第一次的继续往后读
	n, err = file.Read(bs)
	if err != nil {
		fmt.Println(err) //一开始打印EOF是因为我的文件为空，直接就读完了
	}
	fmt.Println(n)
	fmt.Println(bs)
	fmt.Println(string(bs))
	//3.关闭文件
	defer file.Close()
	//type Reader interface {
	//}
	//写操作
	//1.打开
	filename1 := "C:\\Users\\August\\Desktop\\6.17.txt"
	//2.写
	//file, err = os.Open(filename1)//只读，如果没有该文件就会报错
	file, err = os.OpenFile(filename1, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm) //文件名，创建模式，权限（0777）
	//os.O_APPEND,追加写入；
	//O_WRONLY，只写模式打开
	/*
	  0(0000): 不可读写,不能被执行
	   1(0001): 不可读写,能被执行
	   2(0010): 可写不可读,不能被执行
	   3(0011): 可写不可读,能被执行
	   4(0100): 可读不可写,不能被执行
	   5(0101): 可读不可写,能被执行
	   6(0110): 可读写,不能执行
	   7(0111): 可读写,可执行
	*/
	if err != nil {
		fmt.Println(err)
		return
	}
	bs1 := []byte{65, 66, 67, 68, 69} //A,B,C,D,E
	bs2 := []byte{97, 98, 99}
	n1, err := file.Write(bs2) //func (f *File) Write(b []byte) (n int, err error)
	// 写入成功则返回写入的字节数，err 返回 nil，写入失败返回对应的错误信息。
	//也可以不把切片所有的东西写进去，也可以只写一部分
	//示例：
	n2, err := file.Write(bs1[:2]) //截取bs1的前两个字符：65，66（A,B）
	//write还有别的写的操作：
	//直接写出字符串
	/*type StringWriter interface{
	WriteString(s string)(n int,err error)
	}
	*/
	n3, err := file.WriteString("hello write")
	//其实这个WriteString相当于写一个byte切片进去
	n4, err := file.Write([]byte("你好"))
	fmt.Println(n4)
	HandleErr(err)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	HandleErr(err)
	//3.关闭
	defer file.Close()
}

/*
func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}在其他文件下定义过了
*/
