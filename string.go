package main

import (
	"fmt"
	"strconv"
	"strings"
) //按住ctrl可以查看函数的源代码

func main() {
	/*
		Go中字符串是一个字节的切片
		字符是用unicode存储的，类型是UTF-8
			字符串就像字符组成的数组，但是被分配在<只读>内存下，需要修改可以在byte和string之间类型转换
			中文在UTF-8占三个字节
	*/
	//1.定义字符串
	s1 := "hello world"
	s2 := "hello 狗屎"
	//2.字符串长度
	fmt.Println(len(s1), len(s2))
	//3.获取某个字符
	fmt.Println(s1[0])
	fmt.Printf("%c\n", s1[0])
	//4.字符串遍历
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%c\n", s1[i])
	}
	for i := 0; i < len(s2); i++ {
		fmt.Printf("%c\n", s2[i])
	} //后面乱码是因为中文是三个字节
	for i, v := range s2 {
		fmt.Printf("%d %c\n", i, v) //没有乱码但是下标顺序不连续
	}
	//5.字符串是字节的集合
	slice1 := []byte{65, 66, 67, 68, 69}
	s3 := string(slice1) //根据一个字节切片，构建字符串
	fmt.Println(s3)
	//反过来
	s4 := "abcde"
	slice2 := []byte(s4)
	fmt.Println(slice2) //打印的是数值
	//6.字符串不可修改
	//s3[2]="ww",报错，字符串不可修改

	//函数练习
	str := "世界是一个大妓院aaa"
	fmt.Println(strings.Contains(str, "妓院"))       //true
	fmt.Println(strings.ContainsAny(str, "妓院abc")) //true
	fmt.Println(strings.ContainsAny(str, "狗屎"))    //false，包含其一即为true
	fmt.Println(strings.Count(str, "a"))           //统计a在str出现的次数
	str1 := "你的肛门比较松弛"
	if strings.HasPrefix(str1, "你") { //如果str1是以“你”开头
		fmt.Println("没错,", str1)
	} else {
		fmt.Println("不，他的肛门比较松弛")
	}
	if strings.HasSuffix(str1, "松弛") { //若以“松弛”为结尾
		fmt.Println(str1)
	}
	fmt.Println(strings.Index(str, "a"))      //第一次出现"a"的下标,24
	fmt.Println(strings.LastIndex(str, "a"))  //最后一次出现“a"的下标,26
	fmt.Println(strings.Index(str, "b"))      //没有的话返回-1,-1
	fmt.Println(strings.IndexAny(str, "abc")) //出现”abc”任意一个排在前面的字符字符则返回下标，24
	str2 := []string{"abc", "def", "ghi", "jkl"}
	fmt.Println(strings.Join(str2, "+")) //拼接一个字符串数组中的字符
	str3 := "原神,是一款由\"米哈游\"前面忘了,后面忘了"
	fmt.Println(strings.Split(str3, ","))            //从逗号处断开
	fmt.Println(strings.Repeat(str1, 5))             //自己拼接自己拼接了5次
	fmt.Println(strings.Replace(str, "世界", "肛门", 1)) //把str中的“世界”替换为“肛门”，替换一次
	//如果要全部替换则n取负数
	fmt.Println(strings.Replace(str, "a", "去死", -1))
	str4 := "abcdefg"
	str5 := "ABCDEFG"
	fmt.Println(strings.ToLower(str5)) //大写字母全变成小写，ABCDEFG->abcdefg
	fmt.Println(strings.ToUpper(str4)) //小写字母全变成大写，abcdefg->ABCDEFG
	/*
		其他语言：substring(start,end)截取子串
		go没有substring
		str[start:end]->suber
		左闭右开，包含start不包含end
	*/
	fmt.Println(str[0:6]) //截取str的前六个字节，“世界”
	/*
		strcnov包的使用
	*/
	//1.bool
	sss1 := "true"
	b1, err := strconv.ParseBool(sss1)
	if err == nil {
		fmt.Println(err)
	}
	fmt.Printf("%T,%t\n", b1, b1) //bool,true
	//2.int
	sss2 := "100"
	b2, err := strconv.ParseInt(sss2, 10, 64) //（字符串，进制，最大位数）
	if err == nil {
		fmt.Println(err)
	}
	fmt.Printf("%T,%d\n", b2, b2) //int64,100
	//itoa(),atoi()
	i3, err := strconv.Atoi("-42") //字符串转为int
	fmt.Printf("%T,%d\n", i3, i3)  //int,-42
	i4 := strconv.Itoa(42)
	fmt.Printf("%T,%s\n", i4, i4) //string,42

}
