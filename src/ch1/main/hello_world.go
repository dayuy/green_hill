package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// 浮点数
func main1() {
	fmt.Println("Hello world")
	var price float32 = 89.12
	fmt.Println("price=", price)
	var num1 float32 = -0.00089
	var num2 float64 = -7809656.09
	fmt.Println("num1=", num1, "num2=", num2)

	// 浮点数 = 符号位 + 指数位 + 尾数位
	// 尾数部分可能丢失，造成精度损失。 -123.0000901
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=", num3, "num4=", num4)

	// num1= -0.00089 num2= -7.80965609e+06
	// num3= -123.00009 num4= -123.0000901

	// 浮点型默认为float64
	var num5 = 1.1
	fmt.Printf("num5数据类型 %T \n", num5)

	num8 := 5.1234e2  // ? 5.1234 * 10的2次方
	num9 := 5.1234e-2 // 5.1234 / 10的2次方
	fmt.Println("num8=", num8, num9)
}

// 字符 byte
func main2() {
	var c1 byte = 'a'
	var c2 byte = '0' // byte 对应的是字符0，对应的是ASCII码中的值

	// 直接输出的是byte，对应的是字符的ASCII码值er
	fmt.Println("c1= ", c1) // 97
	fmt.Println("c2= ", c2) // 48
	// 格式化输出，可以输出对应字符
	fmt.Printf("c1=%c c2=%c\n", c1, c2) // a 0

	// var c3 byte = '北' // overflow溢出:constant 21271 overflows byte
	var c3 int = '北'
	fmt.Printf("c3=%c c3对应的码值=%d\n", c3, c3) // 北 21271

	var c4 int = 22269 // 22269对应的unicode字符为 ‘国’
	fmt.Printf("c4=%c", c4)
}

func main3() {
	var b = false
	fmt.Println("b=", b)

	// 1. bool类型占用存储空间是一个字节
	fmt.Println("b 的占用空间=", unsafe.Sizeof(b)) // 1

	var c int = '中'
	fmt.Println("c 的占用空间：", unsafe.Sizeof(c)) // 8
}

func main4() {
	var address string = "hello world"
	fmt.Printf(address)

	// 1、字符串一旦赋值了，就不能修改了。字符串不可变
	// var str = "hello"
	// str[0] = 'a'
	var hello string = "hello\nworld"
	fmt.Printf(hello)

	// 数据类型默认值
	var a int          // 0
	var b float32      // 0
	var c float64      // 0
	var isMarried bool // false
	var name string    // ""
	fmt.Println(a, b, c, isMarried, name)
	fmt.Printf("a=%d, b=%v, c=%v, isMarried=%v, name=%v", a, b, c, isMarried, name)

	// 1. 只能显示转换， 也没有低精度向高精度的隐性转换
	var i int32 = 100
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	// var n3 int64 = i  // 没有低精度到高精度
	fmt.Printf("i=%v n1=%v n2=%v", i, n1, n2)

	// 2. 被转换的是变量存储的数据（即值），变量本身的数据类型并没有变化
	fmt.Printf("i=%T\n", i)

	// 3. 转换中，比如将 int64 转成 int8 【-128 127】，编译时不会报错，结果按溢出处理
	var num1 int64 = 999999
	var num2 int8 = int8(num1) // 63 ?
	fmt.Println("num2=", num2)
}

func main5() {
	// 基本数据类型转string
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = '1'
	var str string

	// 1、fmt.Sprintf()
	str = fmt.Sprintf("%d", num1) // Sprintf 根据format参数生成格式化字符串并返回该字符串
	fmt.Printf("str type %T str=%v\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type= %T str=%v\n", str, str)

	str = fmt.Sprintf("%t", b) // bool
	fmt.Printf("str type= %T str=%v\n", str, str)

	str = fmt.Sprintf("%s", myChar) // byte
	fmt.Printf("str type= %T str=%q\n", str, str)

	// 2、strconv 函数: 其他类型转string
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q\n", str, str)

	// strconv 函数：string类型转其他
	var str1 string = "true"
	var b1 bool
	b1, _ = strconv.ParseBool(str1)
	fmt.Printf("b type: %T b=%v\n", b1, b1)

	var str2 string = "1234590"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 type %T n1=%v\n", n1, n1)
	fmt.Printf("n2 type %T n2=%v\n", n2, n2)

	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1=%v\n", f1, f1)
}

