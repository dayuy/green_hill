package main

import (
	"fmt"
	"io"
	"os"
)

// interface{} 是万能通用型数据类型，是一个指针的形式
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	// interface{} 怎么区分类型？类型断言
	value, ok := arg.(float32)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Printf("arg is string type, value = %v, type = %T\n", value, value)
	}
}

type Bookes struct {
	auth string
}

func main() {
	book := Bookes{"golang"}

	myFunc(book)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14)


	// ? 为什么可以使用类型断言
	var a string
	a = "aceld" // pair<staticType:string, value: "aceld">

	var anyType interface{}
	anyType = a // pair<type:string, value: "aceld">

	str, _ := anyType.(string)
	fmt.Println(str)

	osTest()
}

func osTest() {
	tty, err := os.OpenFile("./tty", os.O_RDWR, 0)
	// tty: pair <type: *os.File, value: "/dev/tty">
	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	var r io.Reader // r:pair <type:, value:>
	r = tty // r:pair <type: *os.File, value: "/dev/tty">

	var w io.Writer  // w:pair <type:, value:>
	w = r.(io.Writer) // w:pair <type: *os.File, value: "/dev/tty")

	w.Write([]byte("HELLO THIS is A TEST!!!\n"))
}

type Reader interface {
	ReadBook()
}
type Writer interface {
	WriteBook()
}

type Book1 struct {
}

func (this *Book1) ReadBook() {
	fmt.Println("Read a Book")
}

func (this *Book1) WriteBook() {
	fmt.Println("Write a Book")
}

func test_2() {
	b := &Book1{} // b: pair<type: Book1, value: book{}>
	b.ReadBook()
	b.WriteBook()

	var r Reader // r: pair<type: , value:>
	r = b // r: pair<type: Book, value: book{}>
	r.ReadBook()
	//r.WriteBook Reader没有这个方法

	var w Writer
	// 类型断言
	w = r.(Writer) // w: pair<type: Book, value: book{}>
	w.WriteBook()
}