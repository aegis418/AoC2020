package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := processFile()
	//fmt.Println(input)
	fmt.Println("Part 1: ", findHighestSeat(input))
	fmt.Println("Part 2: ", findMySeat(input))
}

func processFile() []string {
	f, _ := ioutil.ReadFile("input")
	passes := strings.Split(string(f), "\n")
	for i, _ := range passes {
		for _, l := range []string{"F", "B", "L", "R"} {
			switch l {
			case "F":
				passes[i] = strings.Replace(passes[i], l, "0", -1)
			case "B":
				passes[i] = strings.Replace(passes[i], l, "1", -1)
			case "L":
				passes[i] = strings.Replace(passes[i], l, "0", -1)
			case "R":
				passes[i] = strings.Replace(passes[i], l, "1", -1)
			}
		}
	}
	return passes
}

func findHighestSeat(passes []string) int {
	max := 0
	for _, pass := range passes {
		id, _ := strconv.ParseInt(pass, 2, 0)
		if int(id) > max {
			max = int(id)
		}
	}
	return max
}

func findMySeat(passes []string) int {
	ids := make([]int, 0)
	for _, pass := range passes {
		id, _ := strconv.ParseInt(pass, 2, 0)
		ids = append(ids, int(id))
	}

	sort.Ints(ids)
	i, j := 0, 1
	for j < len(ids)-1 {
		if ids[j]-ids[i] == 2 {
			return ids[i] + 1
		}
		i++
		j++
	}
	return 0
}
