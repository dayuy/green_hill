package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unicode/utf8"
)

func main() {
	//binaryOne()

	//shortHand()
	//
	//zeroValue()
	//
	allTogether()
	//
	//iotaT()
	//
	//greet()

	//map_exercise()
	//
	//scannerExercise()

	//intSlicePlus()

	//mapBucket()

	//englishAlphabet()

	//structDefined()

	//interfaceExercise()

	//sortPackage()

	//standardLibraryExample()

	//goRoutines()

	//channelsExample()
}

func binaryOne() {
	fmt.Printf("%d - %b \n", 42, 42)
	fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#x \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#X \n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)

	for i := 1000000; i < 1000100; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}

	for i := 60; i < 122; i++ {
		// d 十进制、b 二进制、x 十六进制、q 对应的ASCII数字
		// decimal、binary、hexadecimal
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}

	fmt.Println(reverseTwo("hello world"))
}

func reverseTwo(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func shortHand() {
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "hello"
	f := `Do you like my hat?`
	g := 'M' // 单引号

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g) // 77

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", g) // int32
}

func zeroValue() {
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("%v \n", a) // 初始化值 0
	fmt.Printf("%v \n", b) // ''
	fmt.Printf("%v \n", c) // 0
	fmt.Printf("%v \n", d) // false

	fmt.Println()

	var message = "hello world"
	var e, f, g = 1, false, 3
	fmt.Println(message, e, f, g)
}

var a = "this is stored in the variable a"
var b, c string = "stored in b", "stored in c"
var d string

func allTogether() {
	d = "stored in d"
	var e = 42
	f := 43
	g := "stored in g"
	h, i := "stored in h", "stored in i"
	j, k, l, m := 44.7, true, false, 'm'
	n := "n"
	o := `0`

	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("c - ", c)
	fmt.Println("d - ", d)
	fmt.Println("e - ", e)
	fmt.Println("f - ", f)
	fmt.Println("g - ", g)
	fmt.Println("h - ", h)
	fmt.Println("i - ", i)
	fmt.Println("j - ", j)
	fmt.Println("k - ", k)
	fmt.Println("l - ", l)
	fmt.Println("m - ", m) // 字符 单引号 int32
	fmt.Println("n - ", n)
	fmt.Println("o - ", o)

	fmt.Printf("%T \n", m)
	fmt.Printf("%T \n", n)

	name := "Todd"
	fmt.Println("hello ", name)

	p := make(map[string]int)
	p["Todd"] = 44
	fmt.Println("p", p)

	q := make([]string, 1, 25)
	fmt.Println("q: ", q, len(q), cap(q))
	q = append(q, "a")
	fmt.Println("q: ", q, len(q), cap(q))
}

func iotaT() {
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	d := 43
	fmt.Println(d)
	fmt.Println(&d)

	var e = &d
	fmt.Println(e, *e)
	fmt.Printf("%T ; %T", e, *e) // *int

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	f := true
	if food := "Chocolate"; f {
		fmt.Println(food)
	}
	//fmt.Println(food)

	data := []float64{43, 56, 87, 12, 45, 57}
	n := average(data...) // data... 表示？
	// n := average(data)
	fmt.Println(n)

	name := "Todd"
	fmt.Println(name)
	changeMe(name)
	name = "Todd a"
	fmt.Println(name)
}

// func average(sf []float64) float64 {
func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func changeMe(z string) {
	fmt.Println(z)
	z = "Rocky"
	fmt.Println(z)
}

func greet() {
	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	for i, currentEntry := range greeting {
		fmt.Println(i, currentEntry)
	}

	for i := 0; i < len(greeting); i++ {
		fmt.Println(greeting[i])
	}

	fmt.Print("[1:2] ")
	fmt.Println(greeting[1:2])
	fmt.Print("[:2] ")
	fmt.Println(greeting[:2])
	fmt.Print("[5:] ")
	fmt.Println(greeting[5:])
	fmt.Print("[:] ")
	fmt.Println(greeting[:])

	// customerNumber
	customerNumber := make([]int, 3) // 3 is length & capacity
	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 15
	customerNumber = append(customerNumber, 5) // ! append is ok
	fmt.Println(customerNumber)

	greeting1 := make([]string, 3, 5) // 3 is length   5 is capacity
	greeting1[0] = "Good morning!"
	greeting1[1] = "Bonjour!"
	greeting1[2] = "dias!"
	//greeting1[3] = "superman"  // panic: runtime error: index out of range [3] with length 3
	fmt.Println(greeting1[2])
	greeting1 = append(greeting1, "superman") // append is ok
	fmt.Println(greeting1[3])

	// length、capacity
	greeting2 := make([]string, 3, 5)
	greeting2[0] = "Good morning!"
	greeting2[1] = "Bonjour!"
	greeting2[2] = "buenos dias!" // [index] 不能超过3个，但是append()却可以
	greeting2 = append(greeting2, "superman")
	greeting2 = append(greeting2, "Zao'an")
	greeting2 = append(greeting2, "Ohayou gozaimasu")
	greeting2 = append(greeting2, "qidday")

	fmt.Println(greeting2[6])
	fmt.Println(len(greeting2))
	fmt.Println(cap(greeting2))

	mySlice := []int{1, 2, 3, 4, 5}
	myOtherSlice := []int{6, 7, 8, 9}
	// slice of int
	mySlice = append(mySlice, myOtherSlice...)

	fmt.Println(mySlice)

	mySliceStr := []string{"Monday", "Tuesday"}
	myOtherSliceStr := []string{"Wednesday", "Thursday", "Friday"}

	mySliceStr = append(mySliceStr, myOtherSliceStr...)
	fmt.Println(mySliceStr)

	// delete
	mySliceStr = append(mySliceStr[:2], mySliceStr[3:]...)
	fmt.Println(mySliceStr)

	// shorthand slice
	student := []string{}
	students := [][]string{}
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil) // false

	// var slice
	var student1 []string
	var students1 [][]string
	fmt.Println(student1)
	fmt.Println(students1)
	fmt.Println(student1 == nil) // true

	// make slice
	student2 := make([]string, 35)
	students2 := make([][]string, 35)
	fmt.Println(student2)
	fmt.Println(students2)
	fmt.Println(student2 == nil)

	// compare shorthand slice
	student3 := []string{}
	students3 := [][]string{}
	//student3[0] = "Todd" // panic: runtime error: index out of range [0] with length 0
	student3 = append(student3, "Todd")
	fmt.Println(student3)
	fmt.Println(students3)

	// var slice
	var student4 []string
	var students4 [][]string
	//student4[0] = "Todd" // panic: runtime error: index out of range [0] with length 0
	student4 = append(student4, "TOdd")
	fmt.Println(student4)
	fmt.Println(students4)

	// make slice
	student5 := make([]string, 3)
	students5 := make([][]string, 3)
	student5[0] = "Todd"
	student5 = append(student5, "Todd")
	fmt.Println(student5)
	fmt.Println(students5)

	var records [][]string
	student6 := make([]string, 4)
	student6[0] = "Foster"
	student6[1] = "Nathan"
	student6[2] = "100.00"
	student6[3] = "74.00"
	records = append(records, student6)

	student7 := make([]string, 4)
	student7[0] = "Gomez"
	student7[1] = "Lisa"
	student7[2] = "92.00"
	student7[3] = "96.00"
	records = append(records, student7)
	fmt.Println(records)

	transactions := make([][]int, 0, 3)
	for i := 0; i < 3; i++ {
		transaction := make([]int, 0, 4)
		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)

	mySlice3 := make([]int, 1)
	fmt.Println(mySlice3[0])
	mySlice3[0] = 7
	fmt.Println(mySlice3[0])
	mySlice3[0]++
	fmt.Println(mySlice3[0])
}

