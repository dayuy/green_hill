package groutine_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGroutine(t *testing.T){
	for i:=0;i<10;i++{
		go func(i int) {   // 在不同的协程里做
			fmt.Println(i) // 9 3 4 5 678012 因为协程被调用时不是按照编程顺序的
		}(i)
	}
	time.Sleep(time.Millisecond * 50)
}


/**
	共享内存并发机制
**/
func TestCounter(t *testing.T) {
	counter := 0
	for i:=0;i<5000;i++{
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter=%d", counter) // counter = 4791; 因为counter在不同的协程里做自增，并发条件 不是线程安全的程序
}

// 改造成为 线程安全的程序:互斥锁
func TestCounterThreadSafe(t *testing.T)  {
	var mut sync.Mutex // 互斥锁
	counter := 0
	for i:=0;i<5000;i++{
		go func() {
			defer func() {
				mut.Unlock() // 解锁
			}()
			mut.Lock() // 每个协程上锁
			counter++
		}()
	}
	time.Sleep(1 * time.Second) // 4984 如果不加sleep，外部的协程会先于循环内的程序执行完，结果是4984 （程序是按顺序执行的，当使用go func(){}(),则是开启里另一个协程）
	t.Logf("counter=%d", counter) // 5000
}

// 改造外部协程先于for循环内的协程执行完
func TestCounterThreadGroup(t *testing.T)  {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i:=0;i<5000;i++{
		wg.Add(1) // 增加等待的count
		go func() {
			defer func() {mut.Unlock()}()
			mut.Lock()
			counter++
			wg.Done() // 每一个count完成
		}()
	}
	wg.Wait() // 等待所有count完成
	t.Logf("counter=%d", counter)
}


/**
	CSP并发机制 communication sequential process
	go 中的channel是有容量限制，并且独立处理Groutine的。
	使程序具有异步的能力
**/
func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()

	//Done
	//working on something else
	//Task is done.
	// 串行的输出，按照执行顺序
}

// 改造为异步，1、在调用service时 启动另外一个协程，而不是阻塞当前程序运行
// 		     2、当返回结果时，放到channel里，需要用的话 从channel里读取
func AsyncService() chan string {
	//retCh := make(chan string) // 声明一个channel
	retCh := make(chan string, 1) // 声明一个buffer channel：是内部不阻塞
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh) // 读取channel

	//working on something else
	//returned result.
	//Task is done.
	//Done
	//service exited. 需要等到有人使用channel，才能往下执行 fmt.Println("service exited").相当于channel内部阻塞了。
	// 于是便可以设置channel 为buffer channel： make(chan string)
	//  working on something else
	//  returned result.
	//	service exited.
	//	Task is done.
	//	Done
}





