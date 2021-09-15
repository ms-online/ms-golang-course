package main

import (
	"fmt"
	"math"
)


func sayGreeting(n string) {
	fmt.Printf("早上好，%v! \n", n)
}

func sayBye(n string) {
	fmt.Printf("再见，%v \n", n)
}

func cycleNames(n []string, f func(string)){
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return	math.Pi * r * r
}


func main() { 
	// len() append()

	sayGreeting("summer")
	sayBye("lucy")
	sayBye("tom")

	cycleNames([]string{"summer","lucy","henry"}, sayGreeting)
	cycleNames([]string{"summer","lucy","henry"}, sayBye)

	 a1 := circleArea(10)
	 a2 := circleArea(15.5)

	 fmt.Println(a1, a2)
	 fmt.Printf("a1的面积为：%0.3f, a2的面积为：%0.3f", a1,a2)
}

