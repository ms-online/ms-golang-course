package main

import (
	"fmt"
	"sort"
	// "strings"
)

func main() { 
	// 标准库链接：https://pkg.go.dev/std
	// greeting := "Hello world My name is Summer"
  
	// strings 包
	// fmt.Println(strings.Contains(greeting, "world"))
	// fmt.Println(strings.ReplaceAll(greeting,"Hello","Hi"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "o"))
	// fmt.Println(strings.Split(greeting, " "))

	 // 原始变量不会被改变
	//  fmt.Println("原始变量不会改变：",greeting)


 // sort 包
   ages := []int{38,26,15,9,49,60,78,90}

	 sort.Ints(ages)
	 fmt.Println(ages)

	 indexOne := sort.SearchInts(ages, 15)
	 indexTwo := sort.SearchInts(ages, 50)
	 fmt.Println(indexOne)
	 fmt.Println(indexTwo)


	 names := []string{"summer","lucy","henry","lily","tom"}

	 sort.Strings(names)
	 fmt.Println(names)

	 fmt.Println(sort.SearchStrings(names, "henry"))

}