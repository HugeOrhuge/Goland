package main

import "fmt"

type huge struct {
	num int
}

func swap(a *int, b *int) {
	var temp int
	temp = *b
	*b = *a
	*a = temp
}

func main() {
	var n int
	var size int
	var h [1000]huge
	var p1 int
	var p2 int
	fmt.Println("请输入数组大小和要翻转的长度：")
	fmt.Scanln(&n, &size)
	//fmt.Printf("%d%d\n", n, size)
	fmt.Println("请输入数组：")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &h[i].num)
		//if _, err := fmt.Scanf("%d", &h[i].num); err != nil {
		//	panic(err)
		//}
		//fmt.Printf("%d", h[i].num)
	}
	//fmt.Println()
	p1 = 0
	p2 = size - 1

	for p1 <= p2 {
		swap(&h[p1].num, &h[p2].num)
		p1++
		p2--
	}
	fmt.Println("翻转之后的数组为：")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", h[i].num)
	}
}
