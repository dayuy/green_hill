package type_test

import (
	cm "github.com/orcaman/concurrent-map"
	"math"
	"testing"
)

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	// b = a // Cannot use 'a' (type int32) as the type int64 不能隐式类型转换
	b = int64(a) // 可以显式类型转换
	var c MyInt
	c = MyInt(b)

	aPrt := &a

	t.Logf("%T %T", a, aPrt)

	t.Log(a, b, c, aPrt, *aPrt)

	var s string // 默认为空字符串
	t.Log("*" + s + "*")

	t.Log(math.MaxUint8, math.MaxInt8) //255 127
}

func TestSliceComparing(t *testing.T) {
	a1 := []int{1, 2, 3, 4}
	b2 := []int{1, 2, 3, 4} //切片
	//slice can only be compared to nil 切片不能比较
	//if a1 == b2 {
	t.Log("equal", len(a1), len(b2), cap(a1), cap(b2))
	//}

	a := [...]int{1, 2, 3, 4} //自动初始化为n个int的数组
	b := [...]int{1, 3, 2, 4}
	c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b) // 顺序不同 false
	t.Log(c, cap(a)) // [1 2 3 4 5] 4
	// t.Log(a == c) 元素个数不相同的进行比较，会报错
	t.Log(a == d) // true

	t.Logf("%T %T", a, a1) // [4]int   []int
}

// 切片的内部结构
func TestSlice2(t *testing.T)  {
	var s0 []int // 切片：未赋值
	t.Log(len(s0), cap(s0)) // 0  0
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0)) // 1  1

	s1 := []int{1,2,3,4} // 切片：初始化 len\cap都为初始化长度
	t.Log(len(s1), cap(s1)) // 4 4

	s2 := make([]int, 3, 5) // 切片声明方式3
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1, 2, 3) // cap分别为 5 5 10；这也就是为什么append需要重新赋值，指向的存储空间每次cap翻倍的时候会变
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3])
}


/**
	Map
**/
// map的初始化方式
func TestInitMap(t *testing.T) {
	m1 := map[int]int{1:1, 2:4, 3:9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1]) // map中，如果此key不存在，则默认值为0
	m1[2] = 0
	t.Log(m1[2]) // 与本身值为0的key 无法区分
	// m1[3] = 0
	if v, ok := m1[3]; ok { // 所以，再取值的时候，会返回两个数，如果不存在 ok为false
		t.Logf("key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}

// Map 与 工厂模式
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

/**
	Set可以确定元素的唯一性
// Go的内置集合中没有Set，可以用 map[type]bool
**/
func TestMapForSet(t *testing.T) {
	mySet := map[int]int{}
	mySet[1] = 3
	n := 1
	if mySet[n] != 0 {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	mySet[3] = 6
	t.Log(len(mySet), mySet)

	delete(mySet, 1)
	t.Log(mySet[5])
	// map 中不存在的数为 0 或 false
}

func TestMapForSet1(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	mySet[3] = true
	t.Log(len(mySet))
	delete(mySet, 3)
	t.Log(len(mySet))
}


/**
	第三方包 ConcurrentMap
	thread-safe map 线程安全的map(读写相对频繁的情况下使用)
**/
func TestConcurrentMap(t *testing.T) {
	m := cm.New()
	m.Set("key", 10)
	t.Log(m.Get("key"))
}



/**
bool
string
int  int8  int16  int32  int64
uint uint8  uint16  uint32  uint64  uintptr
byte  // alias for unit8
rune  // alias for int32, represents a Unicode code point
float32  float64
complex64 complex128


1、Go 不允许隐式类型转换。可以显式的类型转换
2、别名和原有类型也不能进行隐式类型转换
**/
