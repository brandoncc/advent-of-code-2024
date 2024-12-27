package main

import (
	dayonepartone "advent_of_code_2024/src/day_one/part_one"
	dayoneparttwo "advent_of_code_2024/src/day_one/part_two"
	daytwopartone "advent_of_code_2024/src/day_two/part_one"
	"fmt"
)

func main() {
	fmt.Printf("Day 1, part 1: %s\n", dayonepartone.Solve())
	fmt.Printf("Day 1, part 2: %s\n", dayoneparttwo.Solve())
	fmt.Printf("Day 2, part 1: %s\n", daytwopartone.Solve())
}
