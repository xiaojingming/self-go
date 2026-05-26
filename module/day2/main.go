package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func queryDatabase() error {
	return errors.New("connection timeout")
}

func test1() {
	retries := 3
	for i := 0; i < retries; i++ {
		err := queryDatabase()
		if err == nil {
			fmt.Println("数据查询成功！")
			break
		}
		fmt.Printf("第%d次重试, 1s后重试\n", i)
		time.Sleep(time.Second)
	}
}

func handleRequest(method string, path string) {
	switch {
	case method == "GET" && path == "/users":
		fmt.Println("获取用户列表")
	case method == "POST" && path == "/users":
		fmt.Println("创建新用户")
	default:
		fmt.Println("404")
	}
}

func test2() {
	nums := []int{10, 20, 30}
	var ptrs []*int

	for i := range nums {
		ptrs = append(ptrs, &nums[i])
	}

	for _, val := range ptrs {
		fmt.Println(*val)
	}
}

func test3() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func FizzBuzz() {
	for i := 1; i <= 20; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

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

	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 3:
		fmt.Println("中指")
	default:
		fmt.Println("其他手指")
	}

	count := 0
	for count < 3 {
		fmt.Println("count:", count)
		count++
	}

	for {
		fmt.Println("无限循环")
		break
	}

	// num,err作用域仅在if-else块中
	if num, err := strconv.Atoi("123"); err == nil {
		fmt.Println("转换成功，num是：", num)
	} else {
		fmt.Println("转换失败：", err)
	}

	test1()
	handleRequest("GET", "/users")
	test2()
	fmt.Println("===============")
	test3()
	fmt.Println("===============")
	FizzBuzz()
}
