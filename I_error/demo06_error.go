package main

import "fmt"

func main()  {
	length,width:=6.7,-9.1
	area,err:=rectArea(length,width)
	if err!=nil {
		fmt.Println(err)
		if ins,ok:=err.(*areaError);ok {
			if ins.lengthNegative() {
				fmt.Printf("长度：%.2f,小于0\n",ins.lenght)
			}
			if ins.widthNegative() {
				fmt.Printf("宽度：%.2f,小于0\n",ins.width)
			}
		}
		return
	}
	fmt.Println(area)
}
type areaError struct {
	msg string
	lenght float64//发生错误时矩形的长度
	width float64//发生错误时矩形的宽度
}

func (e *areaError)Error()  string	{
	return e.msg
}
func (e *areaError)lengthNegative() bool {
	return e.lenght<0
}
func (e *areaError)widthNegative() bool {
	return e.width<0
}
func rectArea(length,width float64)(float64,error)  {
	msg:=""
	if length<0 {
		msg="长度小于0"
	}
	if width<0{
		if msg=="" {
			msg="宽度小于0"
		}else {
			msg+="，宽度也小于0"
		}
	}
	if msg!="" {
		return 0,&areaError{msg,length,width}
	}
	return length*width,nil
}


