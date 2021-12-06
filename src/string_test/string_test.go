package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s) // 初始化为默认0值
	s = "hello"
	t.Log(len(s))
	// s[1] = '3' string 是不可变的byte slice
	s = "\xE4\xB8\xA5"
	t.Log(s)      // 严
	t.Log(len(s)) // 3

	s = "中"
	t.Log(len(s), s)     // 中文的byte数: 3
	t.Logf("UTF: %x", s) // e4 b8 ad

	c := []rune(s) // rune：取出字符串里的unicode
	t.Log(len(c))  // 1

	t.Logf("中 unicode %x", c[0])      // 「中」以‘unicode’的形式存储为 4e2d
	t.Logf("中 UTF8 %x %d", s, len(s)) // 「中」以‘UTF8’的形式存储为 e4 b8 ad 3
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
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
	if i, err := strconv.Atoi("10"); err == nil { // 字符串转为数字 strconv.Atoi()
		t.Log(10 + i)
	}
}
