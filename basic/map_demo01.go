package main

import "fmt"

func main38()  {
	/*
	map：映射，是一种专门用来存储键值对的集合，属于引用类型
		存储特点：
			1.存储是无序的键值对
			2.键不能重复，并且和value一一对应
				map中的key值不能重复，如果重复，新的value会覆盖原来的，程序不会报错
		语法结构：
			1.创建map
				var map1 map[key类型]value类型
					nil map,无法直接使用
			2. var map2=make(map[key类型]value类型)
			3. var map3=map[key类型] value类型{key:value,key:value...}
		2.添加/修改
			map[key]=value
			如果key存在就是修改数据，如果key不存在就是添加数据
		3.获取
			map[key]-->value

			value,ok:map[key]
			根据key获取对应的value，如果key存在，value就是对应的数值，ok为true
			如果key不存在，value就是值类型的默认值，ok为false
		4.删除数据
			delete(map,key)
			如果key存在，就可以直接删除
			如果key不存在，删除失败
		5.长度：
			len()
	 */
	//1.创建map 三种方法
	var map1 map[int]string
	var map2=make(map[int] string)
	var map3=map[string] int{"GO":12,"Python":22,"Java":32}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	fmt.Println(map1==nil)
	fmt.Println(map2==nil)
	fmt.Println(map3==nil)
	//map1[1]="hello"//panic: assignment to entry in nil map
	if map1==nil {
		map1=make(map[int]string)
		fmt.Println(map1==nil)
	}
	fmt.Println("---------------")
	//2.存储键值对到map中
	map1[1]="hello"
	map1[2]="world"
	map1[3]="张三"
	map1[4]="李四"
	map1[5]=""
	fmt.Println(map1)
	//3.获取数据，根据key值获取对应的value值
	//如果存在，获取数值，如果不存在，获取的是value值类型的零值
	fmt.Println(map1[3])
	fmt.Println(map1[30])

	v1,ok:=map1[5]
	if ok{
		fmt.Println("对应的 数值是",v1)
	}else {
		fmt.Println("操作的key不存在，获取到的是零值")
	}
	//5.修改数据
	fmt.Println(map1)
	map1[1]="如花"
	fmt.Println(map1)
	//6.删除数据
	delete(map1,3)
	fmt.Println(map1)
	delete(map1,30)
	fmt.Println(map1)
	//长度
	fmt.Println(len(map1))

}
