package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var i = 0
	var j = 0
	var res = 0
	for j < len(s) {
		if s[j] >= g[i] {
			res++
			i++
			j++
		} else {
			j++
		}
	}

	return res
}

func main() {
	var greed = []int{1, 2, 1, 3}
	var sizes = []int{1, 1, 2}
	res := findContentChildren(greed, sizes)
	fmt.Println(greed)
	fmt.Println(sizes)
	fmt.Printf("res : %d", res)
}
