package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(filename string) []string{
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

func Part1(calories []int64) []int64{
	ans := []int64{}

	var elfNumber int64 = 1
	var maxCal int64 = 0
	var maxCalElf int64 = 1
	var sum int64 = 0

	for _, v := range calories {
		if v == -1 {
			if sum > maxCal {
				maxCal = sum
				maxCalElf = elfNumber
			}

			elfNumber++
			sum = 0
			continue
		}

		sum += v
	}

	ans = append(ans, maxCalElf, maxCal)

	return ans
}

func Part2(calories []int64) int64{
	var elfNumber int64 = 1
	var maxCal1, maxCal2, maxCal3 int64 = 0, 0, 0
	var sum int64 = 0

	for _, v := range calories {
		if v == -1 {
			if sum > maxCal1 {
				maxCal3 = maxCal2
				maxCal2 = maxCal1
				maxCal1 = sum
			} else if sum > maxCal2{
				maxCal3 = maxCal2
				maxCal2 = sum
			} else if sum > maxCal3{
				maxCal3 = sum
			}

			elfNumber++
			sum = 0
			continue
		}

		sum += v
	}

	return maxCal1 + maxCal2 + maxCal3
}

func main() {
	filename := "input.txt"

	input := ReadInput(filename)

	calories := []int64{}

	for _, v := range input{
		if v == "" {
			calories = append(calories, -1)
			continue
		}

		cal, err := strconv.ParseInt(v, 0, 64)
		if err != nil{
			panic(err)
		}

		calories = append(calories, cal)
	}

	part1 := Part1(calories)
	fmt.Println("Max Calories of 1 elf:", part1[1])

	part2 := Part2(calories)
	fmt.Println("Max calories of top 3 elves:", part2)
}
