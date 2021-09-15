package main

import (
	"fmt"
	"strings"
)

func getFirstWords(n string) (string, string) {
    s := strings.ToUpper(n)
		names := strings.Split(s," ")

		var firstWords []string
		for _, v := range names {
			firstWords = append(firstWords, v[:1])
		}

		if len(firstWords) > 1 {
			return firstWords[0], firstWords[1]
		}

		return firstWords[0], "_"

}

func main() { 
 fn1, sn1 :=	getFirstWords("summer peng")
 fmt.Println(fn1,sn1)

 fn2, sn2 :=	getFirstWords("lucy deng")
 fmt.Println(fn2,sn2)

 fn3, sn3 :=	getFirstWords("henry")
 fmt.Println(fn3,sn3)
}

