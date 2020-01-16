package main

import "fmt"

func main34() {
	/*
	切片slice:
		1.每一个切片引用了一个底层数组
		2.切片本身不存储任何数据，都是这个底层数据存储，所以修改切片也就是修改这个数组中的数据
		3.当向切片中添加数据时，如果没有超过容量，直接添加，如果超过容量，自动扩容，成倍增长
		4.切片一旦扩容，就是重新指向一个新的底层数组
	 */
	s1:=[] int{1,2,3}
	fmt.Printf("长度：%d,容量：%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n",s1)//0xc000012360
	s1 = append(s1, 4,5)
	fmt.Printf("长度：%d,容量：%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n",s1)//0xc00000c330
	s1 = append(s1, 6,7,8)
	fmt.Printf("长度：%d,容量：%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n",s1)//0xc000054060

}
