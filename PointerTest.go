package main

import "fmt"

/*func main()  {
	var a int  = 10
	fmt.Println("变量的地址：%x\n", &a)
}

*/

/*func main()  {
	var a int  = 20 //声明实际变量
	var ip  *int //声明指针变量

	ip = &a  //指针变量的存储地址
	fmt.Println("a 变量的地址是：%x\n", ip)

	//使用指针访问值
	fmt.Println("*ip变量的值：%d\n",  *ip)
}*/

//空指针 nil
func main()  {
	var ptr  *int

	fmt.Println("ptr 的值为： %x\n",ptr)
}