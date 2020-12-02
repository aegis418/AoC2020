package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
	part2Other()
}

func processFile() []int {
	nums := make([]int, 0)
	file, _ := os.Open("input")
	defer file.Close()
	s := bufio.NewScanner(file)

	for s.Scan() {
		num, _ := strconv.Atoi(s.Text())
		nums = append(nums, num)
	}

	return nums
}

func part1() {
	nums := processFile()
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == 2020 {
				fmt.Println(nums[i] * nums[j])
				return
			}
		}
	}

}

// Fuck it we're doing it the easy way.
func part2() {
	nums := processFile()
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			for k := 0; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 2020 {
					fmt.Println(nums[i] * nums[j] * nums[k])
					return
				}
			}

		}
	}
}

// From reddit
func part2Other() {
	nums := processFile()
	for _, n := range nums {
		k := 2020 - n
		m := map[int]bool{}
		for _, e := range nums {
			m[e] = true
			if m[k-e] == true {
				fmt.Println(n * e * (k - e))
				return
			}
		}
	}

}
