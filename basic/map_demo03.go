package main

import (
	"fmt"
)

func main40() {
	/*
	map和slice的结合使用：

	 */
	map1:=make(map[string]string)
	map1["name"]="王二狗"
	map1["age"]="30"
	map1["sex"]="男"
	map1["address"]="北京市"
	fmt.Println(map1)

	map2:=make(map[string]string)
	map2["name"]="李小华"
	map2["age"]="20"
	map2["sex"]="女"
	map2["address"]="上海市"
	fmt.Println(map2)

	map3:=map[string]string{"name":"ruby","age":"25","sex":"女","address":"天津市"}
	fmt.Println(map3)

	//将map存入slice中
	s1:=make([]map[string]string,0,3)
	s1 = append(s1, map1)
	s1 = append(s1, map2)
	s1 = append(s1, map3)
	//遍历切片
	for i,v:=range s1 {
		fmt.Printf("第%d个人的信息是:\n",i+1)
		fmt.Printf("\t姓名：%s,性别：%s,年龄：%s,地址：%s\n",v["name"],v["sex"],v["age"],v["address"])
	}
}
