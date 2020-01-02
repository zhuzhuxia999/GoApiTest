package main

import "fmt"

type Books struct {
	title string
	author string
	suject string
	book_id int
}

func main()  {
	/*//创建一个新的结构体
	fmt.Println(Books{"go 语言" , "王一杰","计算机" , 123})

	///忽略的字段为 O或者 空
	fmt.Println(Books{title:"go 语言"})*/

	var Book1  Books //声明book1 为 Books 类型
	var Book2 Books //声明 book2 为Books 类型

	//book1 描述
	Book1.title = "Go 语言"
	Book1.author = "王一杰"
	Book1.suject = "计算机"
	Book1.book_id = 1234

	//book2描述
	Book2.title = "python 教程"
	Book2.author = "王一杰的小可爱"
	Book2.suject = "设计师"
	Book2.book_id = 521

	//打印 结构体中的属性的信息
	fmt.Println(Book1.title + Book1.author  + Book1.suject)
	fmt.Println(Book2.title + Book2.author  + Book2.suject)
}


