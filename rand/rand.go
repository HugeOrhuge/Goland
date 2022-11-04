package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func bubble(s *[]int) {
	l := len(*s)
	temp := 0
	for i := 0; i < l; i++ {
		for j := l - 1; j > i; j-- {
			if (*s)[j] < (*s)[j-1] {
				temp = (*s)[j]
				(*s)[j] = (*s)[j-1]
				(*s)[j-1] = temp
			}
		}
		fmt.Printf("冒泡排序后的%d个数是%d\n", i+1, (*s)[i])
	}
}

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	var s, p []int
	for i := 0; i < 100; i++ {
		s = append(s, rand.Intn(1000))
		p = append(p, s[i])
		fmt.Printf("第%d个随机数是%d\n", i+1, s[i])
	}

	fmt.Println()

	sort.Ints(s)
	for i := 0; i < 100; i++ {
		fmt.Printf("排序后的%d个数是%d\n", i+1, s[i])
	}

	fmt.Println()
	bubble(&p)
}
