package daytwoparttwo

import (
	"advent_of_code_2024/src/internal/helpers"
	"fmt"
	"strconv"
	"strings"
)

type ReportChangeType int

const (
	Unknown = iota
	Increasing
	Decreasing
	None
)

func Solve() string {
	ch := make(chan string)
	go helpers.StreamInput("day_two/input.txt", ch)

	safeCount := 0

	for line := range ch {
		levels := levelsForReport(line)

		if isReportAcceptable(levels) {
			safeCount++
		}
	}

	return fmt.Sprintf("%d", safeCount)
}

func isReportAcceptable(levels []int) bool {
	return reportHasValidPath(levels, 0, 1, false) || reportHasValidPath(levels, 1, 2, true)
}

func reportHasValidPath(levels []int, leftIndex int, rightIndex int, levelAlreadySkipped bool) bool {
	initialLeftIndex := leftIndex
	left := levels[leftIndex]
	right := levels[rightIndex]
	changeType := identifyReportChangeType(left, right)

	for rightIndex < len(levels) {
		left = levels[leftIndex]
		right = levels[rightIndex]

		if !isLevelSafe(left, right, changeType) {
			if levelAlreadySkipped {
				return false
			}

			if rightIndex >= len(levels)-1 {
				return true
			}

			if leftIndex == initialLeftIndex {
				if isLevelSafe(left, levels[rightIndex+1], Unknown) {
					changeType = identifyReportChangeType(left, levels[rightIndex+1])
					rightIndex++
				} else {
					return false
				}
			} else {
				var changeTypeForBacktrack ReportChangeType

				if leftIndex == initialLeftIndex+1 {
					changeTypeForBacktrack = Unknown
				} else {
					changeTypeForBacktrack = changeType
				}

				if isLevelSafe(levels[leftIndex-1], levels[rightIndex], changeTypeForBacktrack) {
					changeType = identifyReportChangeType(levels[leftIndex-1], levels[rightIndex])
				} else if isLevelSafe(left, levels[rightIndex+1], changeType) {
					rightIndex++
				} else {
					return false
				}
			}

			levelAlreadySkipped = true
		}

		leftIndex = rightIndex
		rightIndex++
	}

	return true
}

func isLevelSafe(left int, right int, expectedChangeType ReportChangeType) bool {
	diff := left - right

	if diff == 0 {
		return false
	}

	if diff < -3 || diff > 3 {
		return false
	}

	if left < right && expectedChangeType == Decreasing {
		return false
	}

	if right < left && expectedChangeType == Increasing {
		return false
	}

	return true
}

func levelsForReport(report string) []int {
	fields := strings.Fields(report)
	fieldsAsInts := make([]int, len(fields))

	for i, val := range fields {
		fieldsAsInts[i] = stringToInt(val)
	}

	return fieldsAsInts
}

func identifyReportChangeType(left int, right int) ReportChangeType {
	if left < right {
		return Increasing
	}

	if left > right {
		return Decreasing
	}

	return None
}

func stringToInt(str string) int {
	converted, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("Couldn't parse int from string: %v", str))
	}

	return converted
}
