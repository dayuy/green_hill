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

func main() {
	defer fmt.Println("main end by defer")
	defer fmt.Println("main defer 2")

	fmt.Println("main:: hello go 1")
	fmt.Println("main:: hello go 2")

	// return 和 defer 谁先执行
	returnAndDefer()
}