func map_exercise() {
	// var map
	var myGreeting map[string]string
	fmt.Println(myGreeting)        // map[]
	fmt.Println(myGreeting == nil) // true     nil map
	//myGreeting["Tim"] = "Good morning" //!! panic: assignment to entry in nil map

	// make map
	myGreeting1 := make(map[string]string)
	// var myGreeting1 = make(map[string]string)
	fmt.Println(myGreeting1)        // map[
	fmt.Println(myGreeting1 == nil) // false
	myGreeting1["Tim"] = "Good morning."
	fmt.Println(myGreeting1)

	// composite map
	myGreeting2 := map[string]string{}
	fmt.Println(myGreeting2)
	fmt.Println(myGreeting2 == nil) // false: 这样的初始化都不是nil
	myGreeting2["Tim"] = "Good morning."
	fmt.Println(myGreeting2)

	// my Greet
	myGreeting3 := map[string]string{
		"Tim":   "Good morning",
		"Jenny": "Bonjour!",
	}
	fmt.Println(myGreeting3["Jenny"])

	// adding entry
	myGreeting4 := map[string]string{
		"Tim":   "Good morning",
		"Jenny": "Bonjour!",
	}
	myGreeting4["Harleen"] = "Howdy"

	fmt.Println(myGreeting4)
	fmt.Println(len(myGreeting4))

	// update
	myGreeting4["Harleen"] = "Gidday"
	fmt.Println(myGreeting4)

	// delete
	delete(myGreeting4, "Jenny")
	fmt.Println(myGreeting4["Harleen"])

	if val, exists := myGreeting4["Harleen"]; exists { // ?
		fmt.Println("That value exists.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	delete(myGreeting4, "haha")
	fmt.Println(myGreeting4)

	for key, val := range myGreeting4 {
		fmt.Println(key, " - ", val)
	}

	// runes
	letter := 'A'
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)

	letter1 := rune("A"[0])
	fmt.Println(letter1)

	word := "hello 你好"
	letter2 := rune(word[1])
	fmt.Println(letter2)

	/*
	* byte 数据类型与 rune相似，都是用来表示字符类型。
	* byte 等同于int8,常用来处理ascii字符
	* rune 等同于int32, 常用来处理unicode或utf8字符
	 */
	fmt.Println("len(word): ", len(word))                            // 12
	fmt.Println("rune: ", len([]rune(word)))                         // 8 (golang中string底层是通过byte数组实现的。中文字符在unicode下占2个字符，在utf-8编码下占3个字符，而golang默认是utf-8)
	fmt.Println("RuneCountInString: ", utf8.RuneCountInString(word)) // 8

	for i := 65; i < 122; i++ {
		fmt.Println(i, " - ", string(i), " - ", i%12) // 65 A 5
	}

	fmt.Println("int(A): ", int('A')) // 65

	// http.get
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	//res, err := http.Get("http://haha")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(res.Body)  // &{0xc0001181c0 {0 0} false <nil> 0x1204640 0x1204740}
	/*
	* ioutil.ReadAll(r io.Reader)
	* ioutil.ReadAll 从 r 中读取数据直达遇到EOF或error
	* io 包提供了对 I/O 原语的基本接口。
	 */
	// 方法一：读取
	bs, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", bs)

	// 方法二：读取
	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func scannerExercise() {
	/*
	* bufio 实现了有缓冲的I/O。
	* bufio.NewScanner(r io.Reader) *Scanner 创建并返回一个从r读取数据的Scanner，默认的分割函数是ScanLines
	* Scanner 提供方便的读取数据的接口，如从换行符分割的文本里读取每一行。（分割函数可以将文件分割为行、字节、unicode码值、空白分隔的word。调用者可以定制自己的分割函数）
	* Reader 类型通过从一个字符串读取数据，
	 */
	const input = "It is not the critic who counts;"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// set the split function for the scanning operation
	scanner.Split(bufio.ScanWords) // bufio.ScanWords 会将空白分割的片段作为一个token返回。
	// Count the words.
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input: ", err)
	}
}

func intSlicePlus() {
	buckets := make([]int, 1)
	fmt.Println(buckets[0])
	buckets[0] = 42
	fmt.Println(buckets[0])
	buckets[0]++
	fmt.Println(buckets[0])

	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)

	//buckets1 := make([]int, 200)
	//for scanner.Scan() {
	//	n := hashBucket(scanner.Text())
	//	buckets1[n]++
	//}
	//fmt.Println(buckets1[65:123])
	//fmt.Println("***************")
	//for i := 28; i < 126; i++ {
	//	// %c 该值对应的unicode码值
	//	// %v 值的默认格式表示
	//	fmt.Printf("%v - %c - %v \n", i, i, buckets1[i])
	//}

	// scan the page
	//buckets2 := make([]int, 12)
	//for scanner.Scan() {
	//	n := hashBucket2(scanner.Text(), 12)
	//	buckets2[n]++
	//}
	//fmt.Println(buckets2)

	// creates a slice with len and cap equal to 12. I can now access each of the twelve positions in the slice by index and assign values to them.
	buckets3 := make([][]string, 12)
	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket2(word, 12)
		// if I "append" to this slice, like this...
		buckets3[n] = append(buckets3[n], word)
	}

	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets3[i]))
	}
	fmt.Println(len(buckets3))
	fmt.Println(cap(buckets3))
}

