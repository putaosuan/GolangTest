package main

import "fmt"

func main() {
	/*
	递归函数

	 */
	//求1-n的和
	s1:=getSum(20)
	fmt.Println(s1)
	//fibnacci数列
	/*
	1 2 3 4 5 6 7  8  9
	1 1 2 3 5 8 13 21 34
	 */
	s2:=getFibnacci(9)
	fmt.Println(s2)

}
func getFibnacci(n int) int  {
	if n<=2 {
		return 1
	}
	return getFibnacci(n-1)+getFibnacci(n-2)
}
func getSum(n int)(int)  {
	if n==1 {
		return 1
	}else {
		return getSum(n-1)+n
	}
}

