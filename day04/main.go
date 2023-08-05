package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func Part1(intervals []string) int {
	var count int = 0

	for _, v := range intervals {
		startVal1, err := strconv.ParseInt(strings.Split(strings.Split(v, ",")[0], "-")[0], 0, 64)
		if err != nil {
			panic(err)
		}
		endVal1, err := strconv.ParseInt(strings.Split(strings.Split(v, ",")[0], "-")[1], 0, 64)
		if err != nil {
			panic(err)
		}


		startVal2, err := strconv.ParseInt(strings.Split(strings.Split(v, ",")[1], "-")[0], 0, 64)
		if err != nil {
			panic(err)
		}
		endVal2, err := strconv.ParseInt(strings.Split(strings.Split(v, ",")[1], "-")[1], 0, 64)
		if err != nil {
			panic(err)
		}

		if startVal1 >= startVal2 && endVal1 <= endVal2 {
			count++
		} else if startVal2 >= startVal1 && endVal2 <= endVal1 {
			count++
		}
	}

	return count
}

func Part2(intervals []string) int {
	var count int = 0

	for _, v := range intervals {
		startVal1, err := strconv.ParseInt(strings.Split(strings.Split(v, ",")[0], "-")[0], 0, 64)
		if err != nil {
			panic(err)
		}
		endVal1, err := strconv.ParseInt(strings.Split(strings.Split(v, ",")[0], "-")[1], 0, 64)
		if err != nil {
			panic(err)
		}


		startVal2, err := strconv.ParseInt(strings.Split(strings.Split(v, ",")[1], "-")[0], 0, 64)
		if err != nil {
			panic(err)
		}
		endVal2, err := strconv.ParseInt(strings.Split(strings.Split(v, ",")[1], "-")[1], 0, 64)
		if err != nil {
			panic(err)
		}

		if startVal1 <= endVal2 && endVal1 >= startVal2 {
			count++
		} else if startVal2 <= endVal1 && endVal2 >= startVal1 {
			count++
		}
	}

	return count
}

func main() {
	filename := "input.txt"
	contents := ReadInput(filename)

	ans1 := Part1(contents)
	fmt.Println("No. of completely overlapping intervals:", ans1)

	ans2 := Part2(contents)
	fmt.Println("No. of overlapping intervals:", ans2)
}
