package main

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (p IntSlice) Len() int {
	return len(p)
}
func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	arr := IntSlice{5, 1, 1, 2, 0, 0}
	fmt.Println("排序前的数组为：", arr)
	sort.Sort(arr)
	fmt.Println("排序后的数组为：", arr)
}