func hashBucket(word string) int {
	return int(word[0])
}
func hashBucket1(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}
func hashBucket2(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
	// comment out the above, then uncomment the below
	// a more uneven distribution
	// return len(word) % buckets
}

func mapBucket() {
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)

	buckets := make(map[int]map[string]int)
	for i := 0; i < 12; i++ {
		buckets[i] = make(map[string]int)
	}
	for scanner.Scan() {
		word := scanner.Text()
		n := hashBucket2(word, 12)
		buckets[n][word]++
	}
	for k, v := range buckets[6] {
		fmt.Println(v, " \t- ", k)
	}
}

func englishAlphabet() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist.wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// 读取数据，方法一：ioutil.ReadAll(res.Body)
	//bs, _ := ioutil.ReadAll(res.Body)
	//str := string(bs)
	//fmt.Println(str)

	// 读取数据，方法二：bufio.NewScanner(res.Body)
	//scanner := bufio.NewScanner(res.Body)
	//defer res.Body.Close()
	//scanner.Split(bufio.ScanLines)
	//for scanner.Scan() {
	//	fmt.Println("-----")
	//	fmt.Println(scanner.Text())
	//}

	words := make(map[string]string)
	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = ""
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input: ", err)
	}

	i := 0
	for k := range words {
		fmt.Println(k)
		if i == 200 {
			break
		}
		i++
	}
}

/*
*
Go is Object Oriented
1. Encapsulation
2. Reusability
3. Polymorphism
4. Overriding

Traditional OOP
Classes
-- data structure describing a type of object
-- you can then create "instances"/"objects" from the class/blue-print
-- classes hold both:
=== state / data / fields
=== behavior / methods
-- public / private

Inheritance
- you don't create classes, you create a type
- you don't instantiate, you create a value of a type
*/
type foo int

type person struct {
	first string
	last  string
	age   int
}

type person1 struct {
	First       string
	Last        string
	Age         int
	notExported int
}

type person2 struct {
	First string
	Last  string `json:"-"`
	Age   int    `json:"wisdom score"`
}

type person3 struct {
	First string
	Last  string
	Age   int
}

func (p person) fullName() string {
	return p.first + p.last
}
func (p person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

type doubleZero struct {
	person
	LicenseToKill bool
	first         string
}

func (dz doubleZero) Greeting() {
	fmt.Println("Miss Moneypenny, so good to see you.")
}

func structDefined() {
	var myAge foo
	myAge = 44
	// %T 值的类型；%t 单词true或false
	fmt.Printf("%T %v %t\n", myAge, myAge, true)

	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())

	p3 := doubleZero{
		person:        person{"James", "Bond", 20},
		LicenseToKill: true,
		first:         "Double Zero Seven",
	}
	p4 := doubleZero{
		person: person{
			first: "Miss",
			last:  "MoneyPenny",
			age:   19,
		},
		first:         "Double",
		LicenseToKill: false,
	}
	fmt.Println(p3.first, p3.last, p3.age, p3.LicenseToKill)
	fmt.Println(p4.first, p4.last, p4.age, p4.LicenseToKill)
	p3.Greeting()
	p4.Greeting()
	p4.person.Greeting()

	// struct pointer
	p5 := &person{"James", "Bound", 20}
	fmt.Println(p5)
	fmt.Println(p5.first)
	fmt.Println(p5.age)

	// Exported
	p6 := person1{"James", "Bond", 20, 007}
	bs, _ := json.Marshal(p6)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))

	// unExported
	p7 := person{"James", "Bound", 20}
	fmt.Println(p7)
	bs1, _ := json.Marshal(p7)
	fmt.Println(string(bs1))

	p8 := person2{"James", "Bond", 20}
	bs2, _ := json.Marshal(p8)
	fmt.Println(string(bs2))

	var p9 person3
	fmt.Println(p9.First)
	fmt.Println(p9.Last)
	fmt.Println(p9.Age)

	bs3 := []byte(`{"First": "James", "Last": "Bond", "Age": 20}`)
	json.Unmarshal(bs3, &p9)

	fmt.Println("-----")
	fmt.Println(p9.First)
	fmt.Println(p9.Last)
	fmt.Println(p9.Age)
	fmt.Printf("%T \n", p9)

	var p10 person2
	fmt.Println(p10.First)
	fmt.Println(p10.Last)
	fmt.Println(p10.Age)

	bs4 := []byte(`{"First": "James", "Last": "Bond", "wisdom score": 20}`)
	json.Unmarshal(bs4, &p10)
	fmt.Println("-------------")
	fmt.Println(p10.First)
	fmt.Println(p10.Last)
	fmt.Println(p10.Age)
	fmt.Printf("%T \n", p10)

	/*
	* Decoder 从输入流解码json对象
	* json.NewDecoder(r io.Reader) *Decoder 创建一个从r读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据
	 */
	var p11 person1
	rdr := strings.NewReader(`{"First": "James", "Last": "Bond", "Age": 20}`)
	// Decode(v interface{}) 从输入流读取下个json编码并保存在v指向的值里
	json.NewDecoder(rdr).Decode(&p11)

	fmt.Println(p11.First)
	fmt.Println(p11.Last)
	fmt.Println(p11.Age)
}

