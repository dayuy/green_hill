package unit_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type A [0][256]int

type S struct {
	x A
	y [1 << 30]A
	z [1 << 30]struct{}
}

type T [1 << 30]S

/*
*  unsafe.Sizeof() 返回类型v本身数据所占用的字节数
*  len() the number of elements
*  1。 一个0尺寸（所占有的字节数为0？）的数组，或元素长度为0的数组，长度也为0
 */
func TestZero(t *testing.T) {
	var a A
	var s S
	var t1 T
	println(unsafe.Sizeof(a), len(a))
	println(unsafe.Sizeof(s))
	println(unsafe.Sizeof(t1))

	var x [1<<63 - 1]struct{}
	var y [2000000000 + 1]byte // PASS
	//var z = make([]byte, 1<<49) // panic: runtime error: makeslice: len out of range
	println(unsafe.Sizeof(x), 1<<2) // 0
	fmt.Printf("%d - %b\n", 4, 4)
	fmt.Printf("%d - %b\n", 5, 5)
	fmt.Printf("%d - %b\n", 6, 6)
	println(unsafe.Sizeof(y)) // 2000000001
	//println(unsafe.Sizeof(z))
}

var g *[0]int
var a, b [0]int

func f() *[0]int {
	return new([0]int)
}

/**
* 2。所有在堆上分配的局部零尺寸的值都共享同一个地址
 */
func TestZero1(t *testing.T) {
	// x , y 被分配在栈上
	var x, y, z, w [0]int
	println("===1==", g, &a, &b, &x, &y, &z, &w, f())
	g = &z
	g = &w // 使 z w 逃逸到堆上；所以可以相等么？
	println("===2==", g, &a, &b, &x, &y, &z, &w, f())
	println(&b == &a)  // false
	println(&x == &y)  // false
	println(&z == &w)  // true
	println(g == &x)   // false
	println(&z == f()) // false
}

/*
* 3. 不要把一个零尺寸字段作为struct的最后一个字段
*    为了防止内存泄漏，官方编译器将确保在取非零尺寸struct的最后一个地址时，不会返回超出为该结构体分配的内存块的地址。通过在最后一个零尺寸字段之后填充一些字节来实现。
*    同时遵循事实a：一个struct的地址对齐，保证为其字段的最大地址对齐。
*              b: 一个类型的尺寸总是该类型的地址对齐保证的倍数
 */
type Ty struct {
	_ [0]func()
	y int64
}
type Tz struct {
	z int64
	_ [0]func()
}

func TestZero2(t *testing.T) {
	var y Ty
	var z Tz
	println(unsafe.Sizeof(y)) // 8
	println(unsafe.Sizeof(z)) // 16

}

/*
* 4. for...range 模拟 for i in 0...N
 */
const N = 8

var n = 8

func TestZero3(t *testing.T) {
	for i := range [N]struct{}{} {
		println(i)
	}

	for i := range [N][0]int{} {
		println(i)
	}
	for i := range make([][0]int, n) {
		println(i)
	}
}

/*
* 5. 创建slice的多种方式
 */
func TestZero4(t *testing.T) {
	var s0 = make([]int, 10)
	var s1 = []int{1, 2, 3, 4}
	var s2 = (&[10]int{})[:]
	var s3 = new([10]int)[:]
	println(len(s0), len(s1), len(s2), len(s3))
}

/*
* 6. for i, v = range aContainer 实际上是迭代 aContainer 的副本
 */
func TestZero5(t *testing.T) {
	var a = [...]int{1, 2, 3}
	for i, n := range a {
		if i == 0 {
			a[1], a[2] = 8, 9
		}
		println(n) // 123 而不是 189（range中，数组是被复制的；但是如果忽略n（for i := range a），就不会被复制）
	}

	// 数组是基本类型？切片是引用类型。在Go中，值的复制都是浅拷贝。
	var s = []int{1, 2, 3}
	for i, n := range s {
		if i == 0 {
			s[1], s[2] = 8, 9
		}
		println(n) // 189
	}
}

/*
* 7. 某些情况，数组指针可以被当作数组使用
 */