// 指针类型
func main6() {
	// 基本数据类型在内存中的分配
	var i int = 10
	// i 的地址是什么，&i
	fmt.Println("i的地址：", &i) // 0xc0000b6008

	var ptr *int = &i                 // ptr是一个指针变量
	fmt.Printf("ptr=%v\n", ptr)       // 0xc0000b6008
	fmt.Printf("ptr 的地址=%v\n", &ptr)  // 0x00000e030
	fmt.Printf("ptr 指向的值=%v\n", *ptr) // 10

	// 1. 修改指针内的值，会修改原来的数
	var num int = 9
	fmt.Printf("num address=%v\n", &num)

	var ptr1 *int
	ptr1 = &num
	*ptr1 = 10 // 这里修改时，会影响到num的值
	fmt.Println("num = ", num, " *ptr1 = ", *ptr1)
}

// 运算
func main7() {
	// 1、「/、%」如果运算的数都是整数，那么除以后也会去掉小数部分，保留整数部分
	fmt.Println(10 / 4) // 2

	var n1 float32 = 10 / 4
	fmt.Println(n1) // 2

	// 如果需要保留小数部分，则需要有浮点数参与运算
	var n2 float32 = 10.0 / 4
	fmt.Println(n2) // 2.5

	// a % b = a - a/b*b
	fmt.Println("10%3=", 10%3)     // 10 - 10/3*3 = 1
	fmt.Println("-10%3=", -10%3)   // -10 - -10/3*3 = -10 - -9 = -1
	fmt.Println("10%-3=", 10%-3)   // 10 - 10/-3*-3 = 1
	fmt.Println("-10%-3=", -10%-3) // -10 - (-10/-3*-3) = -10 - -9 = -1

	// 3、++ 和 -- 只能独立使用，且只有后++、后--
	var i int = 8
	var a int
	// a = i++ 错误
	i++
	fmt.Println("i++: ", i, a)

	// 例: 交换两个变量，不允许使用中间变量
	var m int = 10
	var n int = 20
	m = m + n
	n = m - n
	m = m - n
	fmt.Printf("m=%v n=%v", m, n)

	/*
	* 运算符的优先级
	* 1、（）、++、--
	* 2、单目运算符、赋值运算符是从右向左运算的
	* 3、算术运算符 * / % + -
	* 4、位移运算 <<  >>
	* 5、关系运算 < <=  == !=
	* 6、位运算 & ^ |
	* 7、逻辑运算  &&  ||
	* 8、赋值运算  =  += -=  *=  /=
	 */

	// 用户输入 fmt.Scanln()
	var name string
	var age byte
	fmt.Println("请输入名字：")
	fmt.Scanln(&name)
	fmt.Printf("您的名字是：%v \n", name)
	fmt.Println("请输入年了：")
	fmt.Scanln(&age)
	fmt.Printf("您的名字是：%v \n", age)

	fmt.Println("请输入您的名字，年龄")
	fmt.Scanf("您的名字：%s 年龄是：%d", &name, &age)
}

// 进制
func main8() {
	var i int = 5
	// 二进制输出
	fmt.Printf("%b \n", i)

	// 八进制：0-7，满8进1，以数字0开头表示
	var j int = 011 // 9
	fmt.Println("j=", j)

	// 16进制，以0x或0X开头表示
	var k int = 0x11 // 0x11 => 16 + 1 = 17
	fmt.Println("k=", k)
}

