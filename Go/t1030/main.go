package main

import (
	"container/list"
	"fmt"
)

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	res := [][]int{}
	vis := make([][]bool, R)
	for i := range vis {
		vis[i] = make([]bool, C)
	}
	q := list.New()
	q.PushBack([]int{r0, c0})
	vis[r0][c0] = true

	for q.Len() != 0 {
		cur := q.Front().Value.([]int)
		q.Remove(q.Front())
		x, y := cur[0], cur[1]
		res = append(res, cur)
		if x+1 < R && !vis[x+1][y] {
			q.PushBack([]int{x + 1, y})
			vis[x+1][y] = true
		}
		if x-1 >= 0 && !vis[x-1][y] {
			q.PushBack([]int{x - 1, y})
			vis[x-1][y] = true
		}
		if y-1 >= 0 && !vis[x][y-1] {
			q.PushBack([]int{x, y - 1})
			vis[x][y-1] = true
		}
		if y+1 < C && !vis[x][y+1] {
			q.PushBack([]int{x, y + 1})
			vis[x][y+1] = true
		}
	}
	return res
}

func main() {
	fmt.Println("Hello World")
}
