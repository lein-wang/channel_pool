package lang

import (
	"fmt"
)

func Test() {
	var i = 1
	fmt.Println(i << 3)
	{
		i = 2
		fmt.Println(i >> 3)
	}

}

func check1()(x int,err error){
	if err != nil {
		return x,err
	}
	return x,nil
}

func check2(x int)(y int,err error){
	if err != nil {
		return x,err
	}
	return x,nil
}
func Shadow()(err error)  {
	x,y := 1,2
	x,err = check1()
	fmt.Println(x)
	if err != nil{
		return
	}
	if y,err = check2(x);err != nil{
		return
	}else {
		fmt.Println(y)
	}
	return

}