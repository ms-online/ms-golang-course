package main

import "fmt"

func main() { 
	// 数组：长度固定，类型固定
  //  var ages [3]int = [3]int{18,20,22}
	var ages = [3]int{18,20,22}

	names := [4]string{"summer","lucy","henry","lily"}
	names[1] = "John"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// 切片:动态数组
	var scores = []int{80,85,90,100}
	 scores[2] = 95
	scores = append(scores,98)

	fmt.Println(scores, len(scores))

	// 范围截取 - 包含起始下标,不包含终止下标
	rangeOne := names[1:3]
	rangeTwo := scores[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne,rangeTwo,rangeThree)

	rangeOne = append(rangeOne,"Tom")
	fmt.Println(rangeOne)


}