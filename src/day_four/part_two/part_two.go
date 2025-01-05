package dayfourparttwo

import (
	"advent_of_code_2024/src/internal/helpers"
	"fmt"
)

func Solve() string {
	ch := make(chan string)
	go helpers.StreamInput("day_four/input.txt", ch)

	lines := []string{}

	for line := range ch {
		lines = append(lines, line)
	}

	occurrences := countXMasOccurrences(lines)

	return fmt.Sprintf("%d", occurrences)
}

func countXMasOccurrences(lines []string) int {
	rowCount := len(lines)
	colCount := len(lines[0])

	occurrences := 0

	for x := 1; x <= colCount-2; x++ {
		for y := 1; y <= rowCount-2; y++ {
			if lines[y][x] == 'A' && xMasExistsAtLocation(lines, x, y) {
				occurrences++
			}
		}
	}

	return occurrences
}

func xMasExistsAtLocation(lines []string, x int, y int) bool {
	backslash :=
		lines[y+1][x-1] == 'M' && lines[y-1][x+1] == 'S' ||
			lines[y+1][x-1] == 'S' && lines[y-1][x+1] == 'M'

	forwardslash :=
		lines[y-1][x-1] == 'M' && lines[y+1][x+1] == 'S' ||
			lines[y-1][x-1] == 'S' && lines[y+1][x+1] == 'M'

	return backslash && forwardslash
}
