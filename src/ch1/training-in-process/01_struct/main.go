package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
	"unsafe"
)

type person struct {
	name string
	age  int
}

type mySentence string

type mySentences []string

type myType int
type myTypes []int

func main() {
	p1 := person{"James", 20}
	fmt.Println(p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)
	fmt.Printf("%T\n", p1)

	var message mySentence = "Hello World!"
	fmt.Println(message)
	fmt.Printf("%T\n", message)
	fmt.Printf("%T\n", string(message)) // main.mySentence -> string
	//fmt.Printf("%T\n", message.(string)) // non-interface type mySentence on the left

	var p2 person
	fmt.Println(p2)
	fmt.Println(p2.name)
	fmt.Println(p2.age) // age 初始化为0
	fmt.Printf("%T\n", p2)

	p3 := person{
		name: "James",
		age:  20,
	}
	fmt.Println(p3)
	fmt.Println(p3.name)
	fmt.Println(p3.age)
	fmt.Printf("%T\n", p3) // main.person

	var messages mySentences = []string{"Hello World!", "More coffee"}
	fmt.Println(messages)
	fmt.Printf("%T\n", messages)
	fmt.Printf("%T\n", []string(messages)) // main.mySentences -> []string

	var x myType = 32
	fmt.Println(x)
	fmt.Printf("%T\n", x)         // main.myType
	fmt.Printf("%T\n", int(x))    // main.myType -> int
	fmt.Printf("%T\n", string(x)) // main.myType -> string

	var xs myTypes = []int{32, 44, 57}
	fmt.Println(xs)
	fmt.Printf("%T\n", xs)
	fmt.Printf("%T\n", []int(xs)) // main.myTypes -> []int

	// slice 声明切片
	slice1 := make([]float32, 0)
	slice2 := make([]float32, 3, 5)
	fmt.Println(len(slice1), cap(slice1))
	fmt.Println(len(slice2), cap(slice2))

	// 使用切片
	slice2 = append(slice2, 1, 2, 3, 4)
	fmt.Println(len(slice2), cap(slice2)) //
	// 子切片
	sub1 := slice2[3:]
	sub2 := slice2[:3]
	// 合并切片
	combined := append(sub1, sub2...) // sub2... 切片解构
	fmt.Println(combined)

	// 字典 map
	// 仅声明
	m1 := make(map[string]int)
	// 声明时初始化
	m2 := map[string]string{
		"Sam":   "Male",
		"Alice": "Female",
	}
	// 赋值/修改
	m1["Tom"] = 19
	fmt.Println(m2)

	// 指针（Pointer）
	// 指针即某个值的地址，类型定义时使用符合*，对一个已经存在的变量，使用&获取该变量地址。
	str := "Golang"
	var p *string = &str // 通过&获取一个变量的地址
	fmt.Println(str, p)
	*p = "Hello" // 通过 *获取一个地址的值
	fmt.Println(str, p)

	num := 100
	add(num)
	fmt.Println("num===", num)
	realAdd(&num)
	fmt.Println(num)

	fmt.Println(get(5))
	fmt.Println("finished")

	// 接口
	var person1 Person = &Student{
		name: "Tom",
		age:  18,
	}
	fmt.Println(person1.getName())
	// 1. 检测某个类是否实现类接口的所有方法: 将nil转换为Student类型，再转换为Person接口，如果失败则说明并没有实现接口上的所有方法
	var _ Person = (*Student)(nil)
	//var _ Person = (*Worker)(nil)

	// 2. 实例可以强制类型转换为接口，接口也可以强制类型转换为实例
	stu1 := person1.(*Student) // 接口转为实例
	fmt.Println(stu1.getName())

	// 3. 如果定义了一个没有任何方法的空接口，那么这个接口可以表示任意类型。
	m3 := make(map[string]interface{})
	m3["name"] = "Tom"
	m3["age"] = 18
	m3["scores"] = [3]int{98, 99, 85}
	fmt.Println(m3)

	// 并发编程
	goroutine1()

	// interview
	goInterview()

	// defer
	fmt.Println("return", goDefer())

	// slice equal
	slice3 := make([]string, 3)
	slice4 := make([]string, 3)
	slice3 = append(slice3, "a", "b")
	slice4 = append(slice4, "a", "b", "c")
	b := StringSliceEqual(slice3, slice4)
	fmt.Println("slice equal: ", b)

	// 表示枚举
	type StuType int32
	const (
		Type1 StuType = iota
		Type2
		Type3
	)
	fmt.Println(Type1, Type2, Type3)

	// 空struct{}用途
	// 1。可以节省内存，一般作为占位符使用，表面这里并不需要一个值
	fmt.Println(unsafe.Sizeof(struct{}{})) // 0
	// 2。适用：如map，可以使用struct{}作为占位符。若int或bool会浪费内存和歧义
	set := make(map[string]struct{})
	for _, item := range []string{"a", "b", "c"} {
		set[item] = struct{}{}
	}
	fmt.Println(len(set)) // 3
	if _, ok := set["a"]; ok {
		fmt.Println("a exists")
	}
	// 3. 适用：channel控制并发时，只想要一个信号，并不想传递值，也可以使用struct{}代替。
	ch2 := make(chan struct{})
	go func() {
		<-ch2
		// do...
	}()
	ch2 <- struct{}{}
	// do...
	// 4. 适用：声明只包含方法的结构体
	type Lamp struct{}
	//func (l Lamp) On() {
	//	println("on")
	//}

	// channel
	channelBuffer()
}
func add(num int) {
	num += 1
}
func realAdd(num *int) {
	*num += 1
}

