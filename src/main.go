package main

import (
	dayonepartone "advent_of_code_2024/src/day_one/part_one"
	dayoneparttwo "advent_of_code_2024/src/day_one/part_two"
	daythreepartone "advent_of_code_2024/src/day_three/part_one"
	daytwopartone "advent_of_code_2024/src/day_two/part_one"
	daytwoparttwo "advent_of_code_2024/src/day_two/part_two"
	"fmt"
)

func main() {
	fmt.Printf("Day 1, part 1: %s\n", dayonepartone.Solve())
	fmt.Printf("Day 1, part 2: %s\n", dayoneparttwo.Solve())
	fmt.Printf("Day 2, part 1: %s\n", daytwopartone.Solve())
	fmt.Printf("Day 2, part 2: %s\n", daytwoparttwo.Solve())
	fmt.Printf("Day 3, part 1: %s\n", daythreepartone.Solve())
}
