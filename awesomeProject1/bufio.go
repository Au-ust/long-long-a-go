package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	/*
		bufio的io区别在于多了个缓冲区，因为在数据传输中时间都浪费在程序和硬盘打交道了
		所以创建一个缓冲区，在内存部分，然后把数据（一般缓冲区读取的大小要比程序大）读取或者写入放到缓冲区里
		程序再从缓冲区读
		像匿名管道
	*/
	filename := "C:\\Users\\August\\Desktop\\6.17.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	b1 := bufio.NewReader(file) //返回值：rd,块的大小（默认4096）
	/*
		bufio.Reader高效读取

		p := make([]byte, 1000)     //此处1000<4096
		n1, err := b1.Read(p)
		fmt.Println(string(p))
		fmt.Println(n1)

	*/
	/*
		ReadLine,只读取一行,不建议用
	*/
	/*	data, flag, err := b1.ReadLine()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(flag)
		fmt.Println(string(data))*/
	for {
		s1, err := b1.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读取完毕") //EOF
			break
		}
		fmt.Println(s1)
	}
	//ReadBytes
	data, err := b1.ReadBytes('\n') //读取到第一个\n前
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	//Scanner
	//s2 := ""
	//fmt.Scanln(&s2) //遇到空格就停止读取
	//fmt.Println(s2)
	//所以用bufio
	/*
		b2 := bufio.NewReader(os.Stdin)
		s2, _ := b2.ReadString('\n')
		fmt.Println(s2)
	*/

	//bufio下的io是对io.Writer的封装
	filenameW := "C:\\Users\\August\\GolandProjects\\awesomeProject1\\6.16.txt"
	fileW, err := os.OpenFile(filenameW, os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileW.Close()
	w1 := bufio.NewWriter(fileW)
	n, err := w1.WriteString("hello world")
	fmt.Println(n)
	fmt.Println(err)
	w1.Flush() //冲刷缓冲区

}
