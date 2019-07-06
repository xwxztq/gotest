package main

import "fmt"

type stringer interface{
	string()string
}

type tester interface{
	stringer               // 嵌入其他接口
	test()
}

type data struct{}

func(*data)test() {
	fmt.Println("121213")
}

func(data)string()string{
	return""
}

func main() {
	var d data

	var t tester= &d
	t.test()
	fmt.Println(t.string())
	fmt.Println("123")
}
