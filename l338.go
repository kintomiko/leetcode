package main

import (
	"fmt"
)

// func countBits(num int) []int {
// 	sl := make([]int, num+1)
// 	sl[0] = 0
// 	sl[1] = 1
// 	if num == 0 {
// 		return []int{0}
// 	} else if num == 1 {
// 		return sl
// 	}
// 	for i := 2; i <= num; i++ {
// 		sl[i] = sl[i/2] + i%2
// 	}
// 	return sl
// }

func toString(arr []int) string {
	return fmt.Sprintf("%v", arr)
}

func main() {
	m := make(map[string]bool)
	m[toString([]int{1, 2, 3})] = true
	fmt.Println(toString([]int{1, 2, 3}))
	fmt.Println(m[toString([]int{1, 2, 3})])
	// fmt.Println(countBits(0))
}
