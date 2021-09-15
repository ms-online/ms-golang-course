package main

import "fmt"

func main() { 
	// age := 55
 // 布尔值
	// fmt.Println( age < 40)
	// fmt.Println( age <= 40)
	// fmt.Println( age > 40)
	// fmt.Println( age >= 40)
	// fmt.Println( age == 35)
	// fmt.Println( age != 35)


	// 条件语句
	// if age < 40{
	// 	fmt.Println("年龄在40周岁以下")
	// } else if age < 50 {
	// 	fmt.Println("年龄在50周岁以下")
	// } else {
	// 	fmt.Println("年龄在55周岁及其以上")
	// }
	
 	names := []string{"summer","lucy","henry","tom","lily"}

	 for index, value := range names {
    if index == 1 {
			fmt.Println("当前下标为：", index)
			continue
		}

		if index > 2 {
			fmt.Println("下标值满足大于2的条件，此时索引为：", index)
			break
		}

		 fmt.Printf("下标 %v 的值为 %v \n", index, value)
	 }
}