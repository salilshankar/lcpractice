package main

import (
	"fmt"
)

func main() {
	input := []int{0, 1, 0, 1, 1, 0, 0, 1}
	fmt.Println(input)
	fmt.Println()
	fmt.Println()

	fmt.Println("AAAA", prisonAfterNDays(input, 7))
}

func prisonAfterNDays(cells []int, N int) []int {

	m := make(map[int]int)
	ctr := 0

	cells = getNewCells(cells)

	for !found(m, cells) {
		cells = getNewCells(cells)
		ctr++
	}

	N = N % ctr
	if N == 0 {
		N = ctr
	}

	for i := 1; i < N; i++ {
		cells = getNewCells(cells)
	}

	return cells
}

func getNewCells(cells []int) []int {
	cache := make([]int, len(cells))
	for i := 1; i < len(cells)-1; i++ {
		if cells[i-1] == cells[i+1] {
			cache[i] = 1
		}
	}

	return cache
}

func found(a map[int]int, b []int) bool {
	x := turnToInt(b)
	if _, ok := a[x]; ok {
		return true
	} 
        
    a[x] = 0
	return false
}

func turnToInt(a []int) int {
	x := 0
	incrementer := 1
	for _, v := range a {
		x = x*incrementer + v
		incrementer *= 10
	}

	return x
}
