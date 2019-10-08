package main

import (
	"fmt"
)

func getDisplayName(firstName, lastName *string) {
	defer func(){
		r := recover()
		fmt.Println("recovered from panic-", r)

	}()
	if firstName == nil {
		panic("runtime error: firstname must not be nil")
	}
	if lastName == nil {
		panic("runtime error: lastname must not be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("return from getDisplayName func")
}

// divide two number
func Divide(a int, b int) (int, error) {
	// can not divide by `0`
	if b == 0 {
		return 0, fmt.Errorf("Can not devide by Zero!")
	}
	return 0, nil
}

func main() {
	defer fmt.Println("deferred call in main func")
	firstName := "Calsoft"
	getDisplayName(&firstName, nil)
	// divide 5 by 0
	if _, err := Divide(5, 0); err != nil {
		fmt.Println("Error occured: ", err)
	}
	fmt.Println("return from main func")
}