func TestZero6(t *testing.T) {
	var a = [12]int{3: 789}
	var pa = &a
	// 1。 遍历数组指针
	for i, v := range pa { // 无需复制数值来遍历数组元素；复制一个指针的成本很低；
		println(i, v)
	}
	// 2。 获取数组的长度和容量
	_, _ = len(pa), cap(pa)
	// 3. 通过数组指针访问数组元素
	_ = pa[0]
	pa[3] = 555
	// 4. 从数组指针派送切片
	var _ []int = pa[:]

	// 5. 遍历 nil 数组指针时，如果第二个变量是被忽略的，将不会发生指针panic。
	var po *[5]string
	for i := range po {
		println(i) // 0 1 2 3 4
	}
	for i, _ := range po {
		println(i) // 0 1 2 3 4
	}
	for _, v := range po {
		_ = v // panic: runtime error: invalid memory address or nil pointer dereference.
	}
}

/*
* 8. unsafe.Sizeof、unsafe.Offsetof 和unsafe.Alignof 的调用都在编译时进行估值
 */
// a. 不会panic
func f1() {
	var v *int64 = nil
	println(unsafe.Sizeof(*v)) // 8
}
func g1() {
	var t *struct{ s [][16]int } = nil
	println(len(t.s[99]), unsafe.Sizeof(t.s[99])) // 16 128
}

// b. 会 panic
func f2() {
	var v *int64 = nil
	_ = *v
}
func g2() {
	var t *struct{ s [][16]int } = nil
	_ = t.s[99]
}

// c. 内置的len()会在for-range中被隐式调用。所以前两个循环正常，最后一个panic
type T3 struct {
	s []*[5]int
}

func g3() {
	var t *T3
	for i, _ := range t.s[99] { // len(t.s[99])、len(*t.s[99]) 的调用是编译时估值。此t.s[99]的长度5是确认的。
		print(i) // 0 1 2 3 4
	}
	for i := range *t.s[99] {
		print(i) // 0 1 2 3 4
	}
	// 而len(t.s)是在运行时估值的，当t是空指针的时候，会panic
	for i := range t.s { // panic: runtime error: invalid memory address or nil pointer dereference
		print(i)
	}
}
func TestZero7(t *testing.T) {
	f1()
	g1()
	f2() // panic: runtime error: invalid memory address or nil pointer dereference
	g2() // panic: runtime error: invalid memory address or nil pointer dereference
	g3() // panic: runtime error: invalid memory address or nil pointer dereference

	var c chan int
	var s []byte
	println(s, c)
	//const X = len([1]int{<-c})    // is not constant len包含通道读取操作，编译时就报错
	//const Y = cap([1]int{len(s)}) //is not constant len包含非常量函数调用，编译时就报错

	const X = 1 + 2i
	var y = 1 + 2i
	type A [8]float64

	var _ [len(A{imag(X)})]int
	var z = A{imag(y)}
	var _ [len(z)]int // 表达式len(z)不含有任何非常量函数调用，所以编译通过
	//var _ [len(A{imag(y)})]int // 编译不通过：y是变量，表达式含有非常量，所以不会在编译时被估值。 nvalid array bound 'len(A{imag(y)})', must be a constant expression
}

/*
* 9. 不支持声明超过2GB的**数组**
* 	  但是分配在堆上的数组可以?
 */
func TestZero8(t *testing.T) {
	//var x [20000000000000000 + 1]byte
	//println(len(x)) // type [20000000000000001]byte larger than address space

	var y *[20000000000000000 + 1]byte // 堆上的数组也不行？
	println(len(y))
}

/*
* 10. 切片、数组、map、struct字段的可寻址性
	a. 一个切片的元素总是可以寻址的，无论此切片本身是否可寻址。
    b. 可寻址的数组，其元素也是可寻址的；不可寻址的数组，其元素也是不可寻址的
    c. map的值的元素总是不可寻址的。
 */
type T9 struct {
	x int
}

func TestZero9(t *testing.T) {
	// a. 字面量是不可寻址的
	_ = &([10]bool{}[1]) // error: Cannot take the address of '[10]bool{}[1]'
	// 隐射元素?是不可寻址的
	var mi = map[int]int{1: 0}
	_ = &(mi[1]) // error: Cannot take the address of 'mi[1]
	var ma = map[int][10]bool{2: [10]bool{}}
	_ = &(ma[2][1])
	_ = &{T9{}.x}
	// 从不可寻址的数组中派生切片也是非法的
	var aSlice = [10]bool{}[:]

	var _ = &([]int{1: 0}[1]) // okay
	var a [10]bool
	_ = &(a[1]) // okay: 任何变量都是可以寻址的
	var t9 T9
	_ = &(t9.x)
}
