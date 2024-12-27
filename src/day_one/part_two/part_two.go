package dayoneparttwo

import (
	"advent_of_code_2024/src/internal/helpers"
	"fmt"
	"strconv"
	"strings"
)

func Solve() string {
	ch := make(chan string)
	go helpers.StreamInput("day_one/input.txt", ch)

	left := []int{}
	counts := map[int]int{}

	for line := range ch {
		fields := strings.Fields(line)

		value1, err := strconv.Atoi(fields[0])
		if err != nil {
			panic(fmt.Errorf("Error parsing value 1 from %s", line))
		}

		left = append(left, value1)

		value2, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(fmt.Errorf("Error parsing value 2 from %s", line))
		}

		counts[value2]++
	}

	total := 0

	for _, val := range left {
		total += val * counts[val]
	}

	return fmt.Sprintf("%d", total)
}
