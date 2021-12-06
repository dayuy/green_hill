package main

import "fmt"

// const 定义枚举类型
const (
	BEIJING = 0
	SHANGHAI = 1
	SHENZHEN = 2
)

const (
	GUANGZHOU = iota  // 可以在 const () 里加一个关键字 iota，每行的iota都会累加1，第一行的iota默认值为0
	CHONGXING
	SHIJIAZHUANG
)

const (
	a, b = iota+1, iota+2    // 1 2
	c, d                     // 2 3
	e, f                     // 3 4
	g, h = iota *2, iota * 3  // 6 9
	i, k                      // 8 12
)

func main() {
	const length int = 10
	fmt.Println("length = ", length)

	//length = 11; 常量是不允许被修改的

	fmt.Println("length = ", BEIJING, CHONGXING)

	fmt.Println(a,b,c,d,e,f,g,h,i,k)
}

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100
	return c
}

func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 666, 777
}

func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("----foo3----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// 给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000

	return
}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("---foo4---")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	r1 = 1000
	r2 = 2000
	return
}
