package main

import "fmt"

func main() {
	score := 85

	if score >= 60 {
		fmt.Println("成绩及格")
	} else {
		fmt.Println("成绩不及格")
	}

	for i := 1; i < 3; i++ {
		fmt.Println(i)
	}
}
