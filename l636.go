package main

import (
	"fmt"
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	callstack := make([]int, 0, n)
	rst := make([]int, n)
	f, _, _ := parse(logs[0])
	callstack = push(callstack, f)
	for i := 1; i < len(logs); i++ {
		_, ls, lt := parse(logs[i-1])
		f, s, t := parse(logs[i])
		if s == "start" && ls == "end" {
			if len(callstack) != 0 {
				rst[back(callstack)] += t - lt - 1
			}
			callstack = push(callstack, f)
		} else if ls == "start" && s == "end" {
			if len(callstack) != 0 {
				rst[back(callstack)] += t - lt + 1
			}
			_, callstack = pop(callstack)
		} else {
			if len(callstack) != 0 {
				rst[back(callstack)] += t - lt
			}
			callstack = push(callstack, f)
		}
		fmt.Println(callstack)
	}
	return rst
}

func push(slice []int, item int) []int {
	return append(slice, item)
}

func pop(slice []int) (int, []int) {
	return slice[len(slice)-1], slice[:len(slice)-1]
}

func back(slice []int) int {
	return slice[len(slice)-1]
}

func parse(log string) (int, string, int) {
	segs := strings.Split(log, ":")
	function, _ := strconv.Atoi(segs[0])
	time, _ := strconv.Atoi(segs[2])
	return function, segs[1], time
}

func main() {
	fmt.Println(exclusiveTime(3, []string{"0:start:0", "0:end:0", "1:start:1", "1:end:1", "2:start:2", "2:end:2", "2:start:3", "2:end:3"}))
}
