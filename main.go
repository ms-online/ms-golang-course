package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 获取输入值的方法
func getInput(prompt string ,r *bufio.Reader) (string,error){
  fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	// 阅读器
	reader := bufio.NewReader(os.Stdin)

	//输出询问
	// fmt.Print("创建账单名称：")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("创建账单名称：",reader)

	// 创建新账单
	 b := newBill(name)
	 fmt.Println("账单已创建：", b.name)

	 return b

}

//可选项
func promptOptions(b bill){
// 阅读器
	reader := bufio.NewReader(os.Stdin)

	opt,_ := getInput("可选项（a - 添加商品， s - 保存账单， t - 添加税额）",reader)
	
	// switch语句
	switch opt {
		case "a":
			name, _ := getInput("商品名称：", reader)
			price, _ := getInput("商品价格：", reader) 

			p, err := strconv.ParseFloat(price, 64)

			//nil是一个预先声明的标识符，表示指针、通道、函数、接口、映射或切片类型。
			// 当出现不等于nil的时候,说明出现某些错误了,需要我们对这个错误进行一些处理,而如果等于nil说明运行正常。
			if err != nil {
				fmt.Println("金额必须为数字...")
				promptOptions(b)
			}
			// 添加商品
			b.addItem(name,p)

			fmt.Println("商品已经添加：",name, price)
			promptOptions(b)
		case "t":
			tax,_ := getInput("输入商品税额：", reader)

				t, err := strconv.ParseFloat(tax, 64)
			if err != nil {
				fmt.Println("金额必须为数字...")
				promptOptions(b)
			}
			b.updateTax(t)
			fmt.Println("税额已经更新：",t)
			promptOptions(b)
		case "s":
			fmt.Println("你的账单已保存", b)
		default:
				fmt.Println("选项无效...")
				promptOptions(b)
		}
}

func main() { 
  mybill := createBill()
  promptOptions(mybill)

}
