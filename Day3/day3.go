package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type slope struct {
	right int
	down  int
}

func main() {
	m := processFile()
	countTrees(m, 3, 1)
	countTreesMultipleSlopes(m)
}

func processFile() [][]string {
	file, _ := os.Open("input")
	defer file.Close()
	s := bufio.NewScanner(file)

	arr := make([][]string, 0)

	for s.Scan() {
		splitLine := strings.Split(s.Text(), "")
		arr = append(arr, splitLine)
	}

	return arr
}

func countTrees(m [][]string, right, down int) int {
	i, j := 0, 0
	numTrees := 0
	for i < len(m) {
		if m[i][j] == "#" {
			numTrees++
		}
		j += right
		if j >= len(m[i]) {
			j = j % len(m[i])
		}

		i += down
	}

	fmt.Println(numTrees)
	return numTrees
}

func countTreesMultipleSlopes(m [][]string) {
	slopes := [...]slope{slope{1, 1}, slope{3, 1}, slope{5, 1}, slope{7, 1}, slope{1, 2}}
	result := 1
	for _, slope := range slopes {
		result *= countTrees(m, slope.right, slope.down)
	}
	fmt.Println(result)
}
