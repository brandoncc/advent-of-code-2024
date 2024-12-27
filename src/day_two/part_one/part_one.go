package daytwopartone

import (
	"advent_of_code_2024/src/internal/helpers"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type ReportChangeType int

const (
	Unknown    ReportChangeType = iota
	Increasing ReportChangeType = iota
	Decreasing
)

func Solve() string {
	ch := make(chan string)
	go helpers.StreamInput("day_two/input.txt", ch)

	safeReports := 0

	for line := range ch {
		if isReportSafe(line) {
			safeReports++
		}
	}

	return fmt.Sprintf("%d", safeReports)
}

func isReportSafe(report string) bool {
	fields := strings.Fields(report)
	fieldsAsNumbers := make([]int, len(fields))

	for i, val := range fields {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			panic(fmt.Sprintf("Can't parse string to integer: %s", val))
		}

		fieldsAsNumbers[i] = intVal
	}

	reportChangeType := Unknown

	for i := 0; i < len(fields)-1; i++ {
		difference := fieldsAsNumbers[i] - fieldsAsNumbers[i+1]
		absoluteDifference := math.Abs(float64(difference))

		if absoluteDifference < 1 || absoluteDifference > 3 {
			return false
		}

		var currentChangeType ReportChangeType

		if difference > 0 {
			currentChangeType = Increasing
		} else {
			currentChangeType = Decreasing
		}

		if reportChangeType == Unknown {
			reportChangeType = currentChangeType
		}

		if currentChangeType != reportChangeType {
			return false
		}
	}

	return true
}
