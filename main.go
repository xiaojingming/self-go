package main

import "fmt"

func main() {
	var age int = 26
	fmt.Println("Age is:", age)

	var name = "Xiao"
	fmt.Println("Name is:", name)

	isDeveloper := true
	fmt.Println("Is developer:", isDeveloper)

	var count int
	fmt.Println("count:", count)

	var isFinished bool
	fmt.Println("isFinish:", isFinished)

	var str string
	fmt.Printf("str: %q\n", str)

	price := 100
	taxRate := 0.08
	total := float64(price) * taxRate
	println("total is", total)
}
