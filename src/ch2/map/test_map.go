package main

import "fmt"

func main() {
	var myMap map[string]string
	if myMap == nil {
		fmt.Println("myMap1 空")
	}

	// 开辟空间
	myMap = make(map[string]string, 10)
	myMap["one"] = "haha"
	myMap["two"] = "lala"
	myMap["three"] = "xixi"

	if myMap != nil {
		fmt.Println(myMap)
	}

	myMap3 := map[string]string {
		"one": "php",
		"two": "haha",
	}

	fmt.Println(myMap3)

	myMap4 := make(map[string]string)
	myMap4["C"] = "beijing"
	myMap4["J"] = "tokyo"
	myMap4["U"] = "NewYork"

	for k, v := range myMap4 {
		fmt.Println(k, " = ", v)
	}

	delete(myMap4, "C")

	myMap4["J"] = "Haha"

	printMap(myMap4)
}

func printMap(cityMap map[string]string)  { // 传递的是一个引用
	cityMap["E"] = "hala"
	for k, v := range cityMap {
		fmt.Println("key = ", k)
		fmt.Println("value = ", v)
	}
}
