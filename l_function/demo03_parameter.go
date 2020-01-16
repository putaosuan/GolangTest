package main

import "fmt"

func main()  {
	/*
	可变参数:一个函数的参数类型确定，但是个数不确定，就可以使用可变参数
	语法：
		参数名...参数的类型
		对于函数，可变参数相当于一个切片
		调用函数的时候，就可以传入0-多个参数
			Println()，Printf()，Print(),append()
	注意事项：
		如果一个函数的参数是可变参数，同时还有其他参数，可变参数要放在参数列表的最后
		一个函数的参数列表中，最多只能有一个可变参数
	 */
	//求和
	getSum(1,2,3,4,5,6,7,8,9,10)
	//切片
	s1:=[]int{1,2,3,4,5}
	getSum(s1...)
}
func getSum(nums...int)  {
	fmt.Printf("%T\n",nums)
	sum:=0
	for i := 0; i< len(nums);i++  {
		sum+=nums[i]
	}
	fmt.Println(sum)
}
//func fun1(s1,s2 string,nums...floats)  {

//}
