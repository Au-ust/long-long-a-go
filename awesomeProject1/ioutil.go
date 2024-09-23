package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	/*
		ioutil包:
		ReadFile()
		WriteFile()
		ReadDir()
	*/
	/*
		filename := "C:\\Users\\August\\GolandProjects\\awesomeProject1\\6.16.txt"
		data, err := ioutil.ReadFile(filename) //已经被废弃，和os包合并
		fmt.Println(string(data), err)
		filename1 := "C:\\Users\\August\\GolandProjects\\awesomeProject1\\6.17.txt"
		//s1 := "hello golang"
		s2 := "bye bye golang"
		err = ioutil.WriteFile(filename1, []byte(s2), os.ModePerm) //会清空已有内容
		fmt.Println(err)
	*/
	//ReadAll什么都可以读，还可以读字符串
	/*
		s3 := "我会永远的视奸你，永远永远"
		r1 := strings.NewReader(s3) //将字符串封装成一个 Reader 对象，使得字符串可以像文件或其他输入流一样逐字节读
		data3, err := ioutil.ReadAll(r1)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(data3))
	*/
	//ReadDir读取目录信息，但是只能读取一层
	/*
		dirName := "C:\\Users\\August\\GolandProjects\\awesomeProject1"
		fileInfo, err := ioutil.ReadDir(dirName)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(len(fileInfo)) //len(fileDir)=14,不加len显示的是文件地址
		for i, file := range fileInfo {
			fmt.Printf("第%d个文件，文件名：%s，是否为目录：%t \n", i, file.Name(), file.IsDir())
		}
	*/
	//创建临时目录和临时文件
	dir, err := ioutil.TempDir("C:\\Users\\August\\GolandProjects\\awesomeProject1", "Test")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dir)
	file, err := ioutil.TempFile(dir, "Test")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file.Name())
	defer os.Remove(file.Name())

}
