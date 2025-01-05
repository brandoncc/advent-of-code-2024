package dayfourpartone

import (
	"advent_of_code_2024/src/internal/helpers"
	"fmt"
)

type Movement struct {
	x int
	y int
}

var movements = []Movement{
	{x: 1, y: 0},
	{x: 1, y: -1},
	{x: 0, y: -1},
	{x: -1, y: -1},
	{x: -1, y: 0},
	{x: -1, y: 1},
	{x: 0, y: 1},
	{x: 1, y: 1},
}

func Solve() string {
	ch := make(chan string)
	go helpers.StreamInput("day_four/input.txt", ch)

	lines := []string{}
	needle := "XMAS"

	for line := range ch {
		lines = append(lines, line)
	}

	occurrences := countOccurrences(lines, needle)

	return fmt.Sprintf("%d", occurrences)
}

func countOccurrences(lines []string, word string) int {
	firstCharInWord := word[0]
	occurrences := 0

	for y, line := range lines {
		for x, char := range line {
			if char == rune(firstCharInWord) {
				occurrences += countOccurrencesAtLocation(lines, word, x, y)
			}
		}
	}

	return occurrences
}

func countOccurrencesAtLocation(lines []string, word string, x int, y int) int {
	colCount := len(lines[0])
	rowCount := len(lines)

	occurrences := 0

	for _, movement := range movements {
		currX := x
		currY := y
		movesMade := 0

		for movesMade < len(word) {
			currX += movement.x
			currY += movement.y
			movesMade++

			if currX < 0 || currX >= colCount || currY < 0 || currY >= rowCount {
				break
			}

			if word[movesMade] != lines[currY][currX] {
				break
			}

			if movesMade == len(word)-1 {
				occurrences++
				break
			}
		}
	}

	return occurrences
}