// 不可预知的错误 panic，其他语言用try catch处理，这里用defer recover机制
func get(index int) (ret int) { // 函数中也可以给返回值命名，简化return
	defer func() {
		if r := recover(); r != nil { // defer... recover 扑获错误
			fmt.Println("Some error happened!", r)
			//ret = -1 默认为0
		}
	}()
	arr := [3]int{2, 3, 4}
	return arr[index]
}

// 接口（interface）
// 接口定义了一组方法的集合，接口不能被实例化，一个类型可以实现多个接口
type Person interface {
	getName() string
}
type Student struct {
	name string
	age  int
}

func (s *Student) getName() string {
	return s.name
}

type Worker struct {
	name   string
	gender string
}

// 并发编程：sync和channel两种方式都可以支持协程（goroutine）的并发。
// 场景1：希望并发下载N个资源，多个并发协程之间不需要通信，就可以使用sync.WaitGroup，等待所有并发协程执行结束。
var wg sync.WaitGroup

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	wg.Done() // 2.为wg减去一个计数
}
func goroutine1() {
	for i := 0; i < 3; i++ {
		wg.Add(1)                             // 1.为wg添加一个计数
		go download("a.com/" + string(i+'0')) // 启动新的协程并发执行download
	}
	wg.Wait() // 等待所有的协程执行结束
	fmt.Println("Done!")
	fmt.Println(runtime.NumGoroutine())
}

// 场景2：channel。可以在协程之间传递消息；阻塞等待并发协程返回消息。
var ch = make(chan string, 10) // 创建大小为 10 的缓冲信道
func download2(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	ch <- url // 将url发送给信道
}
func goroutine2() {
	for i := 0; i < 3; i++ {
		go download2("a.com/" + string(i+'0'))
	}
	for i := 0; i < 3; i++ {
		msg := <-ch // 阻塞：等待信道返回信息
		fmt.Println("finish", msg)
	}
	fmt.Println("Done!")
}

// test
func goInterview() {
	// 1. 如何高效拼接字符串：Go中，字符串是只读的，意味着每次修改操作都会创建一个新的字符串，如果需要拼接多次，应该使用strings.Builder，最小化内存拷贝次数
	var str strings.Builder
	for i := 0; i < 1000; i++ {
		str.WriteString("a")
	}
	fmt.Println(str.String())

	// 2. 什么是rune类型
	// a. ASCII码只需要7bit 但只表示的是英文字母在内的128个字符，为表示世界上大部分文字系统，发明了Unicode，它是ASCII的超集；
	//    并为每个字符分配了一个标准编号（Unicode CodePoint）,在Go中称之为rune，为int32的别名
	// b. GO中字符串的底层表示是byte(8 bit)序列，而非rune（32 bit）序列。
	//	  中文使用UTF-8编码各占3个byte，所以len("Go语言")==8
	//    也可以将字符串转换为rune序列
	fmt.Println(len("Go语言"))         // 8
	fmt.Println(len([]rune("Go语言"))) // 4
}

