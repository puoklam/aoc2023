package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func ScanInput(name string) [][]rune {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)

	lines := make([][]rune, 0)
	for fs.Scan() {
		lines = append(lines, []rune(fs.Text()))
	}

	return lines
}

func validIndex(lines [][]rune, i, j int) bool {
	if i < 0 || i >= len(lines) {
		return false
	}
	if j < 0 || j >= len(lines[0]) {
		return false
	}
	return true
}

func connected(lines [][]rune, i, j int, dir rune) bool {
	r := lines[i][j]

	if dir == 'T' {
		if i-1 < 0 {
			return false
		}
		n := lines[i-1][j]
		if r == 'S' || r == '|' || r == 'J' || r == 'L' {
			if n == '|' || n == '7' || n == 'F' {
				return true
			}
		}
		return false
	}
	if dir == 'D' {
		if i+1 >= len(lines) {
			return false
		}
		n := lines[i+1][j]
		if r == 'S' || r == '|' || r == '7' || r == 'F' {
			if n == '|' || n == 'J' || n == 'L' {
				return true
			}
		}
		return false
	}

	if dir == 'L' {
		if j-1 < 0 {
			return false
		}
		n := lines[i][j-1]
		if r == 'S' || r == '-' || r == 'J' || r == '7' {
			if n == '-' || n == 'L' || n == 'F' {
				return true
			}
		}
		return false
	}

	if dir == 'R' {
		if j+1 < 0 {
			return false
		}
		n := lines[i][j+1]
		if r == 'S' || r == '-' || r == 'L' || r == 'F' {
			if n == '-' || n == 'J' || n == '7' {
				return true
			}
		}
		return false
	}
	return false
}

func Adjs(visited [][]bool, lines [][]rune, i, j int) [][2]int {
	adjs := make([][2]int, 0)

	if connected(lines, i, j, 'T') && !visited[i-1][j] {
		adjs = append(adjs, [2]int{i - 1, j})
	}
	if connected(lines, i, j, 'D') && !visited[i+1][j] {
		adjs = append(adjs, [2]int{i + 1, j})
	}
	if connected(lines, i, j, 'L') && !visited[i][j-1] {
		adjs = append(adjs, [2]int{i, j - 1})
	}
	if connected(lines, i, j, 'R') && !visited[i][j+1] {
		adjs = append(adjs, [2]int{i, j + 1})
	}
	return adjs
}

func FindDist(lines [][]rune) int {
	Max := 0
	m, n := len(lines), len(lines[0])

	var pq *PriorityQueue

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for i, line := range lines {
		for j, cell := range line {
			if cell == 'S' {
				pq = &PriorityQueue{&Item{
					i:    i,
					j:    j,
					dist: 0,
				}}
			}
		}
	}

	heap.Init(pq)

	for {
		item := heap.Pop(pq).(*Item)
		i, j := item.i, item.j
		if visited[i][j] {
			break
		}
		visited[i][j] = true

		if item.dist > Max {
			Max = item.dist
		}

		adjs := Adjs(visited, lines, i, j)
		for _, adj := range adjs {
			heap.Push(pq, &Item{
				i:    adj[0],
				j:    adj[1],
				dist: item.dist + 1,
			})
		}
	}

	return Max
}

func main() {
	fname := "./input.txt"
	lines := ScanInput(fname)
	Max := FindDist(lines)
	fmt.Println(Max)
}
