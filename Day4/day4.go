package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
)

func main() {
	passports := processFile()
	fmt.Println(len(passports))
	fmt.Printf("Part 1: %d\n", validatePassports(passports))
	fmt.Printf("Part 2: %d\n", validatePassportsPt2(passports))
	//for _, p := range passports {
	//	fmt.Println(p)
	//}
	//testRegex()
}

func processFile() []string {
	file, _ := ioutil.ReadFile("input")
	passports := strings.Split(string(file), "\n\n")
	for i := range passports {
		passports[i] = strings.Replace(passports[i], "\n", " ", -1)
	}
	return passports
}

func getPassportInfo(passport string) []string {
	re := regexp.MustCompile(`([\w\d:#]+)`)
	fields := re.FindAllString(passport, -1)

	return fields
}

func validatePassports(passports []string) int {
	valid := 0
	re := regexp2.MustCompile(`(?=.*ecl)(?=.*pid)(?=.*eyr)(?=.*hcl)(?=.*byr)(?=.*iyr)(?=.*hgt)`, 2)
	for _, passport := range passports {
		if res, _ := re.MatchString(passport); res {
			valid++
		}
	}

	return valid
}

// For some reason we are off by 10 here.
func validatePassportsPt2(passports []string) int {
	valid := 0
	for _, passport := range passports {
		fields := getPassportInfo(passport)
		isValid := true
		fmt.Println(passport, len(fields))
		for _, f := range fields {
			key, val := f[:3], f[4:]
			switch key {
			case "byr":
				re := regexp.MustCompile(`^\d{4}$`)
				s := re.FindString(val)
				year, _ := strconv.Atoi(s)
				if !(year >= 1920 && year <= 2002) {
					fmt.Println("\tinvalid", key, val)
					isValid = false
					break
				}
			case "iyr":
				re := regexp.MustCompile(`^\d{4}$`)
				s := re.FindString(val)
				year, _ := strconv.Atoi(s)
				if !(year >= 2010 && year <= 2020) {
					fmt.Println("\tinvalid", key, val)
					isValid = false
					break
				}
			case "eyr":
				re := regexp.MustCompile(`^\d{4}$`)
				s := re.FindString(val)
				year, _ := strconv.Atoi(s)
				if !(year >= 2020 && year <= 2030) {
					fmt.Println("\tinvalid", key, val)
					isValid = false
					break
				}
			case "hgt":
				re := regexp.MustCompile(`^(\d+)(in|cm)$`)
				height := re.FindStringSubmatch(val)
				if len(height) != 0 {
					n, _ := strconv.Atoi(height[1])
					if height[2] == "cm" {
						if !(n >= 150 && n <= 193) {
							fmt.Println("\tinvalid", key, val)
							isValid = false
							break
						}
					} else if height[2] == "in" {
						if !(n >= 59 && n <= 76) {
							fmt.Println("\tinvalid", key, val)
							isValid = false
							break
						}
					} else {
						isValid = false
						break
					}
				} else {
					isValid = false
					break
				}
			case "hcl":
				re := regexp.MustCompile(`^#[0-9a-f]{6}$`)
				if !re.MatchString(val) {
					fmt.Println("\tinvalid", key, val)
					isValid = false
					break
				}
			case "ecl":
				re := regexp.MustCompile(`amb|blu|brn|gry|grn|hzl|oth`)
				if !re.MatchString(val) {
					fmt.Println("\tinvalid", key, val)
					isValid = false
					break
				}
			case "pid":
				re := regexp.MustCompile(`^\d{9}$`)
				if !re.MatchString(val) {
					fmt.Println("\tinvalid", key, val)
					isValid = false
					break
				}
			default:
				break
			}
		}
		if isValid && (len(fields) == 7 || len(fields) == 8) {
			valid++
		}
	}
	return valid
}
