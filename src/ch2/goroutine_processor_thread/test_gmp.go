package main

import (
	"fmt"
	"time"
)

// 子goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

/**
	结论：创建一个子routine，此空间依赖于主routine空间：
		如果主routine退出，则子routine将被退出
**/
// 主goroutine
func main1() {
	// 正常流程会按顺序执行
	//newTask()

	// 创建一个go线程 去执行 newTask() 流程
	go newTask()

	i := 0
	for i < 8 {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("main-线程===")
}
