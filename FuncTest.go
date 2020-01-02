package main

import "fmt"

func main()  {
	var result = max(2, 3)
	fmt.Println("max : ",result)
	a,b := swap("王一杰","王一杰的小可爱")
	fmt.Println("swap() :",a,b)
}

//函数示例
func max(num1,num2 int)  int{
	//声明局部变量
	var result int;

	if(num1>num2){
		result = num1
	}else {
		result = num2
	}
	return result
}

//函数返回多个值
func swap(x,y string)(string,string)  {
	return x,y
}