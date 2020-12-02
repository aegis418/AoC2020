package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

func processFile() []int {
	nums := make([]int, 0)
	file, _ := os.Open("input")
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
				println(nums[i] * nums[j])
				return
			}
		}
	}

}

func part2() {
	nums := processFile()
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			for k := 0; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 2020 {
					println(nums[i] * nums[j] * nums[k])
					return
				}
			}

		}
	}
}
