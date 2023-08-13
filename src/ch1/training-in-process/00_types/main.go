package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(32 / 16)
	fmt.Println(32 / 3.74)
	answer := 32 / 3.74
	fmt.Println(answer)

	fmt.Println("Hello\tWorld\nHow are you?")

	intro := "Four score and seven years ago..."
	fmt.Println(intro)
	fmt.Println(&intro)
	intro = "Hahaha!" // 字符串也可以变更？
	fmt.Println(intro)
	fmt.Println(&intro)
	fmt.Println([]byte(intro))
	fmt.Println(string(111))

	// runes
	letter := 'A'
	fmt.Println(letter)         // 65
	fmt.Printf("%T \n", letter) // int32

	letter1 := rune("A"[0])
	fmt.Println(letter1, "A"[0]) // 65 65   "A"[0]也就是runes

	word := "hello 你好"
	letter2 := rune(word[1])
	fmt.Println(letter2)

	/*
	* byte 数据类型与 rune相似，都是用来表示字符类型。
	* byte 等同于uint8,常用来处理ascii字符
	* rune 等同于int32, 常用来处理unicode或utf8字符
	 */
	fmt.Println("len(word): ", len(word))                            // 12
	fmt.Println("rune: ", len([]rune(word)), []rune(word))           // 8 (golang中string底层是通过byte数组实现的。中文字符在unicode下占2个字符，在utf-8编码下占3个字符，而golang默认是utf-8)
	fmt.Println("RuneCountInString: ", utf8.RuneCountInString(word)) // 8

	fmt.Println(len("hello"))
	fmt.Println(len("世界")) // go中默认在utf-8中占3个字符，所以这里是6

	intro = "Four 世"
	fmt.Printf("%T\n", intro)
	fmt.Println(intro)
	bs := []byte(intro)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)            // []unit8
	fmt.Printf("%T\n", []rune(intro)) // []int32
	fmt.Println("***********")
	fmt.Printf("%d\n", bs)

	fmt.Println("*****====================******")
	// 字符串是以byte数组形式存储的，类型是unit8，英文字符占1个byte，中文一般是占3个，打印时需要用string进行类型转换，否则打印的是编码值
	fmt.Println(intro)                           // Four 世
	fmt.Println([]byte(intro))                   // [70 111 117 114 32 228 184 150]
	fmt.Println(reflect.TypeOf(intro[5]).Kind()) // unit8
	fmt.Println(intro[5], string(intro[5]))
	fmt.Println("len(intro): ", len(intro))
	// 正确的方式是把string转为rune数组
	// []rune 类型，字符串中的每个字符，无论占多少个字节都用int32表示，因而可以正确处理中文
	runeArr := []rune(intro)
	fmt.Println(runeArr) // [70 111 117 114 32 19990]
	fmt.Println(reflect.TypeOf(runeArr[5]).Kind())
	fmt.Println(runeArr[5], string(runeArr[5]))
	fmt.Println("len(runeArr): ", len(runeArr))

	for _, v := range bs {
		// %#x: 十六进制前加0X
		// %#o：八进制前加0
		fmt.Printf("%d\t\t %#x\t %b\n", v, v, v)
	}
	fmt.Println("**********")
	y := 999999999999999999

	fmt.Printf("%d\t\t %#x\t %b\n", y, y, y)
	fmt.Println(&y)
	fmt.Sprint(y)
	fmt.Println("*********")

	z := 'h'
	fmt.Printf("%T\n", z) // int32

	greeting := "Hello"
	fmt.Println(greeting)
	fmt.Println(greeting[0]) // rune
	fmt.Println(greeting[4])
	fmt.Println("-------------")
	fmt.Println("What the ... ")
	fmt.Println(greeting[:4])
	fmt.Println("... did that just do?")
	fmt.Println(greeting[0:4]) // 左包含 右舍弃 （左闭右开）
	fmt.Println(greeting[1:4])
	fmt.Println(greeting[1:])
	//fmt.Println(greeting[:-2]) // must be non-negative
	fName := "James"
	lName := "Bond"
	fmt.Println(fName + " " + lName)

	var x int = 5
	// int -> string : 1. strconv.Itoa(x)  2. fmt.Sprint(x)
	str := "Hello world " + strconv.Itoa(x)
	fmt.Println(str)
	str = "Hello world " + fmt.Sprint(x)
	fmt.Println(str)

	// string -> int:  1. strconv.Atoi("32")
	i, _ := strconv.Atoi("32")
	fmt.Println(i + 10)

	up := 34.705945
	down := 34.405945
	fmt.Println(math.Floor(up + 0.5))
	fmt.Println(math.Floor(down + 0.5)) // 向下取整
	fmt.Println(math.Ceil(up))          // 向上取整
	fmt.Println(math.Ceil(down))
	fmt.Println(math.Round(up)) // 四舍五入
	fmt.Println(math.Round(down))

	a := "this is stored in the variable a"
	b := 42
	c, d, e, f := 44.7, true, false, 'm'
	g := "g"
	h := `h`

	// 判断类型：1. reflect.TypeOf(a)   2. fmt.Printf("%T", h)
	fmt.Println("a - ", reflect.TypeOf(a), " - ", a)
	fmt.Println("b - ", reflect.TypeOf(b), " - ", b)
	fmt.Println("c - ", reflect.TypeOf(c), " - ", c)
	fmt.Println("d - ", reflect.TypeOf(d), " - ", d)
	fmt.Println("e - ", reflect.TypeOf(e), " - ", e)
	fmt.Println("f - ", reflect.TypeOf(f), " - ", f) // int32
	fmt.Println("g - ", reflect.TypeOf(g), " - ", g)
	fmt.Println("h - ", reflect.TypeOf(h), " - ", h)
	fmt.Printf("h - %T\n", h)

	fmt.Println("==========++++=========")
	result := isEqual("abcd", "dcba")
	fmt.Println(result)
}

func isEqual(s string, g string) bool {
	if len(s) != len(g) {
		return false
	}
	i := 0
	for i < len(s) {
		if in := strings.IndexByte(g, s[i]); in < 0 {
			return false
		}
		for j := i + 1; j < len(s); j++ {
			s1 := []byte(s)
			s1[i], s1[j] = s1[j], s1[i]
			if string(s1) == g {
				return true
			}
		}
		i++
	}
	return false
}
