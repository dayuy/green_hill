package string_test

import (
	"strconv"
	"strings"
	"testing"
)

/**
	1、string 是数据类型，不是引用或指针类型
	2、**string 是只读的 byte slice***，len函数显示它所包含的byte数
	3、string 的byte数组可以存放任何数据
**/

func TestString(t *testing.T) {
	var s string
	t.Log(s) // 初始化为默认0值
	s = "hello"
	t.Log(len(s))
	// s[1] = '3' string 是不可变的byte slice
	s = "\xE4\xB8\xA5"
	t.Log(s)      // 严
	t.Log(len(s)) // 3  len表示的是byte数，而不是字符数

	s = "中"
	t.Log("中：", len(s), s)     // 中文的byte数: 3
	t.Logf("UTF: %x", s) // e4 b8 ad

	c := []rune(s) // rune：取出字符串里的**unicode**
	t.Log(len(c))  // 1

	t.Logf("中 unicode %x", c[0])      // 「中」以‘unicode’的形式存储为 4e2d
	t.Logf("中 UTF8 %x %d", s, len(s)) // 「中」以‘UTF8’的形式存储为 e4 b8 ad 3
}

/**
	Unicode是一种字符集（code point）
	UTF8是unicode的存储实现 （转换为byte序列的规则）
	字符    「中」
	Unicode 0x4E2D
	UTF-8   0xE4B8AD
	string/[]byte [0xE4,0xB8,0xAD]
**/

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s { // 遍历字符串 默认转换为rune，而不是byte
		t.Logf("%[1]c %[1]d", c) // 表示第一个参数以%c格式化，以%d格式化
	}
}

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConvFn(t *testing.T) {
	s := strconv.Itoa(10) // 数字转换为字符串 strconv.Itoa()
	t.Log("str" + s)
	// d := 10 //Invalid operation: "int" + d (mismatched types string and int)
	// t.Log("int" + d) // mismatched types string and int
	if i, err := strconv.Atoi("10"); err == nil { // 字符串转为数字 strconv.Atoi()
		t.Log(10 + i)
	}
}

const (
	Readable = 1 << iota  // 1 << 0
	Writable // 1 << 1  0010 //2
	Executable  // i << 2   0100 // 4
)

func TestConstantTry(t *testing.T) {
	a := 1 // 0001
	t.Log(a, Readable, Writable, Executable) // 1 1 2 4
	t.Log(a&Readable, a&Writable, a&Executable) // 1 0 0
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable) // true false false

	b := 7 // 0111
	b = b &^ Readable
	t.Log(b&Readable, b)
}
