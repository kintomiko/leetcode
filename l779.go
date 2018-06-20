package main

import (
	"fmt"
	"math"
)

func kthGrammar(N int, K int) int {
	if N == 1 {
		return 0
	}
	preLen := int(math.Pow(2, float64(N-2)))
	fmt.Println(preLen)
	if K <= preLen {
		return kthGrammar(N-1, K)
	} else {
		return (kthGrammar(N-1, K-preLen) + 1) % 2
	}
}

func main() {
	fmt.Println(kthGrammar(2, 2))
}
