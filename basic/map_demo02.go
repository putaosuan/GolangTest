package main

import (
	"fmt"
	"sort"
)

func main39() {
	/*
	map的遍历
		使用for range
	 */
	map1:=make(map[int]string)
	map1[1]="小米"
	map1[2]="华为"
	map1[3]="OPPO"
	map1[4]="Vivo"
	//1.遍历map
	for k,v:=range map1 {
		fmt.Println("key；",k,"value:",v)
	}
	//
	fmt.Println("-----------")
	for i := 1; i<= len(map1);i++  {
		fmt.Println(i,"-->",map1[i])
	}
	/*
	1.获取所有的key，存在切片中
	2.进行排序
	3.遍历key
	 */
	keys:=make([]int,0, len(map1))
	for k,_:=range map1{
		keys=append(keys,k)
	}
	fmt.Println(keys)
	//冒泡排序，使用sort包下的排序方法
	sort.Ints(keys)
	for _,key:=range keys{
		fmt.Println(key,map1[key])
	}

	s1:=[]string{"Apple","Oracle","windows","abc","aaa","张三"}
	fmt.Println(s1)
	sort.Strings(s1)
	fmt.Println(s1)
}
