package main

import "fmt"

/*//局部变脸示例
func main()  {
	//声明局部变量
	var a , b ,c int

	//初始化参数值
	a = 10
	b = 20
	c = a + b
	fmt.Println("相加的结果是： ", c )
}*/

var g  int

func main()  {
	//声明局部变量
	var a , b int

	//初始化参数
	a =  10
	b =  20
	g = a + b

	fmt.Println("相加的结果是：", g )
}