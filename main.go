package main

import "fmt"


func main() { 
  //初始化账单
	mybill := newBill("summer's bill")
	
  mybill.addItem("apple",5.45)
  mybill.addItem("peach",2.56)
  mybill.addItem("orange",2.78)
  mybill.addItem("watermelon",1.88)
	mybill.updateTax(10)

	fmt.Println(mybill.format())
}
