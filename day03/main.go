package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "unicode"
)

func ReadInput(filename string) []string {
	contents, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var lines []string
	sc := bufio.NewScanner(strings.NewReader(string(contents)))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines
}

func Part1(rucksacks []string) int {
	var totalPriority int
	
	for _, v := range rucksacks {
		var priority [53]int
		
		for i := 0; i < len(v) / 2; i++ {
			if v[i] < 97 {
				idx := int((v[i] % 65) + 27)
				priority[idx] = 1 
			} else{
				idx := int((v[i] % 97) + 1)
				priority[idx] = 1 
			}
		}

		for i := len(v) / 2; i < len(v); i++ {
			if v[i] < 97 {
				idx := int((v[i] % 65) + 27)
				if priority[idx] == 1 {
					totalPriority += idx
					break
				}
			} else{
				idx := int((v[i] % 97) + 1)
				if priority[idx] == 1 {
					totalPriority += idx
					break
				}
			}
		}
	}

	return totalPriority
}

func Part2(rucksacks []string) int {
	var totalPriority int
	var priority [53]int
	
	for i, v := range rucksacks {
		if i % 3 == 0 {
			priority = [53]int{}
		}

		for j := 0; j < len(v); j++ {
			if v[j] < 97 {
				idx := int((v[j] % 65) + 27)
				if i % 3 == 0 {
					priority[idx] = 1
				} else if i % 3 == 1 && priority[idx] == 1 {
					priority[idx] = 2 
				} else if i % 3 == 2 && priority[idx] == 2 {
					totalPriority += idx
					break
				}

			} else{
				idx := int((v[j] % 97) + 1)
				if i % 3 == 0 {
					priority[idx] = 1
				} else if i % 3 == 1 && priority[idx] == 1 {
					priority[idx] = 2 
				} else if i % 3 == 2 && priority[idx] == 2 {
					totalPriority += idx
					break
				}
			}
		}
	}

	return totalPriority
}

func main() {
	filename := "input.txt"

	contents := ReadInput(filename)

	ans1 := Part1(contents)
	fmt.Println("Total Priority 1:", ans1)

	ans2 := Part2(contents)
	fmt.Println("Total Priority 2:", ans2)
}