func main9() {
	b := strings.Contains("seafood", "sea")
	fmt.Printf("b=%v\n", b)

	num := strings.Count("cheese", "e")
	fmt.Printf("num=%v\n", num)

	b = strings.EqualFold("abc", "Abc")
	fmt.Printf("b=%v\n", b)

	fmt.Println("result: ", "abc" == "Abc")

	index := strings.Index("NLT_abcabcabc", "abc")
	fmt.Printf("index=%v\n", index)

	index = strings.LastIndex("NLT_abcabcabc", "abc")
	fmt.Printf("index=%v\n", index)

	str2 := "go go help"
	str := strings.Replace(str2, "go", "北京", -1)
	fmt.Printf("str=%v str2=%v\n", str, str2)

	strArr := strings.Split("hello,world,ok", ",")
	fmt.Printf("strArr=%v\n", strArr)
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
	}

	str = "golang Hello"
	str = strings.ToLower(str)
	fmt.Printf("str=%v\n", str)

	str = strings.TrimSpace(" the a lone gopher turn    ")
	fmt.Println(str)

	// 去掉空格，同时去掉！
	str = strings.Trim("! hel!lo! ", " !")
	fmt.Println(str)

	b = strings.HasPrefix("ftp://192.168.10.1", "ftp")
	fmt.Printf("b=%v\n", b)
}

func test03() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main10() {
	now := time.Now()
	fmt.Printf("now=%v now type=%T\n", now, now)

	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	// 格式化
	fmt.Printf(now.Format("2006/01/02 15:04:05")) // 2006/01/02 15:04:05 数字必须是这个固定的。
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()

	time.Sleep(time.Millisecond + 100)

	fmt.Printf("unix时间戳=%v unixnano时间戳=%v\n", now.Unix(), now.UnixNano())

	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Printf("执行test03()耗费时间为%v秒\n", end-start)
}

func main11() {
	num1 := 100
	fmt.Printf("num1的类型%T, num1的值=%v, num1的地址%v\n", num1, num1, &num1)

	/*
	* 内置函数
	* 1、len：用来求长度
	* 2、new：用来分配内存，主要用来分配值类型：int、float32、struct，返回的是指针。
	* 3、make：用来分配内存，主要用来分配引用类型：chan、map、slice
	 */
	num2 := new(int) // *int
	// num2的类型%T => *int
	// num2的值 = 地址0xc00001e0c0 （此地址是系统分配）
	// num2的地址%v = 地址0xc00000e030
	// num2指向的值 = 0
	*num2 = 100
	// num2指向的值 = 100
	fmt.Printf("num2的类型%T, num2的值=%v, num2的地址%v，num2这个指针指向的值=%v", num2, num2, &num2, *num2)
}

func test() {
	// 使用defer + recover 来捕获和处理异常
	defer func() {
		err := recover() // recover() 内置函数，可以捕获到异常
		if err != nil {
			fmt.Println("err= ", err)
		}
	}()

	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res= ", res)
}

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文章失败。。。")
	}
}

func test02() {
	err := readConf("config2.ini")
	if err != nil {
		panic(err) // 就不再往后执行了
	}
	fmt.Println("test02() 继续执行。。。")
}

func main12() {
	/*
	* 错误处理
	* 1、不支持传统的try...catch...
	* 2、Go中引用的处理方式为：defer、panic、recover
	*    Go中可以抛出一个异常，然后在defer中通过recover捕获这个异常，然后正确处理
	 */
	//test()
	test02()
	fmt.Println("main() 下面的代码")
}

func main() {
	var intArr [4]int
	// int 占8个字节, 定义完数组后，数组的默认值为0
	fmt.Println(intArr)
	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	intArr[3] = 40
	fmt.Printf("intArr的地址=%p intArr[0]的地址=%p, int[1]=%p, int[2]=%p, int[3]=%p\n", &intArr, &intArr[0], &intArr[1], &intArr[2], &intArr[3])

	// 初始化数组
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01=", numArr01)

	var numArr02 = [3]int{1, 2, 3}
	fmt.Println("numArr02=", numArr02)

	var numArr03 = [...]int{1, 2, 3}
	fmt.Println("numArr03=", numArr03)

	var numArr04 = [...]int{1: 800, 0: 900, 2: 999}
	fmt.Println("numArr04=", numArr04)

	numArr05 := [...]int{4, 5, 3, 2}
	fmt.Println("numArr05=", numArr05)
}
