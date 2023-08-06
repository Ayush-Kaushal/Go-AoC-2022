package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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

func createSupplyStack(stackContent []string) map[int][]string {
	supplyStack := map[int][]string{}

	for _, v := range stackContent {
		for i, c := range v {
			if i%4 == 1 && unicode.IsLetter(c) {
				supplyStack[i/4+1] = append([]string{string(c)}, supplyStack[i/4+1]...)
			}
		}
	}

	return supplyStack
}

func createStackOperations(operations []string) [][]int {
	stackOperations := [][]int{}

	for _, v := range operations {
		cmd := strings.Split(v, " ")

		quantity, err := strconv.Atoi(cmd[1])
		if err != nil {
			panic(err)
		}

		initialStack, err := strconv.Atoi(cmd[3])
		if err != nil {
			panic(err)
		}

		finalStack, err := strconv.Atoi(cmd[5])
		if err != nil {
			panic(err)
		}

		op := []int{quantity, initialStack, finalStack}
		stackOperations = append(stackOperations, op)
	}

	return stackOperations
}

func Part1(supplyStack map[int][]string, stackOperations [][]int) string {
	var topCrates string

	for i := range stackOperations {
		moves := stackOperations[i][0]
		initialStack := stackOperations[i][1]
		finalStack := stackOperations[i][2]

		for moves > 0 {
			supplyStack[finalStack] = append(supplyStack[finalStack], supplyStack[initialStack][len(supplyStack[initialStack])-1])
			supplyStack[initialStack] = supplyStack[initialStack][:len(supplyStack[initialStack])-1]

			moves--
		}
	}

	for i := 1; i <= len(supplyStack); i++ {
		topCrates += supplyStack[i][len(supplyStack[i])-1]
	}

	return topCrates
}

func Part2(supplyStack map[int][]string, stackOperations [][]int) string {
	var topCrates string

	for i := range stackOperations {
		moves := stackOperations[i][0]
		initialStack := stackOperations[i][1]
		finalStack := stackOperations[i][2]

		cratesToMove := len(supplyStack[initialStack]) - moves

		supplyStack[finalStack] = append(supplyStack[finalStack], supplyStack[initialStack][cratesToMove:]...)
		supplyStack[initialStack] = supplyStack[initialStack][:cratesToMove]

	}

	for i := 1; i <= len(supplyStack); i++ {
		topCrates += supplyStack[i][len(supplyStack[i])-1]
	}

	return topCrates
}

func main() {
	filename := "input.txt"
	contents := ReadInput(filename)

	si := 0
	for i, v := range contents {
		if v == "" {
			si = i
			break
		}
	}

	stackContent := contents[0:si]
	supplyStack := createSupplyStack(stackContent)

	operationContent := contents[si+1:]
	stackOperations := createStackOperations(operationContent)

	ans1 := Part1(supplyStack, stackOperations)
	fmt.Println("Top crates 9000:", ans1)

	supplyStack = createSupplyStack(stackContent)

	ans2 := Part2(supplyStack, stackOperations)
	fmt.Println("Top crates 9001:", ans2)
}
