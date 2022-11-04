package main

import "fmt"

func main() {
	var a, b int
	var ch string
	fmt.Println("请输入整数计算式：") //形式如1 + 2，2 * 1 + 3这种；中间要有空格
	fmt.Scan(&a, &ch, &b)
	switch ch {
	case "+":
		fmt.Println("计算结果为", a, ch, b, "=", a+b)
	case "-":
		fmt.Println("计算结果为", a, ch, b, "=", a-b)
	case "*":
		fmt.Println("计算结果为", a, ch, b, "=", a*b)
	case "/":
		fmt.Println("计算结果为", a, ch, b, "=", a/b)
	}
}
