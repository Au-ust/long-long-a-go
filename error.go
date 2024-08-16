package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		if ins, ok := err.(*os.PathError); ok {
			fmt.Println("1.Op:", ins.Op)
			fmt.Println("2.Path:", ins.Path)
			fmt.Println("3.Err:", ins.Err)
			//open test.txt: The system cannot find the file specified.
		}
		return
	}
	fmt.Println(f.Name())
	addr, err := net.LookupHost("www.baidu.com")
	fmt.Println(addr, err)
}