type square struct {
	side float64
}
type circle struct {
	radius float64
}

func (z square) area() float64 {
	return z.side * z.side
}

// which implements the shape interface
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(z shape) { // interface 作为参数z
	fmt.Println(z)
	fmt.Println(z.area())
}
func totalArea(shapes ...shape) float64 { // 扩展运算符？
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
func interfaceExercise() {
	// no interface
	s := square{10}
	fmt.Println("Area: ", s.area())

	// interface
	s1 := square{20}
	fmt.Printf("%T\n", s1)
	info(s1)

	c1 := circle{5}
	info(c1)
	fmt.Println("Total Area: ", totalArea(c1, s1))

	msg := "Do not dwell in the past, do not dream of the future, concentrate the mind on the present."
	rdr := strings.NewReader(msg)
	io.Copy(os.Stdout, rdr)

	// bytes.NewBuffer(buf []byte) *Buffer
	// 使用buf作为初始内容创建并初始化一个Buffer，本函数用于读取已存在数据的buffer
	// buf应作为具有指定容量但长度为0的切片
	rdr2 := bytes.NewBuffer([]byte(msg))
	io.Copy(os.Stdout, rdr2)

	//res, _ := http.Get("http://www.mcleods.com")
	//io.Copy(os.Stdout, res.Body)
	//res.Body.Close()

	// error checking
	rdr3 := strings.NewReader(msg)
	_, err := io.Copy(os.Stdout, rdr3)
	if err != nil {
		fmt.Println(err)
		return
	}

	//res, err := http.Get("http://www.mcleods.com")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//io.Copy(os.Stdout, res.Body)
	//if err := res.Body.Close(); err != nil {
	//	fmt.Println(err)
	//}

	// package sort
}

type man []string

func (p man) Len() int { return len(p) }
func (p man) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p man) Less(i, j int) bool {
	return p[i] < p[j]
}
func sortPackage() {
	studyGroup := man{"Zeno", "John", "Al", "Jenny"}
	studyGroup1 := []string{"HaHa", "LiLy", "Al"}

	fmt.Println(studyGroup)
	// sort.Sort(d Interface)
	sort.Sort(studyGroup)
	// sort.Sort(studyGroup1) 报错：Cannot use 'studyGroup1' (type []string) as the type Interface Type does not implement 'Interface' as some methods are missing
	// sort.StringSlice 给 studyGroup1 添加方法(Len() Swap() Less())以满足 Interface接口，才能调用sort.Sort()
	sort.Sort(sort.StringSlice(studyGroup1))
	fmt.Println(studyGroup)
	fmt.Println(studyGroup1)

	s := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(s) // 将s排序为递增顺序
	fmt.Println(s)
	fmt.Printf("Just s: %T\n", s) // []string
	t := sort.StringSlice(s)
	fmt.Printf("just t: %T\n", t)                                    // sort.StringSlice
	fmt.Printf("s reverse: %T\n", sort.Reverse(sort.StringSlice(s))) // *sort.reverse
	//sort.Reverse(sort.StringSlice(s)) // 没有拍
	sort.Sort(sort.Reverse(sort.StringSlice(s))) // 所以倒序要在加层sort.Sort()
	fmt.Println(s)
	fmt.Println("---------")

	i := sort.Reverse(sort.StringSlice(s))
	fmt.Println(i)        // &{[Zeno John Jenny Al]}
	fmt.Printf("%T\n", i) // *sort.reverse
	sort.Sort(i)
	fmt.Println(s) // [Zeno John Jenny Al]

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println(n)
	sort.Sort(sort.IntSlice(n))
	fmt.Println(n)

	n1 := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println(n1)
	sort.Sort(sort.Reverse(sort.IntSlice(n1)))
	fmt.Println(n1)

	n2 := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(n2)
	fmt.Println(n2)
}

type men struct {
	Name string
	Age  int
}

func (m men) String() string {
	return fmt.Sprintf("YAYAYA %s: %d", m.Name, m.Age)
}

// ByAge implements sort.Interface for []person based on the Age Field
type ByAge []men

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

type animal struct {
	sound string
}
type dog struct {
	animal
	friendly bool
}
type cat struct {
	animal
	annoying bool
}

// 1. 利用空interface，传入任意struct
func specs(a interface{}) {
	fmt.Println(a)
}

