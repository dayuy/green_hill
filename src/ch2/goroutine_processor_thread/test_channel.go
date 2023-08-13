package main

import (
	"fmt"
	"time"
)

/**
	channel：
	goroutine之间互相通信，通过channel
	注意现象：主goroutine没有立即执行完退出，使子goroutine得以执行
	注意结果：主goroutine中有管道，会等待管道中的数据（阻塞在这里）
**/

func channel1() {
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine结果")
		fmt.Printf("子goroutine，正在运行")

		c <- 666 // 将666传输给c管道
	}()

	num := <-c // 从c管道中接收数据，并赋值给num
	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束...")
}

/**
	缓冲channel  make(chan int, 3)
	注意：当channel已满，再往里存数据，会阻塞，等待被取出
		 当channel为空，取数据会报错 fatal error: all goroutines are asleep - deadlock!
**/
func channel2() {
	c := make(chan int, 3)

	fmt.Println("len(c) = ", len(c), ", cap(c)", cap(c))

	go func() {
		defer fmt.Println("子goroutine结束")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子go程正在输送，发送的元素=", i, " len(c) = ", len(c), " cap(c) = ", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-c // 从c管道里取数据，并赋值给num
		fmt.Println("num = ", num)
	}

	fmt.Println("main 结束")
}

func main() {
	//channel1()
	//channel2()
	//channel3()
	channelSelect()
}

/**
close() 关闭channel管道
注意：向关闭的channel，传送数据，会引发panic，导致接收立即返回0值  panic：send on closed channel
	 关闭channel后，可以继续从channel中接收数据
     对于nil channel，无论收发都会被阻塞
**/
func channel3() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		close(c) // close可以关闭一个channel
	}()

	for {
		// ok如果为true，表示channel没有关闭，如果为false，表示channel已经关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	// 也可以使用range迭代
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Main Finished...")
}

/**
	select 监控多个channel的状态
	：单流程下一个go只能监控一个channel的状态，select可以完成监控多个channel的状态
**/
func fibonacii(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x = y
			y = x + y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}

func channelSelect() {
	c := make(chan int)
	quit := make(chan int)

	// sub goroutine
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	// main goroutine
	fibonacii(c, quit)
}
