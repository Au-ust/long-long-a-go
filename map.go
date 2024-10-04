package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		map:映射，键值对的集合，引用类型
		存储特点：存储的是无序的键值对
				键不能重复，与value值一一对应，否则新的会覆盖旧的

		1.map创建

		2.判断map是否为空

		每种数据类型未赋值的初始值：
		int:0
		float:0.0->0
		string:""
		array:[00000]

		slice:nil(空，但是因为有底层数组，所以未赋值也可直接使用
		map:nil(空，但是不可以直接使用
	*/
	var map1 map[int]string                                       //[key的类型]值的类型，没有盘子没有饺子
	var map2 = make(map[int]string)                               //有盘子没有饺子
	var map3 = map[string]int{"GO": 98, "Python": 79, "Html": 93} //有盘子，也装了饺子
	//map1[1] = "hello"会报错
	fmt.Println(map1, map2, map3)
	if map1 == nil {
		map1 = make(map[int]string)
		fmt.Println(map1 == nil)
	}
	//3.存键值对到map中
	map1[3] = "xixixi"
	map1[4] = ""
	//4.获取数据
	fmt.Println(map1[3])
	fmt.Println(map1[2]) //2为空，打印出string的空值，即”“
	/*如何才能知道自己对应键位存的是"0"值，还是空呢？
	在map的返回中除了value值，还有"ok-idiom",代表
	*/
	v4, ok := map1[4]
	if ok == true {
		fmt.Println("v4==", v4)
	}
	//5.修改数据
	map1[4] = "hahaha"
	fmt.Println(map1[4])
	//6.删除数据，函数delete
	delete(map1, 4)
	fmt.Println(map1)
	//7.获取长度
	fmt.Println(len(map1))
	//8.遍历map
	map4 := make(map[int]string)
	map4[1] = "今天是5.27号"
	map4[2] = "我的自行车被学校拖走了"
	map4[3] = "我的雪顶咖啡化完了"
	map4[4] = "风是透明的河流"
	map4[5] = "水是流动的山川"
	for k, v := range map4 {
		fmt.Println(k, v) //map为无序存储，所以每次打印顺序不一样
	}
	//如何有序打印
	fmt.Println("-----------------")
	for i := 1; i <= len(map4); i++ {
		fmt.Println(i, "-->", map4[i])
	} //如果k没有顺序就无法这样打印
	/*
		1.获取所有的key,->切片/数组存储
		2.进行排序
		3.遍历key,->map[key]
	*/
	key := make([]int, 0, len(map4))
	fmt.Println(key)
	for k := range map4 {
		key = append(key, k)
	}
	fmt.Println(key) //每次顺序都不一样
	//排序
	sort.Ints(key)
	fmt.Println(key)
	for _, keys := range key {
		fmt.Println(keys, map4[keys])
	}
	//9.map和slice结合使用，创建map->存入信息到map里->把map存入到slice里->遍历输出
	map5 := make(map[string]string)
	map5["name"] = "艾美莉卡"
	map5["age"] = "244"

	map6 := make(map[string]string)
	map6["name"] = "西安邮电大学"
	map6["age"] = "74"

	map7 := make(map[string]string)
	map7["name"] = "荷叶饭"
	map7["age"] = "18"
	str := make([]map[string]string, 0, 3)
	str = append(str, map5)
	str = append(str, map6)
	str = append(str, map7)
	for i, v := range str {
		fmt.Printf("第%d个人的姓名为%s\n", i+1, v["name"])
		fmt.Printf("年龄为%s\n", v["age"])
		fmt.Println()
	}
	//10.map的类型
	fmt.Printf("%T\n", map7) //类型：map[string][string]
	//也可以在map里套map
	map8 := make(map[string]map[string]string) //k为string,v为map[string][string]
	m1 := make(map[string]string)
	m1["name"] = "Go"
	m1["age"] = "15"
	map8["message"] = m1
	fmt.Println(map8)
}
