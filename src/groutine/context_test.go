package groutine_test

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
	"unsafe"
)

/**
	Context
	1. 根Context：通过context.Background()创建
	2. 子Context：context.WithCancel(parentContext)创建
		ctx,cancel:=context.WithCancel(context.Background())
	3. 当前Context被取消时，基于他的子context都会被取消
	4. 接收取消通知 <-ctx.Done()
**/

func isCancelledContext(ctx context.Context) bool {
	select {
	case <-ctx.Done(): // 接收取消通知
		return true
	default:
		return false
	}
}

func TestCancelContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for  {
				if isCancelledContext(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}


/**
	并发任务
	1。 在多线程的情况下保证代码只执行一次 懒汉式
		单例模式 sync.Once
**/
type Singleton struct {

}
var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			obj:=GetSingletonObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}

/**
	并发模式
	first response
**/
func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}
func FirstResponse() string {
	numOfRunner:=10
	//ch:=make(chan string)
	ch:=make(chan string, numOfRunner) // buffer Channel,不需要等待消息的接收，不会阻塞
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch<-ret
		}(i)
	}
	return <-ch
}

func AllResponse() string {
	numOfRunner:=10
	ch:=make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret:=runTask(i)
			ch <- ret
		}(i)
	}
	finalRet:=""
	for j:=0;j<numOfRunner;j++{
		finalRet += <-ch + "\n"
	}
	return finalRet
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before", runtime.NumGoroutine()) // 所有协程数 2
	//t.Log(FirstResponse())
	t.Log(AllResponse())
	time.Sleep(time.Second * 1) // 等所有任务完成，但只取了一次，还有11条没有释放
	t.Log("After:", runtime.NumGoroutine()) // 11
}