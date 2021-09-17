package main

import "fmt"

// 自定义数据类型
type bill struct {
	name string
	items map[string]float64
	tax float64
}

// 创建新账单
 func newBill(name string) bill{
	 b := bill{
		 name: name,
		 items: map[string]float64{},
		 tax: 0,
	 }
	 return b
 }

 // 格式化账单
 func (b bill) format() string{
   fs := "Bill Details: \n"
	 var total float64 = 0

	 // 遍历商品集合
    for k, v := range b.items{
			fs += fmt.Sprintf("%-25v ...$%v \n", k + ":", v)
			total += v
		}

		//税额
		fs += fmt.Sprintf("%-25v ...$%v \n","tax:", b.tax)

		//总金额
		fs += fmt.Sprintf("%-25v ...$%0.2f","total:", total + b.tax)

		return fs

 }

 // 更新税额
 func (b *bill) updateTax(tax float64) {
	  b.tax = tax
 }

 // 添加新商品到账单中
 func (b *bill) addItem(name string, price float64){
	 b.items[name] = price
 }