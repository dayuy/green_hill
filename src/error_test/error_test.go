package error_test

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThanHundredError = errors.New("n should be not larger than 100")

func GetFibonacciNoError(n int) []int {
	fibList := []int{1,1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2] + fibList[i-1])
	}
	return fibList
}

// 处理边缘情况 使用errors.New()
func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}

	if n > 100 {
		return nil, LargerThanHundredError
	}
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(-10); err != nil {
		if err == LessThanTwoError { // 直接判断是哪种错误
			fmt.Println("it is less")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}

func TestGetFibonacciNoError(t *testing.T) {
	t.Log(GetFibonacciNoError(10))
}

/**
	panic 会执行defer指定的函数
		  会输出调用栈的信息
		  不会继续执行后面的语句
	os.Exit 退出时不会调用 defer 指定的函数
			退出时不输出当前调用栈信息
			退出时没有报错信息，Pass
**/
func TestPanicVxExit(t *testing.T) {
	defer func() {
		fmt.Println("Finally defer!")
	}()

	fmt.Println("Start")
	//panic(errors.New("something wrong"))
	os.Exit(-1)
	fmt.Println("End")
}

/**
  recover 覆盖错误。小心使用，应该let it crash。防止僵尸服务进程
**/
func TestPanicVxRecover(t *testing.T) {
	defer func() {
		/*
		2、在有 recover() 的情况下，程序执行通过。掩盖了错误
			PASS: TestPanicVxExit
			PASS
		*/
		if err := recover(); err != nil { // recover 阻挡了panic。
			fmt.Println("recovered from", err)
		}
	}()
	fmt.Println("Start")
	/**
	1、在没有 recover() 的情况下，程序执行失败而中断
	   FAIL: TestPanicVxExit
	   panic: something wrong!
	**/
	panic(errors.New("something wrong!"))
}



/**
	Go 的错误机制
	1、没有异常机制
	2、error类型实现了error接口
	   ：type error interface{ Error() string }
	3、可以通过errors.New 来快速创建错误实例
		：errors.New("n must be in the range [0, 10]")
**/
