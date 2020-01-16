package main

import "fmt"

func main() {
	/*
	函数的参数；
		形式参数：也叫形参。函数定义的时候，用于接收外部传入的数据的变量。
			函数中，某些变量的数值无法确定，需要由外部传入数据
		实际参数：也叫实参，函数调用的时候，给形参赋值的实际数据

	 */
	getSum2(1000)

	//求两个数的和
	getAdd(10,20)

	getAdd2(1,2)
}
func getSum2( n int)  {
	sum:=0
	for i:=1;i<=n ;i++  {
		sum+=i
	}
	fmt.Printf("1到%d的和是:%d\n",n,sum)
}
func getAdd(a int,b int)  {
	sum:=a+b
	fmt.Printf("%d + %d= %d\n",a,b,sum)
}
func getAdd2(a,b int)  {//参数类型一致，可以简写在一起
	fmt.Printf("%d,%d\n",a,b)
}
func fun1(a,b float64 ,c string)  {

}