func standardLibraryExample() {
	men := []men{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(men[0]) // 自定义String()覆盖
	fmt.Println(men)
	sort.Sort(ByAge(men))
	fmt.Println(men)

	fido := dog{animal{"woof"}, true}
	fifi := cat{animal{"meow"}, true}
	specs(fifi)
	specs(fido)

	shadow := dog{animal{"woof"}, true}
	// 2. 利用interface，组成任意类型的数组
	critters := []interface{}{fido, fifi, shadow}
	fmt.Println(critters)

	c := circle{4}
	info(&c) // &{4} 50.265482
	info(c)  // {4} 50.265482 如果func (c *circle) area() float64{} 参数是指针的话，这里必须info(&c)

	// conversion
	var x = 12
	var y = 12.1230123
	fmt.Println(y + float64(x)) // int to float64
	fmt.Println(int(y) + x)     // float64 to int

	var x1 rune = 'a' // rune is an alias for int32; normally omitted in this statement
	var y1 int32 = 'b'
	var z1 byte = 'c'
	fmt.Println(x1) // 97
	fmt.Println(y1) // 98
	fmt.Println(z1) // 99
	fmt.Println(string(x1))
	fmt.Println(string(y1)) // rune to string
	fmt.Println(string(z1))

	fmt.Println(string([]byte{'h', 'e', 'l', 'l', 'o'})) // []bytes to string
	fmt.Println(string([]rune{'w', 'o', 'r', 'l', 'd'}))
	fmt.Println(string([]int32{'t', 'h', 'i', 's'}))

	fmt.Println([]byte("hello")) // string to []bytes

	/*
	* strconv 基本数据类型和其字符串表示的相互转换
	* strconv.Atoi(s string) (i int, err error) 是ParseInt(s, 10, 0)的简写
	* strconv.Itoa(i int) string 是FormatInt(i, 10)的简写
	 */
	var x2 = "12"
	var y2 = 6
	z2, _ := strconv.Atoi(x2)
	fmt.Println(y2 + z2)

	x3 := 12
	y3 := "I have this many: " + strconv.Itoa(x3)
	fmt.Println(y3)

	// convert strings to value: ParseBool, ParseFloat, ParseInt, and ParseUnit
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	i, _ := strconv.ParseInt("-42", 10, 64)
	u, _ := strconv.ParseUint("42", 10, 64)
	fmt.Println(b, f, i, u)

	// convert values to string: FormatBool, FormatFloat, FormatInt, and FormatUnit
	w4 := strconv.FormatBool(true)
	// fmt 表示格式：b（二进制） e（十进制） E（十进制指数）
	// prec 控制精度（排除指数部分）：对 f, e, E 指小数点后几位；对 g G，指的是总的数字个数
	x4 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	y4 := strconv.FormatInt(-42, 16) // base指的是进制
	z4 := strconv.FormatUint(42, 16)

	fmt.Println(w4, x4, y4, z4)

	//name := "Sydney"
	//str, ok := name.(string) // invalid operation
	//if ok {
	//	fmt.Printf("%q\n", str)
	//} else {
	//	fmt.Printf("value is not a string\n")
	//}
	var name interface{} = "Sydney"
	str, ok := name.(string) // ? interface 转 string？
	if ok {
		fmt.Printf("%q, %T, %T \n", str, str, name) // "Sydney", string string
	} else {
		fmt.Printf("value is not a string\n")
	}

	var name1 interface{} = 7
	str1, ok := name1.(string)
	fmt.Printf("%q, %T, %T\n", str1, str1, name1) // "" string int
	fmt.Println(ok)                               // false
	//fmt.Println(name1 + 6) //? Invalid operation: name1 + 6 (mismatched types interface{} and untyped int)
	fmt.Println(name1.(int) + 6) // interface to int ?
	if ok {
		fmt.Printf("%q, %T, %T\n", str1, str1, name1)
	} else {
		fmt.Printf("value is not a string\n")
	}

	name2 := 7.24
	fmt.Printf("%T\n", name2)      // float64
	fmt.Printf("%T\n", int(name2)) // int

	var name3 interface{} = 7
	fmt.Println("%T\n", name3)
	//fmt.Printf("%T\n", int(name3)) // interface 转 int 用 name3.(int)
	fmt.Printf("%T\n", name3.(int))
}

/*
* go routines
* 1. sync 包提供了基本的同步：如互斥锁
*    sync.WaitGroup 用于等待一组线程的结束。父线程调用Add方法设定应等待的线程数。被等待的线程结束时应调用Done方法。
*					同时主线程可以调用Wait方法阻塞至所有线程结束。
 */
var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func goRoutines() {
	//foo1()
	//bar1()

	// concurrency
	//wg.Add(2)
	//go foo1()
	//go bar1()
	//wg.Wait()

	//wg.Add(2)
	//go incrementor("Foo: ")
	//go incrementor("Bar: ")
	//wg.Wait()
	//fmt.Println("Final Counter: ", counter)

	wg.Add(2)
	go incrementor1("Foo: ")
	go incrementor1("Bar: ")
	wg.Wait()
	fmt.Println("Final Counter: ", counter1)
}
func foo1() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}
func bar1() {
	for i := 0; i < 50; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(20 * time.Millisecond)
	}
	wg.Done()
}
func init() {
	/*
	* runtime包提供与go运行时环境的互操作，如控制go程的函数；也用于reflect包的低层次类型信息
	* 1。 环境变量 GOGC 设置最初的垃圾收集目标百分比。GOGC=off会完全关闭垃圾收集
	* 2。 GOMAXPROCS 限制可以同时运行用户层次的go代码的操作系统进行数
	* 3。 GOTRACEBACK 控制当go程序因为不能恢复的panic或不期望的运行时情况失败时的输出
	* 4。 GOARCH、GOOS、GOPATH、GOROOT：在编译时确认
	 */
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("----GOROOT: ", runtime.GOROOT(), runtime.NumCPU())
}

