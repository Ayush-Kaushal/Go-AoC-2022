package main

import (
	"fmt"
	"os"
)

func ReadInput(filename string) string{
	contents, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(contents)
}

func Part1(dataBuffer string) int {
	freq := map[string]int{}

	var i, j = -1, -1

	for i < len(dataBuffer) - 1 {
		i++
		charFront := string(dataBuffer[i])

		length := i - j

		count, isPresent := freq[charFront]
		if isPresent {
			freq[charFront] = count + 1
		} else {
			freq[charFront] = 1
		}

		if length == 4 {
			if len(freq) == 4 {
				return i + 1
			}

			j++
			charBack := string(dataBuffer[j])

			if freq[charBack] == 1 {
				delete(freq, charBack)
			} else {
				freq[charBack] -= 1
			}
		}
	}

	return i
}

func Part2(dataBuffer string) int {
	freq := map[string]int{}

	var i, j = -1, -1

	for i < len(dataBuffer) - 1 {
		i++
		charFront := string(dataBuffer[i])

		length := i - j

		count, isPresent := freq[charFront]
		if isPresent {
			freq[charFront] = count + 1
		} else {
			freq[charFront] = 1
		}

		if length == 14 {
			if len(freq) == 14 {
				return i + 1
			}

			j++
			charBack := string(dataBuffer[j])

			if freq[charBack] == 1 {
				delete(freq, charBack)
			} else {
				freq[charBack] -= 1
			}
		}
	}

	return i
}

func main() {
	filename := "input.txt"
	dataBuffer := ReadInput(filename)

	packetMarker := Part1(dataBuffer)
	fmt.Println("No. of characters before start-of-packet marker:", packetMarker)

	mssgMarker := Part2(dataBuffer)
	fmt.Println("No. of characters before start-of-message marker:", mssgMarker)
}
