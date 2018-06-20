package kin

import "golang.org/x/tour/tree"

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	_walk(t, ch)
	close(ch)
}

func _walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v, ok := <-ch1; ; {
		if !ok {
			break
		}
		if v != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
}