/*
* rand 用于更安全的加解密的更安全的随机数生成器
* sync.Mutex 是一个互斥锁：Mutex类型的锁与线程无关，可以由不同的线程加锁和解锁。
* 			mutex.Lock()
*           mutex.Unlock() 如果m未加锁会导致运行时错误
 */
func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(s, i, "counter: ", counter)
		mutex.Unlock()
	}
	wg.Done()
}

/*
*  sync/atomic 提供了底层的原子级内存操作，对于同步算法的实现很有用。
*				应通过通信来共享内存，而不是通过共享内存实现通信。
 */
var counter1 int64

func incrementor1(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		atomic.AddInt64(&counter1, 1)
		fmt.Println(s, i, "Counter:", atomic.LoadInt64(&counter1))
	}
	wg.Done()
}

/*
* channel
 */
func channelsExample() {
	// unbuffered channels
	//channel1()

	//channel2()

	//channel3()

	//channel4()
	//channel5()
	//channel6()
	//channel7()
	//channel8()
	//channel9()
	//channel10()
	//channel11()
	//channelClosures()
	//channelClosure2()
	//channelClosure3()
	//deadlock1()
	//deadlock1Solution()
	//deadlock2()
	//deadlock3()
	//deadlock3Solution()
	//factorial()
	//factorialMain2()
	//channel12()
	//channel13()
	//channel14()
	//channel15()
	//channel16()
	//channel17()
	//channel18()
	//channel19()
	//channel20()
	//channel22()
	//channel23()
	//channel24()
	//channel25()
	//channel26()
	//goRoutines5()
	//goRoutines6()
	forFun()
	//errorHandling()
	//errorHandling2()
	//errorHandling3()
	//errorHandling4()
	errorHandling5()
}
func channel1() {
	// unbuffered channels
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("c <-", i)
			c <- i
		}
	}()
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	time.Sleep(time.Second) // 1s
}
func channel2() {
	// unbuffered channel
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("c <- ", i)
			c <- i
		}
		close(c)
	}()
	// range 取代 <-c
	for n := range c {
		fmt.Println(n)
	}
}
func channel3() {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			fmt.Println("--1---c <- ", i)
			c <- i
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			fmt.Println("--2---c <- ", i)
			c <- i
		}
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(c) // panic: send on closed channel(如果没有wg.Add(2))
	}()
	for n := range c {
		fmt.Println(n)
	}
}
func channel4() {
	c := make(chan int)
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("1--c <- ", i)
			c <- i
		}
		fmt.Println("1--done <- true")
		done <- true
	}()
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("2--c <- ", i)
			c <- i
		}
		fmt.Println("2--done <- true")
		done <- true
	}()
	go func() {
		<-done
		<-done
		close(c) // 此方式能运行完并关掉
	}()
	for n := range c {
		fmt.Println(n)
	}
}
func channel5() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("1--done <- true")
			c <- i
		}
		done <- true
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()
	// we block here until done <- true
	<-done
	<-done
	close(c)

	// to unblock above
	// we need to take values off of chan c here, but we never get here, because we are blocked above
	for n := range c {
		fmt.Println(n)
	}
}
func channel6() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			done <- true
		}()
	}
	go func() {
		for i := 0; i < n; i++ {
			<-done // 10 个 done channel都取完了，才关闭
		}
		close(c)
	}()
	for n := range c { // 10 * 10
		fmt.Println(n)
	}
}
func channel7() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 100000; i++ {
			fmt.Println("c <-", i)
			c <- i
		}
		close(c)
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		fmt.Println("done <- true")
		done <- true
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		fmt.Println("done <- true")
		done <- true
	}()

	<-done
	<-done // 会等待所有go程执行完
}
func channel8() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 100000; i++ {
			fmt.Println("c <-", i)
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ {
		go func() {
			for n := range c {
				fmt.Println(n)
			}
			fmt.Println("done <- true")
			done <- true
		}()
	}

	for i := 0; i < n; i++ {
		fmt.Println("<-done")
		<-done
	}
}
func channel9() {
	c := incrementor2()
	cSum := puller(c)
	for n := range cSum {
		fmt.Println(n)
	}
}
func incrementor2() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("out <- ", i)
			out <- i
		}
		fmt.Println("incrementor: close out")
		close(out)
	}()
	return out
}
func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		fmt.Println("puller: close out")
		close(out) // channel close之后，依然可以读取？
	}()
	return out
}

