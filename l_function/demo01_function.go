package main

import "fmt"

func main1() {
	/*
		定义函数的语法：
			func  funcName(parametername type1,parametername type2)(output1 type1,output2,type2)
		注意事项：
			1.函数必须先定义再调用
			2.函数名不能冲突
			3.main()，是一个特殊的函数，作为程序的入口，由系统自动调用
	*/
	getSum()

}
func getSum()  {
	sum:=0
	for i:=1;i<=100 ;i++  {
		sum+=i
	}
	fmt.Println(sum)
}

