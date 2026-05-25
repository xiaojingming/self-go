package main

import "fmt"

func main() {
	fmt.Println("程序启动了")
	test1() // 在 main 里调用你的业务函数
	test2()
}

func test1() {
	score := 100
	if score > 50 {
		score = 200
		fmt.Println("score:", score)
	}
	fmt.Println("score:", score)
}

// 场景：你需要编写一段代码，计算一位用户购买商品后的账户余额。
// 用户初始余额（balance）为 100.50 元（浮点数）。
// 用户购买了一件商品，价格（price）为 30 元（整数）。
// 购买后，由于系统赠送，用户余额额外增加了 5 元（整数）。

func test2() {
	balance := 100.50
	price := 30
	gift := 5
	balance -= float64(price)
	balance += float64(gift)
	fmt.Println("balance:", balance)
}
