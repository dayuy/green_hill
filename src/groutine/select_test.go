package groutine_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
	多渠道选择，select，channel都没有返回的话会阻塞
**/
func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 30):
		t.Error("time out")
	}
}


/**
	关闭channel（因为ch<-i,必须与 <-ch相同。否则会报错）
	1。 向关闭的channel发送数据，会导致panic
	2、v,ok <- ch;ok为true时表示正常接受，false表示通道关闭
	3、所有的channel接收者都会在channel关闭时，立刻从阻塞等待中返回且ok为false。
		使用场景：这个广播机制常被利用，进行向多个订阅者同时发送信号（退出信号）
**/
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i:=0;i<11;i++{
			ch <- i
		}
		close(ch) // 增加close，防止阻塞
		wg.Done()
	}()
}
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		//for i:=0;i<11;i++{
		//	data:=<-ch
		//	fmt.Println(data)
		//}
		for {
			if data, ok := <-ch; ok{
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}
func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch:=make(chan int)
	wg.Add(1)
	dataProducer(ch,&wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}


/**
	例2：利用cancel()的广播机制，关闭所有协程
**/
func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{} // 发了个任意的data 到chan上，告诉其取消任务
}

func cancel_2(cancelChan chan struct{})  {
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i:= 0;i<5;i++{
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh){
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}
	//cancel_1(cancelChan) // 4 Cancelled 有5个人协程，只有其中一个被cancel了，因为只有一个channel信号
	cancel_2(cancelChan) // 4 cancelled 2 cancelled 3 cancelled 0 cancelled 1 cancelled
	time.Sleep(time.Second * 1)
}