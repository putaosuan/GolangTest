package main

import "fmt"

func main41() {
	/*
	一、数据类型
		基本数据类型：int,float,bool,string
		复合数据类型：array,slice,map,function,pointer,struct...
			array:[size]数据类型
			slice:[]数据类型
			map
	二、存储特点
		值类型：int,float,string,bool,array,struct
		引用类型：slice,map,chan
			其实make()创建的都是引用类型
	 */
	map1:=make(map[string]map[string]string)
	m1:=make(map[string]string)
	m1["name"]="王二狗"
	m1["age"]="23"
	m1["sex"]="男"
	m2:=make(map[string]string)
	m2["name"]="李小花"
	m2["age"]="20"
	m2["sex"]="女"
	map1["经理"]=m1
	map1["总经理"]=m2
	fmt.Println(map1)
	fmt.Println("-----------")
	map2:=make(map[string]string)
	map2["王二狗"]="住在一楼"
	map2["李小花"]="住在二楼"
	map2["ruby"]="住在三楼"

	map3:=make(map[string]string)
	map3=map2
	map3["ruby"]="不住了"
	fmt.Println(map2)
	fmt.Println(map3)
}
