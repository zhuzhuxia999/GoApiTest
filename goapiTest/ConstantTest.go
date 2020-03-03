package main

/**
常量
 */
/*func main()  {
	const length  = 10
	const width = 5
	var area  int
	const  a , b , c  = 1 , false , "str" //多重赋值

	area = length * width
	fmt.Println("面积为 ：", area)
	println()
	println(a , b , c)
}

/**
变量用作枚举

const (
	unKnow = 1
	female = 2
	male = 3
)*/

/*import "unsafe"
声明常量时使用 内置函数
const(
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main()  {
	println(a , b , c)
}

*/


/*
iota 实例
import "fmt"

func main()  {
	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota //恢复计数
		i
	)
	fmt.Println(a , b , c , d ,  e , f , g , h , i )

}

*/

import "fmt"

const (
	i = 1<<iota
	j = 3<<iota
	k
	l
)

func consantTest1()  {
	fmt.Println("I=", i )
	fmt.Println("j=", j )
	fmt.Println("k=", k )
	fmt.Println("l=", l )
}
