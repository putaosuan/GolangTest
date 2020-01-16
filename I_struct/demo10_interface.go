package main

import (
	"fmt"
	"math"
)

func main() {
	/*
	断言接口：
		方式一：
			1.instance:=接口对象.(实际类型)//不安全，会panic
			2.instance,ok:=接口对象.(实际类型)//安全
		方式二：
			switch instance:=接口对象.(type){
			case 实际类型:
				...
			case 实际类型:
				...
			...
			}
	 */
	var t1 Triangle=Triangle{3,4,5}
	fmt.Println(t1.peri())
	fmt.Println(t1.area())
	fmt.Println(t1.a,t1.b,t1.c)
	var c1 Circle=Circle{4}
	fmt.Println(c1.peri())
	fmt.Println(c1.area())
	fmt.Println(c1.radius)
	var c2 Shape
	c2=c1
	fmt.Println(c2.peri())
	fmt.Println(c2.area())

	testShape(t1)
	testShape(c1)
	testShape(c2)
	fmt.Println("-----------")
	getType(t1)
	getType(c1)
	getType(c2)
	//getType(100)

	var t2 *Triangle=&Triangle{3,4,5}
	fmt.Printf("%T,%p\n",t2,&t2)
	getType(t2)
	fmt.Println("--------")
	getType2(t1)
	getType2(c1)
	getType2(t2)
}
func getType2(s Shape)  {
	switch ins:=s.(type) {
	case Triangle:
		fmt.Println("是三角形啊",ins.a,ins.b,ins.c)
	case Circle:
		fmt.Println("是圆形啊",ins.radius)
	case *Triangle:
		fmt.Println("是三角形啊指针",ins.a,ins.b,ins.c)

	}
}
func getType(s Shape)  {
	if ins,ok:=s.(Triangle) ;ok{
		fmt.Println("三角形。。",ins.a,ins.b,ins.c)
	}else if ins,ok:=s.(Circle);ok {
		fmt.Println("圆形。。",ins.radius)
	}else if ins,ok:=s.(*Triangle);ok {
		fmt.Println("是三角形")
		fmt.Printf("%T,%p\n",ins,&ins)
		fmt.Printf("%T,%p\n",s,&s)
	}else{
		fmt.Println("我也不知道了")
	}
}
func testShape(s Shape)  {
	fmt.Printf("周长:%.2f,面积:%.2f\n",s.peri(),s.area())
}
//1.定义一个接口
type Shape interface {
	peri() float64
	area() float64
}
//2.定义实现类：三角形
type Triangle struct {
	a,b,c float64
}

func (t Triangle)peri() float64 {
	return t.a+t.b+t.c
}
func (t Triangle)area() float64 {
	p:=t.peri()/2
	s:=math.Sqrt(p*(p-t.a)*(p-t.b)*(p-t.c))
	return s
}
//实现类
type Circle struct {
	radius float64
}

func (c Circle)peri() float64{
	return c.radius*2*math.Pi
}
func (c Circle)area() float64 {
	return math.Pi*math.Pow(c.radius,2)
}
