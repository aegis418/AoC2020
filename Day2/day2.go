package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(countPasswordsPart1(processFile()))
	fmt.Println(countPasswordsPart2(processFile()))
}

func processFile() (hi, low []int, letters, passwords []string) {
	file, _ := os.Open("input")
	defer file.Close()
	s := bufio.NewScanner(file)

	hi = make([]int, 0)
	low = make([]int, 0)
	letters = make([]string, 0)
	passwords = make([]string, 0)

	for s.Scan() {
		parts := strings.Split(s.Text(), " ")

		nums := strings.Split(parts[0], "-")
		h, _ := strconv.Atoi(nums[1])
		l, _ := strconv.Atoi(nums[0])
		hi = append(hi, h)
		low = append(low, l)

		s := strings.Split(parts[1], "")[0]
		letters = append(letters, s)

		p := parts[2]
		passwords = append(passwords, p)
	}

	return
}

func countPasswordsPart1(hi, low []int, letters, passwords []string) int {
	validPasswords := 0
	for i, password := range passwords {
		regex := regexp.MustCompile(fmt.Sprintf(`[%s]`, letters[i]))
		res := regex.FindAllString(password, -1)

		count := len(res)

		if count >= low[i] && count <= hi[i] {
			validPasswords++
		}

	}

	return validPasswords
}

func countPasswordsPart2(hiPos, lowPos []int, letters, passwords []string) int {
	validPasswords := 0
	for i, password := range passwords {
		//fmt.Println(password, lowPos[i], hiPos[i])
		pos1Char := string(password[lowPos[i]-1])
		pos2Char := string(password[hiPos[i]-1])

		if letters[i] == pos1Char && letters[i] == pos2Char {
			continue
		} else if letters[i] == pos1Char || letters[i] == pos2Char {
			validPasswords++
		}
	}

	return validPasswords
}
