package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//复制文件本质上就是io操作
	//小文件一次读写就够了，大文件可能需要边读边写
	srcFile := "C:\\Users\\August\\Desktop\\6.17.txt"
	destFile := "C:\\Users\\August\\Desktop\\6.15.txt"
	CopyFile1(srcFile, destFile)
	//拷贝图片也可以
	//io包下有一个已经写好的函数：Copy
	total, err := CopyFile2(srcFile, destFile)
	fmt.Println(total, err)

}
func CopyFile1(srcFile, destFile string) (int, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}
	//关闭
	defer file1.Close()
	defer file2.Close()
	//读写
	bs := make([]byte, 1024, 1024)
	n := -1 //读取的数据量
	total := 0
	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("copy over")
			break
		} else if err != nil {
			return total, err
		}
		total += n
		file2.Write(bs[:n])
	}
	return total, nil
}
func CopyFile2(srcFile, destFile string) (int64, error) { //注意这里是int64
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file1.Close()
	defer file2.Close()
	return io.Copy(file2, file1) //这个函数的返回值正好是外层函数Copy的返回值
}

// ioutil包
func CopyFile3(srcFile, destFile string) (int, error) {
	bs, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return 0, err
	}
	err = ioutil.WriteFile(destFile, bs, 0777)
	if err != nil {
		return 0, err
	}
	return len(bs), nil
}
