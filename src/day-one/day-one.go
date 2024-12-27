package dayone

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solve1() string {
	ch := make(chan string)
	go loadInput("day-one/input1.txt", ch)

	list1 := []int{}
	list2 := []int{}

	for line := range ch {
		fields := strings.Fields(line)

		value1, err := strconv.Atoi(fields[0])
		if err != nil {
			panic(fmt.Errorf("Error parsing value 1 from %s", line))
		}

		list1 = append(list1, value1)

		value2, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(fmt.Errorf("Error parsing value 2 from %s", line))
		}

		list2 = append(list2, value2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	total := 0

	for i := 0; i < len(list1); i++ {
		total += int(math.Abs(float64(list1[i] - list2[i])))
	}

	return fmt.Sprintf("%d", total)
}

func Solve2() string {
	ch := make(chan string)
	go loadInput("day-one/input2.txt", ch)

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

func loadInput(inputPath string, ch chan string) {
	defer close(ch)

	file, err := os.Open(inputPath)
	if err != nil {
		panic(fmt.Errorf("Couldn't load inputs from %s, error: %w", inputPath, err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ch <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
