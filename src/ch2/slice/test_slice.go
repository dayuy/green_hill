package main

import "fmt"

func main() {
	// 声明slice1是一个切片，并且初始化，默认值是1，2，3。长度len是3
	var slice1 = []int{1,2,3,6,7}

	fmt.Println("len = %d, slice = %v\n", len(slice1), slice1)
	slice1[0] = 1

	var slice2 []int
	fmt.Println("slice2 = %d, %v\n", len(slice2), slice2)

	slice2 = make([]int, 3)  // 使用make开辟空间，初始化值为0
	slice2[0] = 100
	fmt.Println("slice2 = %d, %v\n", len(slice2), slice2)

	var slice3 []int
	if slice3 == nil {
		fmt.Println("空切片", slice3)
	}

	// 追加元素
	var nums = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(nums), cap(nums), nums)
	nums = append(nums, 1) // cap为一次性开辟多少空间
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(nums), cap(nums), nums)
	fmt.Println(slice1[:2], slice1[2:], slice1[1:3])
	slice4 := slice1[0:2] // 值的引用
	slice4[0] = 200
	fmt.Println(slice4, slice1)

	// copy 深拷贝
	slice5 := make([]int, 5)
	copy(slice5, slice1)
	slice5[0] = 777
	fmt.Println("len = %d, cap = %d, slice5 = %v, slice1 = %v\n", len(slice5), cap(slice5), slice5, slice1)

}
