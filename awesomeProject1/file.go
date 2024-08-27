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
}