/*
*
defer
1.多个defer语句，执行顺序：后进先出
2.defer 在return语句之后执行，但在函数退出之前，defer可以修改返回值
*/
func goDefer() int {
	i := 0
	defer func() {
		fmt.Println("defer 1")
	}()
	defer func() {
		i++ // 1
		fmt.Println("defer 2")
	}()
	return i // return 0
	// defer2
	// defer1
	// return 0
	// 这里返回值没有被修改，是因为Go的返回机制，执行 return 语句后，Go 会创建一个临时变量保存返回值，因此，defer 语句修改了局部变量 i，并没有修改返回值
}

func goDefer2() (i int) {
	i = 0
	defer func() {
		i += 1
		fmt.Println("defer2")
	}()
	return i
	// defer2
	// return 1
	// 这里的返回值被修改了，对于有名返回值的函数，执行 return 语句时，并不会再创建临时变量保存，因此，defer 语句修改了 i，即对返回值产生了影响
}

/*
*
如何判断两个字符串slice是相等的？
1。使用反射 reflect.DeepEqual(a, b), 但是不推荐，较耗性能
2。遍历
*/
func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

/*
* init()函数是什么时候执行的？
* 1。import -> const -> var -> init() -> main()
* 解释：每个包首先初始化包作用域的常量和变量（常量优先于变量），然后执行包的init()函数。
*	   同一个包或同一个源文件可以有多个init()函数。
*		init()函数没有入参和返回值，不能被其他函数调用。
* 		同一个包内多个init()函数执行顺序不能保证。
 */
/*
* 两个interface能否比较？能，使用==或!=
* 1。interface内部实现了2个字段：类型T 和 值V。
* 2。两种情况下相等
*    a.两个interface均等于nil（此时V、T都处于unset状态）
*	 b.类型T相同，且对应的值V相等。
 */
func equalInterface() {
	type Stu struct {
		Name string
	}
	type IntStu interface{}
	var stu1, stu2 IntStu = &Stu{"Tom"}, &Stu{"Tom"}
	var stu3, stu4 IntStu = Stu{"Tom"}, Stu{"Tom"}
	fmt.Println(stu1 == stu2) // false  ：类型 是 *Stu，值是地址
	fmt.Println(stu3 == stu4) // true   ：类型是Stu，值是Stu{"Tom"}，且各个字段相等
}

/*
* 函数返回局部变量的指针是否安全？
* 1。在go中是安全的，Go编译器将会对每个局部变量进行逃逸分析。如果局部变量的作用域超出该函数，则不会将内存分配在栈上，而是堆上。
 */
func foo1() *int {
	v := 11
	return &v
}
func boo1() {
	m := foo1()
	println(*m) // 11
	// foo() 函数中，如果 v 分配在栈上，foo 函数返回时，&v 就不存在了，但是这段函数是能够正常运行的。Go 编译器发现 v 的引用脱离了 foo 的作用域，会将其分配在堆上。
}

/*
* Channel
* 1。无缓冲channel：发送方将阻塞该信道，直到接收方从该信道接受到数据。而接收方也将阻塞该信道，直到发送方将数据发送到该信道。
* 2。有缓冲channel：发送方在缓冲区用完的情况下阻塞，而接收方在信道为空的情况下阻塞
 */
func channelno() {
	str := time.Now()
	ch := make(chan bool)
	go func() {
		time.Sleep(time.Second * 4)
		<-ch
	}()
	ch <- true
	// 阻塞4s
	fmt.Printf("cost %.1f s\n", time.Now().Sub(str).Seconds())
	time.Sleep(time.Second * 5)
}
func channelBuffer() {
	st := time.Now()
	ch := make(chan bool, 2)
	go func() {
		time.Sleep(time.Second * 2)
		<-ch
	}()
	ch <- true
	ch <- true
	// 不阻塞
	fmt.Printf("cost %.1f s\n", time.Now().Sub(st).Seconds())
	ch <- true
	// 阻塞2s
	fmt.Printf("cost %.1f s\n", time.Now().Sub(st).Seconds())
	time.Sleep(time.Second * 5)
}
