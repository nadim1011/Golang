package main

import "fmt"

func main() {
	var age int
	fmt.Scan(&age)
	if age == 17 {
		fmt.Println("hi")
	} else {
		fmt.Println("Who are you?")
	}
}
