package main

import "fmt"

func deferFun() int {
	fmt.Println("defer func called...")
	return 0
}

func returnFun() int {
	fmt.Println("return func called...")
	return 0
}

func returnAndDefer() int {
	defer deferFun()
	return returnFun()
}

/*
* defer:
* 1. 当执行到defer时，暂时不执行，而是会将defer后面的语句压入到**独立的栈**（defer栈）
* 2. 当函数执行完毕后，再从defer栈，按照**先入后出**的方式出栈（倒叙）：先执行最后一个defer
* 3。当压入栈时，会将相关的值拷贝，
 */
func main() {
	n1 := 0
	defer fmt.Println("main end by defer", n1) // 0
	defer fmt.Println("main defer 2")

	n1++                                 // 1
	fmt.Println("main:: hello go 1", n1) // 1
	fmt.Println("main:: hello go 2")

	// return 和 defer 谁先执行
	returnAndDefer()
}
