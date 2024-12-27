package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func StreamInput(inputPath string, ch chan string) {
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