// 结果和 channel9() 一摸一样。所以参数 <-chan 和 chan 一样么？
func channel10() {
	c := incrementor3()
	//cSum := puller3(c)
	//for n := range cSum {
	//	fmt.Println(n)
	//}
	for n := range puller3(c) {
		fmt.Println(n)
	}
}
func incrementor3() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("out <- ", i)
			out <- i
		}
		fmt.Println("incrementor: close out")
		close(out)
	}()
	return out
}
func puller3(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		fmt.Println("puller: close out")
		close(out)
	}()
	return out
}
func channel11() {
	c1 := incrementor4("Foo: ")
	c2 := incrementor4("Bar: ")
	c3 := puller4(c1)
	c4 := puller4(c2)
	fmt.Println("Final Counter: ", <-c3+<-c4)
}
func incrementor4(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			out <- 1
			fmt.Println(s, i)
		}
		close(out)
	}()
	return out
}
func puller4(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
func channelClosures() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func() {
			fmt.Println(v) // c c c
			done <- true
		}()
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
}
func channelClosure2() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func(u string) {
			fmt.Println(u) // a b c
			done <- true
		}(v)
	}

	for _ = range values {
		<-done
	}
}
func channelClosure3() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		v := v
		go func() {
			fmt.Println(v) // a b c
			done <- true
		}()
	}

	for _ = range values {
		<-done
	}
}
func deadlock1() {
	c := make(chan int)
	c <- 1
	fmt.Println("dead ") // 为什么这里会被锁？
	fmt.Println(<-c)
}
func deadlock1Solution() {
	c := make(chan int)
	go func() { // 在 go 程里就不会被锁？
		c <- 1
	}()
	fmt.Println(<-c)
}
func deadlock2() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()
	fmt.Println(<-c) // 为什么只打印个0？
}
func deadlock3() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	for {
		fmt.Println(<-c) // 打印里0-9，但是all goroutines are sleep-deadlock
	}
}
func deadlock3Solution() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // 1. remember to close your channel, or fatal error: all goroutines are asleep
	}()
	for n := range c { // 用range 而不是 for
		fmt.Println(n)
	}
}
func factorial() {
	f := factorial1(4)
	fmt.Println("Total: ", f) // 1 ??
}
func factorialMain2() {
	c := factorial2(4)
	for n := range c {
		fmt.Println(n) // 24 为什么放在go程里就可以呢
	}
}
func factorial1(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= 1
	}
	return total
}
func factorial2(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
func channel12() {
	c := gen(2, 3)
	out := sq(c)

	fmt.Println(<-out)
	fmt.Println(<-out)
}
func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
func channel13() {
	for n := range sq1(sq1(gen1(2, 3))) {
		fmt.Println(n) // 16 then 81
	}
}
func gen1(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
func sq1(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// go 程 和并发的问题
func channel14() {
	c := factorial3(4) // 24
	//c := factorial3(100) // 这里为啥是0 ??
	for n := range c {
		fmt.Println(n)
	}
}
func factorial3(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		fmt.Println("total: ", total)
		out <- total
		close(out)
	}()
	return out
}
func channel15() {
	in := gen15()
	f := factorial15(in)
	for n := range f {
		fmt.Println(n)
	}
}
func gen15() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}
func factorial15(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact15(n)
		}
		close(out)
	}()
	return out
}
func fact15(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

const numFactorials = 100
const rdLimit = 20

func channel16() {
	var w sync.WaitGroup
	w.Add(numFactorials)
	factorial16(&w)
	w.Wait()
}
func factorial16(wmain *sync.WaitGroup) {
	var w sync.WaitGroup
	rand.Seed(42)

	w.Add(numFactorials + 1)

	for j := 1; j <= numFactorials; j++ {
		go func() {
			f := rand.Intn(rdLimit)
			w.Wait()
			total := 1
			for i := f; i > 0; i-- {
				total *= i
			}
			fmt.Println(f, total)
			(*wmain).Done()
			//out <- total
		}()
		w.Done()
	}
	fmt.Println("All done with initialization")
	w.Done()
}

/*
* Channel fan in
* Channel fan out
 */
func channel17() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ { // 只读10次
		fmt.Println(<-c)
	}
	//for n := range c { // range c 会一直读取，
	//	fmt.Println(n)
	//}
	fmt.Println("You're both boring; I'm leaving.")
}
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for true {
			c <- <-input1
		}
	}()
	go func() {
		for true {
			c <- <-input2
		}
	}()
	return c
}
func channel18() {
	in := gen18(2, 3)

	// FAN OUT
	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq18(in)
	c2 := sq18(in)

	// FAN IN
	// Consume the merged output from multiple channels.
	for n := range merge18(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}
func gen18(nums ...int) chan int {
	fmt.Printf("TYPE OF NUMS %T\n", nums)

	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}
func sq18(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
func merge18(cs ...chan int) chan int {
	fmt.Printf("TYPE OF CS: %T\n", cs)
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
func channel19() {
	in := gen19(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq19(in)
	c2 := sq19(in)

	// Consume the merged output from c1 and c2
	for n := range merge19(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}
func gen19(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq19(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
func merge19(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// start an output goroutine for each input channel in cs.
	// output copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are done.
	// This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func channel20() {
	input := make(chan string)
	go workerProcess21(input)
	go workerProcess21(input)
	go workerProcess21(input)
	go publisher21(input)
	go publisher21(input)
	go publisher21(input)
	go publisher21(input)
	time.Sleep(1 * time.Millisecond)
}

var workerID int
var publisherID int

func publisher20(out chan string) {
	publisherID++
	thisID := publisherID
	dataID := 0
	for {
		dataID++
		fmt.Printf("publisher %d is pushing data\n", thisID)
		data := fmt.Sprintf("Data from publisher %d, Data %d", thisID, dataID)
		out <- data
	}
}
func workerProcess20(in <-chan string) {
	workerID++
	thisID := workerID
	for {
		fmt.Printf("%d: waiting for input...\n", thisID)
		input := <-in
		fmt.Printf("%d: input is: %s\n", thisID, input)
	}
}

var workerID21 int64
var publisherID21 int64

func publisher21(out chan string) {
	atomic.AddInt64(&publisherID21, 1)
	thisID := atomic.LoadInt64(&publisherID21)
	dataID := 0
	for {
		dataID++
		fmt.Printf("publisher %d is pushing data\n", thisID)
		data := fmt.Sprintf("Data from publish %d, Data %d", thisID, dataID)
		out <- data
	}
}
func workerProcess21(in <-chan string) {
	atomic.AddInt64(&workerID21, 1)
	thisID := atomic.LoadInt64(&workerID21)
	for {
		fmt.Printf("%d: waiting for input...\n", thisID)
		input := <-in
		fmt.Printf("%d: input is: %s\n", thisID, input)
	}
}
func channel22() {
	in := gen21()
	f := factorial21(in)
	for n := range f {
		fmt.Println(n)
	}
}
func gen21() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 0; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}
func factorial21(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact21(n)
		}
		close(out)
	}()
	return out
}
func fact21(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
func channel23() {
	in := gen23()
	c0 := factorial23(in)
	c1 := factorial23(in)
	c2 := factorial23(in)
	c3 := factorial23(in)
	c4 := factorial23(in)
	c5 := factorial23(in)
	c6 := factorial23(in)
	c7 := factorial23(in)
	c8 := factorial23(in)
	c9 := factorial23(in)

	var y int
	for n := range merge23(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		y++
		fmt.Println(y, "\t", n)
	}
}
func gen23() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}
func factorial23(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact21(n)
		}
		close(out)
	}()
	return out
}
func merge23(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
func fanOut23(in <-chan int, n int) []<-chan int {
	xc := make([]<-chan int, n)
	for i := 0; i < n; i++ {
		xc = append(xc, factorial23(in))
	}
	return xc
}
func channel24() { // 进程未正常结束
	in := gen23()
	xc := fanOut23(in, 10)
	for n := range merge23(xc...) { // 数组扩展运算符
		fmt.Println(n)
	}
}
func channel25() {
	in := gen23()
	xc := fanOut25(in, 10)
	fmt.Printf("%T \n", xc)
	fmt.Println("***********", len(xc))
	for _, v := range xc {
		fmt.Println("******", <-v)
	}
	for n := range merge23(xc...) {
		fmt.Println(n)
	}
}
func fanOut25(in <-chan int, n int) []<-chan int {
	var xc []<-chan int // this needed to be zero , why??
	for i := 0; i < n; i++ {
		xc = append(xc, factorial23(in))
	}
	return xc
}
func channel26() {
	in := gen26()
	out := make(chan int)

	fanOut26(in, 10, out)

	go func() {
		for v := range out {
			fmt.Println("----")
			fmt.Println(v)
		}
	}()
	var a string
	fmt.Println(&a)
}
func gen26() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}
func fanOut26(in <-chan int, n int, out chan<- int) {
	for i := 0; i < n; i++ {
		factorial25(in, out)
	}
}
func factorial25(in <-chan int, out chan<- int) {
	go func() {
		for n := range in {
			out <- fact21(n)
		}
	}()
}
func goRoutines5() {
	wg5.Add(2)
	go incrementor5("1")
	go incrementor5("2")
	wg5.Wait()
	fmt.Println("Final Counter: ", count)
}

var count int64
var wg5 sync.WaitGroup

func incrementor5(s string) {
	for i := 0; i < 20; i++ {
		atomic.AddInt64(&count, 1)
		fmt.Println("Process: "+s+"Printing:", i)
	}
	wg5.Done()
}
func goRoutines6() {
	c := incrementor6(2)
	var count int
	for n := range c {
		count++
		fmt.Println(n)
	}
	fmt.Println("Final Count: ", count)
}
func incrementor6(n int) chan string {
	c := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(i int) {
			for k := 0; k < 20; k++ {
				c <- fmt.Sprint("Process: "+strconv.Itoa(i)+" Printing: ", k)
			}
			done <- true
		}(i)
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()
	return c
}
func forFun() {
	m := map[int]int{}
	m[4] = 7
	m[3] = 87
	m[72] = 19

	ch := make(chan int)
	ch2 := make(chan int)
	go func() {
		var i int
		for n := range ch2 {
			i += n
			if i == len(m) {
				close(ch)
			}
		}
	}()
	go func() {
		for _, v := range m {
			ch <- v
			ch2 <- 1
		}
	}()
	for v := range ch {
		fmt.Println(v)
	}
	close(ch2)
}
func errorHandling() {
	x := 1
	str := evalInt(x)
	fmt.Println(str)
}
func evalInt(n int) string {
	if n > 10 {
		return fmt.Sprint("x is greater than 10")
	} else {
		return fmt.Sprint("x is less than 10")
	}
}
func errorHandling2() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened: ", err)
		log.Println("err happened: ", err)
		//log.Fatalln(err) // 因为 此相当于 {Println();os.Exit(1)}
		panic(err) // log.Fatalln() 后 则不执行此
	}
}
func errorHandling3() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
	}
}
func init7() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("=====init-create log.txt")
	log.SetOutput(nf) // 设置标准logger的输出目的地，默认是标准错误输出。
}

var ErrNegateMath = errors.New("negate math: square root of negative number")

func errorHandling4() {
	fmt.Printf("%T\n", ErrNegateMath)
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		//return 0, errors.New("negate math: square root of negative number")
		return 0, ErrNegateMath
	}
	return 42, nil
}

type NegateMathError struct {
	lat, long string
	err       error
}

func (n *NegateMathError) Error() string {
	return fmt.Sprintf("a negate math error occured: %v %v %v", n.lat, n.long, n.err)
}
func errorHandling5() {
	_, err := Sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}
func Sqrt5(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("negate math redux: square root of negative number: %v", f)
		return 0, &NegateMathError{"50.2289 N", "99.4656 W", nme}
	}
	return 42, nil
}
func Adder(xs ...int) int {
	res := 0
	for _, v := range xs {
		res += v
	}
	return res
}
func TestAdder(t *testing.T) {
	result := Adder(4, 7)
	if result != 11 {
		t.Fatal("4 + 7 did not equal 11")
	}
}
func isEqual(s string, g string) bool {
	if len(s) != len(g) {
		return false
	}
	if s == g {
		return true
	}
	i := 0
	for i < len(s) {
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
