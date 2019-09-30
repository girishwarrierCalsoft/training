package main

import (
	"fmt"
	"strconv"
	"strings"
)


func main() {
	fmt.Println("Forward Pattern")
	for i:= 0; i<=4; i++{
		for k := 0; k <= i; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	fmt.Println("Reverse Pattern")
	for i:= 1; i<=4; i++{
		newString := strings.Repeat("*", i)
		revPattern := fmt.Sprintf("%4s\n", newString)
		fmt.Print(revPattern)
	}
	fmt.Println("Forward Pattern & Reverse Pattern")
	for i := 0; i < 4; i++ {
		for j := 0; j < 4-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	for i:= 4; i>=1; i--{
		str := strings.Repeat("*", i)
		pattern := fmt.Sprintf("%4s%s\n", str, str)
		fmt.Print(pattern)
	}
	fmt.Println("Fibonocci series")
	a := 0
	b := 1
	fmt.Println(0)
	for i:=1; i<=10; i++{
		c := a+b
		b = a
		a = c
		fmt.Println(c)
	}
	fmt.Println("Reverse a string")
	str:="12345"
	var result string
	for _,v := range str {
		result =  string(v) + result
	}
	fmt.Println(result)
	fmt.Println("Forward Number Pattern")
	var number string
	for i:=1; i<=5; i++{
		s := strconv.Itoa(i)
		number = number + s
		fmt.Println(number)
	}
	fmt.Println("Reverse Number Pattern")
	var number1 string
	for i:=1; i<=5; i++{
		s := strconv.Itoa(i)
		number1 = number1 + s
		fmt.Printf("%5s\n", number1)
	}

}

