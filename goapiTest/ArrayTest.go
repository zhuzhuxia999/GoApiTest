package main

import "fmt"

func arrayTest1()  {
	var n [10] int//声明一个长度为10 的数组
	var i , j int

	//为数组初始化元素 的值
	for i = 0; i <10 ; i++  {
		n[i] = i + 100
	}

	//便利输出数组中的值
	for j = 0; j < 10 ; j++  {
		fmt.Println("Element[%d]", j , n[j])
	}
}