package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInputs(filename string) []string {
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


func Part1(rounds []string) int{
	scoreMap := make(map[string]int)

	scoreMap["A X"] = 1 + 3
	scoreMap["A Y"] = 2 + 6
	scoreMap["A Z"] = 3 + 0
	scoreMap["B X"] = 1 + 0
	scoreMap["B Y"] = 2 + 3
	scoreMap["B Z"] = 3 + 6
	scoreMap["C X"] = 1 + 6
	scoreMap["C Y"] = 2 + 0
	scoreMap["C Z"] = 3 + 3

	var ans int

	for _, v := range rounds{
		ans += scoreMap[v]
	}

	return ans
}

func Part2(rounds []string) int {
	scoreMap := make(map[string]int)

	scoreMap["A X"] = 3 + 0
	scoreMap["A Y"] = 1 + 3
	scoreMap["A Z"] = 2 + 6
	scoreMap["B X"] = 1 + 0
	scoreMap["B Y"] = 2 + 3
	scoreMap["B Z"] = 3 + 6
	scoreMap["C X"] = 2 + 0
	scoreMap["C Y"] = 3 + 3
	scoreMap["C Z"] = 1 + 6

	var ans int

	for _, v := range rounds{
		ans += scoreMap[v]
	}

	return ans
}

func main() {
	filename := "input.txt"
	rounds := ReadInputs(filename)

	part1 := Part1(rounds)
	fmt.Println("Total score for Round1:", part1)

	part2 := Part2(rounds)
	fmt.Println("Total score for Round2:", part2)
}
