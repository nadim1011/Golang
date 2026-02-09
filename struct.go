package main

import "fmt"

type stu struct {
	Name string
	Roll int
}

func main() {
	s1 := stu{
		Name: "nadim",
		Roll: 120,
	}
	s2 := stu{
		Name: "Borad",
		Roll: 12,
	}

	fmt.Println(s1.Name)
	fmt.Println(s1.Roll)

	fmt.Println(s2.Name)
	fmt.Println(s2.Roll)
}
