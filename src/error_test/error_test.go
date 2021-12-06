package error_test

import (
	"errors"
	"fmt"
	"testing"
)

var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThanHundredError = errors.New("n should be not larger than 100")

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
		if err == LessThanTwoError {
			fmt.Println("it is less")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}

/**
  recover 覆盖错误。小心使用，应该let it crash。防止僵尸服务进程
**/
func TestPanicVxExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil { // recover 阻挡了panic。
			fmt.Println("recovered from", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("something wrong!"))
}
