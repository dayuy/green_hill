package main

import (
	"fmt"
)

func main() {
	// 固定长度数组, 长度若是固定的，就不能再改变。（长度是数据类型的一部分 [10]int）
	var myArray1 [10]int // 0
	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{11, 22, 33, 44}
	// myArray3[4] = 3 会报错，长度固定

	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	for i, v := range myArray2 {
		fmt.Println("index = ", i, " , value = ", v)
	}

	// 查看数组的数据类型
	fmt.Printf("myArray1 types = %T\n", myArray1) // [10]int
	fmt.Printf("myArray2 types = %T\n", myArray2) // [10]int
	fmt.Printf("myArray3 types = %T\n", myArray3) // [4]int

	printArray(myArray3)

	fmt.Println("=====")

	// 动态数组，可以被改变里面的值  是一个slice，是一个切片
	// 因为动态数组，本身是一个数组的引用
	myArray4 := [4]int{1, 2, 3, 4}
	printArray(myArray4)
	for _, v := range myArray4 {
		fmt.Println("value ==== ", v)
	}
}

func printArray(myArray [4]int) { // 固定长度  值copy
	for i, v := range myArray {
		fmt.Println("index = ", i, ", value = ", v)
	}
	myArray[0] = 222 // 不会改变外面的myArray4,因为**数组是值传递**
}

func printArray2(myArray []int) { // 随机长度
	for _, value := range myArray {
		fmt.Println("value =+ ", value)
	}

	myArray[0] = 100 // 会改变外面的myArray4
}
