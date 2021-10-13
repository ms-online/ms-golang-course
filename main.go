package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// 创建常量
const (
	ROCK = 0
	PAPER = 1
	SCISSORS = 2
)
func main() {
	// 设置随机数种子,保证每次随机都是随机的
	rand.Seed(time.Now().UnixNano())
  playerChoice := ""
	playerValue := -1
	// ***添加变量追踪双方得分
	playerScore := 0
	computerScore := 0



	// 读取标准输入
	reader := bufio.NewReader(os.Stdin)

	//清空屏幕
	clearScreen()

	// *** 输出游戏说明
	fmt.Println("欢迎来到“石头 - 剪刀 - 布”游戏")
	fmt.Println("----------------------")
	fmt.Println("游戏采取三局两胜制，当出现平局的时候，本局重新开始直到决出胜负，最后祝你好运！")
	fmt.Println()
	
	// ***添加循环
	for i := 1; i <= 3 ; i++{
	// ***打印当前回合信息
		fmt.Printf("第%d回合",i)
		fmt.Println()
		fmt.Println("------------")
	
  // 游戏提示:进入游戏
	fmt.Print("请输入rock,paper,scissors -> ")

	// 读取用户输入的选项
	playerChoice, _ = reader.ReadString('\n')
	// 清理用户输入
	playerChoice = strings.Replace(playerChoice,"\n", "", -1)

	// 保存用户选项所对应的值
	if playerChoice == "rock"{
		  playerValue = ROCK
	}else if playerChoice == "paper"{
			playerValue = PAPER
	}else if playerChoice == "scissors"{
			playerValue = SCISSORS
	}else {
		//*** 如果玩家输入无效则重置值为-1
		playerValue = -1
	}

	fmt.Println("玩家的选择",strings.ToUpper(playerChoice))
  // 电脑选项随机生成
		computerValue := rand.Intn(3)
	// 输出电脑随机生成的选项
	switch computerValue {
	  case ROCK :
		  fmt.Println("电脑的选择为ROCK")
		case PAPER :
			fmt.Println("电脑的选择为PAPER")
		case SCISSORS :
			fmt.Println("电脑的选择为SCISSORS")
	}

	fmt.Println()
	
 // 本回合结论
		if playerValue == computerValue {
			fmt.Println("本回合为平局！")
			// *** 将i进行递减，因为平局所以要重复当前回合
			i--
		}else {
			switch playerValue {
			case ROCK :
				if computerValue == PAPER {
				computerScore =	computerWins(computerScore)
				}else {
				playerScore =	playerWins(playerScore)
				}
			case PAPER :
				if computerValue == SCISSORS {
					computerScore =	computerWins(computerScore)
				}else {
					playerScore =	playerWins(playerScore)
				}
			case SCISSORS :
				if computerValue == ROCK {
					computerScore =	computerWins(computerScore)
				}else {
						playerScore =	playerWins(playerScore)
				}
			default:
				fmt.Println("无效输入，请重新选择")
				i--

			}
		}
	}
// ***最终结论
	fmt.Println("最终得分")
	fmt.Printf("玩家：%d/3， 电脑：%d/3", playerScore, computerScore)
	fmt.Println()
  
	if playerScore > computerScore {
		fmt.Println("最终玩家获胜")
	}else{
		fmt.Println("最终电脑获胜")
	}
}

// ***电脑获胜添加得分
func computerWins(score int) int{
	fmt.Println("本回合电脑获胜！")
	return score + 1
}
// ***玩家获胜添加得分
func playerWins(score int) int{
	fmt.Println("本回合玩家获胜！")
	return score + 1
}

// 清空屏幕
func clearScreen(){
	// runtime提供go运行时环境的操作
	//GOOS判断电脑的操作系统
	if strings.Contains(runtime.GOOS,"windows"){
	cmd := 	exec.Command("cmd","/c","cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	} else {
	cmd := 	exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	}
}