package func_test

import (
	"fmt"
	//"interface_test"
	"math/rand"
	"testing"
	"time"

	cm "github.com/easierway/concurrent_map"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

// 自定义方法别名（签名）
type IntConv func(op int) int

// 函数式编程
func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Print("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

// 可变参数
func Sum(ops ...int) int {
	ret := 0
	fmt.Printf("type of ops = %T\n", ops) // []int
	for _, op := range ops {
		ret += op
	}
	return ret
}
func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources.")
}

// defer 函数
func TestDefer(t *testing.T) {
	defer Clear() // 在函数执行完前，执行这个defer；常做一些清理操作。
	fmt.Println("start")
	panic("err") // panic 不会执行后面的语句，但依然要执行defer
	fmt.Println("end")
}
//
//func TestPackage(t *testing.T) {
//	interface_test.GoProgrammer()
//	var p interface_test.Programmer = new(interface_test.GoProgrammer)
//	t.Log(p.WriteHelloWorld())
//}

func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}


/**
	函数是一等公民
	1、可以有多个返回值
	2、所有参数都是值传递：slice、map、channel会有传引用的错觉
	3、函数可以作为变量的值
	4、函数可以作为参数和返回值
**/
