package main

import "fmt"


func main() { 
  //初始化账单
	mybill := newBill("summer's bill")
  
	fmt.Println(mybill.format())
}
