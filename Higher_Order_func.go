package main

import (
	"fmt"
)

func call() func(a, b int) {
	return cat
}
func cat(a, b int) {
	r := a + b
	fmt.Println(r)
}
func main() {
	a, b := 10, 20
	x := call()
	x(a, b)
	// fmt.Println(x)

}

func init() {
	fmt.Println("I am itit!!!")
}
