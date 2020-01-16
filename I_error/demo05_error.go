package main

import (
	"fmt"
	"math"
)

func main() {
	/*
	自定义错误
	 */
	radius:=-3.0
	area,err:=circleArea(radius)
	if err!=nil{
		fmt.Println(err)
		if ins,ok:=err.(*areaError);ok{
			fmt.Printf("半径是：%.2f\n",ins.radius)
		}
		return
	}
	fmt.Println(area)

}
//定义一个结构体，表示错误的类型
type areaError struct {
	msg string
	radius float64
}

func (e *areaError)Error() string {
	return fmt.Sprintf("error:半径：%.2f,%s",e.radius,e.msg)
}
func circleArea(radius float64)(float64,error)  {
	if radius<0 {
		return 0,&areaError{"半径是非法的",radius}
	}
	return math.Pi*radius*radius,nil
}
