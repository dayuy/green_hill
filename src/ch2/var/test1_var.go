package main

import "fmt"

/**
	四种变量声明方式
**/

var gA int = 100
var gB = 'h'

func main()  {
	var a int
	fmt.Println("a = ", a) // 0

	var b int = 100
	fmt.Println("b=", b)

	var c = 100
	fmt.Println("c=", c)
	fmt.Printf("c = %d, type of c = %T\n", c, c)

	var d = "abcd"
	fmt.Printf("d = %s, type of bb = %T\n", d, d)

	e :=100
	fmt.Printf("e= %d, type of e = %T\n", e, e)

	fmt.Printf("ga = %d, type of ga = %T\n", gA, gA)

	fmt.Printf("gb = %d, type of ga = %T\n", gB, gB)

	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, "yy = ", yy)

	var kk, ll = 100, "Aceld"
	fmt.Println("kk = ", kk, "ll = ", ll)

	var (
		vv int = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, " jj = ", jj)

	var hh int = 1
	changeValue(hh)
	fmt.Println("hh = ", hh)

	var l int = 2
	changeValue2(&l) // 传的是内存地址
	fmt.Println("l = ", l)

	swap(&xx, &yy)
	fmt.Println(xx, yy)

	// 二级指针
	var u *int
	u = &l
	fmt.Println(u)
	fmt.Println(&l)

	var uu **int
	uu = &u
	fmt.Println(&u)
	fmt.Println(uu)
}

func changeValue(p int) {
	p = 10
}

func changeValue2(p *int) { // 指针
	*p = 10
}

func swap(pa *int, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}

