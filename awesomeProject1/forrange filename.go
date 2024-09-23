package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	dirname := "C:\\Users\\August\\Desktop\\LedToggle2.1"
	listFiles(dirname, 0)
}
func listFiles(dir string, level int) {
	//用level来记录我们递归了几次，生成tree样式的目录格式
	s := "|——"
	for i := 0; i < level; i++ {
		s = "|  " + s
	}
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, fileInfo := range fileInfos {
		filename := dir + "/" + fileInfo.Name()
		fmt.Printf("%s%s\n", s, filename)
		if fileInfo.IsDir() { //如果目录下的文件还是目录，就把目录下的文件也打印
			listFiles(filename, level+1) //递归打印
		}
	}

}
