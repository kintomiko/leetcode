package main

import "fmt"

func str(arr []int) string {
	return fmt.Sprintf("%v", arr)
}

var m = make(map[string]int)

func removeBoxes(boxes []int) int {
	return _removeBoxes(append(boxes, 9999))
}

func _removeBoxes(boxes []int) int {
	if v, ok := m[str(boxes)]; ok {
		return v
	}

	if len(boxes) == 2 {
		return 1
	}
	if len(boxes) == 1 {
		return 0
	}
	last := boxes[0]
	curStart := 0
	max := 0
	for i := 0; i < len(boxes); i++ {
		if boxes[i] != last {
			candidate := make([]int, curStart)
			copy(candidate, boxes[:curStart])
			candidate = append(candidate, boxes[i:]...)
			fmt.Println(candidate)
			candidateLen := i - curStart
			candidateCount := _removeBoxes(candidate)
			if candidateCount+candidateLen*candidateLen > max {
				max = candidateCount + candidateLen*candidateLen
			}
			last = boxes[i]
			curStart = i
		}
	}
	m[str(boxes)] = max
	return max
}

func main() {
	fmt.Println(removeBoxes([]int{1, 3, 2, 2, 2, 3, 4, 3, 1}))
}
