package main

import "fmt"

func main() {
	p1:=Person{"王二狗",30}
	fmt.Println(p1.name,p1.age)
	p1.eat()
	//创建Student类型
	s1:=Student{Person{"Ruby",8},"潭州"}
	fmt.Println(s1.name)//s1.Person.name
	fmt.Println(s1.age)//子类对象，可以直接访问父类的字段，（其实就是提升字段）
	fmt.Println(s1.school)//子类对象，访问自己新增的对象属性
	s1.eat()//子类对象，访问父类的方法
	s1.study()//子类对象，访问自己新增的方法
	s1.eat()//如果存在方法的重写，子类对象访问重写的方法
}
//定义一个“父类”
type Person struct {
	name string
	age int
}
//定义一个“子类”
type Student struct {
	Person
	school string
}
//
func (p Person)eat()  {
	fmt.Println("这是父类中的方法，吃窝窝头。。。")
}
func (s Student)study()  {
	fmt.Println("子类新增的方法，在学习。。。")
}
func (s Student)eat()  {
	fmt.Println("子类重写的方法，在吃炸鸡，喝啤酒")
}

