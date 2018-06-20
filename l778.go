package main

import (
	"fmt"
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
}

func (p pos) key() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func swimInWater(grid [][]int) int {
	covers := make(map[string]bool)
	adj := make(map[string]bool)

	mark(pos{x: 0, y: 0}, covers, grid, true)
	mark(pos{x: 0, y: 1}, adj, grid, true)
	mark(pos{x: 1, y: 0}, adj, grid, true)

	max := 0
	for {
		next := min(adj, grid)
		if grid[next.x][next.y] > max {
			max = grid[next.x][next.y]
		}
		fmt.Println(next)
		if next.x == len(grid[0])-1 && next.y == len(grid)-1 {
			return max
		}
		mark(next, covers, grid, true)
		delete(adj, next.key())

		markAdj(pos{x: next.x, y: next.y + 1}, adj, grid, true, covers)
		markAdj(pos{x: next.x, y: next.y - 1}, adj, grid, true, covers)
		markAdj(pos{x: next.x + 1, y: next.y}, adj, grid, true, covers)
		markAdj(pos{x: next.x - 1, y: next.y}, adj, grid, true, covers)
	}
}

func atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func parse(k string) pos {
	cor := strings.Split(k, ",")
	return pos{atoi(cor[0]), atoi(cor[1])}
}

func min(poses map[string]bool, grid [][]int) pos {

	m := 99999
	var minPos pos
	for k, _ := range poses {
		p := parse(k)
		if grid[p.x][p.y] < m {
			m = grid[p.x][p.y]
			minPos = p
		}
	}
	return minPos
}

func mark(p pos, marker map[string]bool, grid [][]int, val bool) {
	if p.x >= 0 && p.y >= 0 && p.x < len(grid[0]) && p.y < len(grid) {
		marker[p.key()] = val
	}
}

func markAdj(p pos, adj map[string]bool, grid [][]int, val bool, covers map[string]bool) {
	if _, ok := covers[p.key()]; !ok {
		mark(p, adj, grid, val)
	}
}

func main() {
	g := [][]int{{0, 2}, {1, 3}}

	fmt.Println(swimInWater(g))
}
