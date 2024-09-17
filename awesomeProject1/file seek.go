package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//seek方法：断点续传（自己设定读哪一部分）
	/*
		Seek(offset,int64,whence int)(int64,error)设置指针光标的位置
		第一个参数：偏移量
		第二个参数：如何设置:
						0:seekstart,表示相对于文件开始
						1.seekcurrent,表示相对于当前光标位置
						2.seekend，相对于结尾
	*/
	fileName := "C:\\Users\\August\\Desktop\\6.17.txt"
	file, err := os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err) //log.Fatal记录日志并调用 os.Exit(1) 终止程序。
	}
	defer file.Close()
	//读写
	bs := []byte{0}
	file.Read(bs)
	fmt.Println(string(bs))
	file.Seek(4, io.SeekStart) //读取距离开头的第四个字符
	file.Read(bs)
	fmt.Println(string(bs))
	/*
		const (
			SeekStart   = 0 // seek relative to the origin of the file
			SeekCurrent = 1 // seek relative to the current offset
			SeekEnd     = 2 // seek relative to the end
		)
	*/
	file.Seek(6, 1)
	file.Read(bs)
	fmt.Println(string(bs))
	file.Seek(0, io.SeekEnd) //在文件末尾写
	file.WriteString("wertyui")
	//fmt.Println(string(bs))
	//断点续传
	/*
				复制文件的时候机器突然断电了，再次启动的时候接着上次的继续传
				创建临时文件，存断电前存入的字节数
			"C:\Users\August\Pictures\锁屏22.jpg"
		边复制，边记录复制总量
	*/
	srcFile := "C:\\Users\\August\\Pictures\\锁屏22.jpg"
	destFile := srcFile[strings.LastIndex(srcFile, "\\")+1:]
	fmt.Println(destFile)
	tempFile := strings.TrimSuffix(destFile, ".jpg") + "_temp.txt" //去除字符串末尾指定的后缀。如果原字符串以给定的后缀结尾
	// 那么该函数会返回去掉后缀后的字符串；如果不以该后缀结尾，则返回原字符串。

	//tempFile := destFile + "temp.txt"
	fmt.Println(tempFile)

	file1, err := os.Open(srcFile)
	file2, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	file3, err := os.OpenFile(tempFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file1.Close()
	defer file2.Close()
	//defer file3.Close()
	//1.先读取临时文件中的数据，再seek
	file3.Seek(0, io.SeekStart)
	bs1 := make([]byte, 100, 100)
	n1, err := file3.Read(bs1)
	//HandleErr(err) //这行出错了是因为我们的临时文件里没有东西可读
	contStr := string(bs1[:n1])
	count, err := strconv.ParseInt(contStr, 10, 64) //将字符串转化为指定进制的整数
	//HandleErr(err)//这行出错也是因为临时文件里没有东西
	fmt.Println(count)
	//设置我们被打断后继续读/写的位置
	file1.Seek(count, io.SeekStart)
	file2.Seek(count, io.SeekStart)
	data := make([]byte, 1024, 1024)
	n2 := -1            //读入的数据量
	n3 := -1            //写出的数据量
	total := int(count) //读取的总量
	//复制文件
	for {
		n2, err = file1.Read(data)
		if err == io.EOF || n2 == 0 {
			fmt.Println("copy over")
			file3.Close()
			os.Remove(tempFile)
			break
		}
		n3, err = file2.Write(data[:n2])
		total += n3
		//将复制的总量存到temp中
		file3.Seek(0, io.SeekStart)
		file3.WriteString(strconv.Itoa(total))
		fmt.Printf("total==%d\n", total)
		//模拟复制的时候断电
		//if total > 6000 {
		//	panic("假装断电了")
		//}

	}
}

// err判断函数
func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
