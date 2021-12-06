package main

import (
	"fmt"
	"runtime"
	"time"
)

func main2() {
	// 开启一个routine： 用go创建承载一个形参为空，返回值为空的函数
	go func() {
		defer fmt.Println("A.defer")
		// return
		func() {
			defer fmt.Println("B.defer")
			// return B.defer -- A -- A.defer 只是退出当前routine
			runtime.Goexit() // B.defer -- A.defer 退出goroutine 到主routine
			fmt.Printf("B")
		}()

		fmt.Println("A")
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}
