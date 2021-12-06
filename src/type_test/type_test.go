package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)

	aPrt := &a

	t.Logf("%T %T", a, aPrt)

	t.Log(a, b, c, aPrt)

	var s string
	t.Log("*" + s + "*")
}

func TestSliceComparing(t *testing.T) {
	a1 := []int{1, 2, 3, 4}
	b2 := []int{1, 2, 3, 4}
	// slice can only be compared to nil
	// if a1 == b2 {
	t.Log("equal", len(a1), len(b2), cap(a1), cap(b2))
	// }

	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	t.Log(c)
	// t.Log(a == c)
	t.Log(a == d)
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	// m1[3] = 0
	if v, ok := m1[3]; ok {
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

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

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
