package main

/**
	面向对象 OOP
**/

import "fmt"

// 声明一种数据类型, 是 int 的一个别名
type myint int

// 定义一个结构体
type Book struct {
	title string
	auth string
}

func main1() {
	var a myint = 10
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhang3"

	fmt.Printf("%v\n", book1)

	changeBook(book1)
	fmt.Printf("change 1 %v\n", book1)

	changeBook2(&book1)
	fmt.Printf("change 2 %v\n", book1)
}

func changeBook(book Book) { // 传递的是一个book的副本
	book.auth = "666"
}

func changeBook2(book *Book)  { // 指针传递
	book.auth = "777"
}
