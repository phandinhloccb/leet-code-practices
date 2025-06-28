package main

import (
	"fmt"
	"strconv"
	"strings"
)

func equalPairs(grid [][]int) int {
	rowCount := make(map[string]int)

	for _, row := range grid {
		var b strings.Builder
		for _, v := range row {
			b.WriteString(strconv.Itoa(v))
			b.WriteString(",")
		}
		key := b.String()
		rowCount[key]++
	}
	fmt.Println(rowCount)

	result := 0
	for i := 0; i < len(grid); i++ {
		var b strings.Builder
		for j := 0; j < len(grid); j++ {
			b.WriteString(strconv.Itoa(grid[j][i]))
			b.WriteString(",")
		}
		key := b.String()

		result += rowCount[key]

	}

	return result
}

func main() {
	grid := [][]int{
		{3, 2, 1},
		{1, 7, 6},
		{2, 7, 7},
	}

	fmt.Println(equalPairs(grid)) // Output: 1
}

/*
Given a 0-indexed n x n integer matrix grid, return the number of pairs (ri, cj) such that row ri and column cj are equal.

A row and column pair is considered equal if they contain the same elements in the same order (i.e., an equal array).
*/
