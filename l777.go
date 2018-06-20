package main

import (
	"fmt"
)

func canTransform(start string, end string) bool {
	if len([]rune(start)) != len([]rune(end)) {
		return false
	}
	var rls []rune

	for _, s := range start {
		if s != 'X' {
			rls = append(rls, s)
		}
	}

	k := 0
	for _, s := range end {
		if s != 'X' {
			if k >= len(rls) || s != rls[k] {
				return false
			}
			k += 1
		}
	}
	return true
}

func main() {
	if canTransform("X", "L") {
		fmt.Printf("true")
	}
}
