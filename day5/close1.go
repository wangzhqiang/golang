package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{1, 3, 2, 6, 4, 9}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	fmt.Println(s)
}
