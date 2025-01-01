package daythreeparttwo

import (
	"advent_of_code_2024/src/internal/helpers"
	"fmt"
	"strconv"
	"unicode"
)

type SeekMode int

const (
	Mul   SeekMode = iota
	Left  SeekMode = iota
	Right SeekMode = iota
)

const MaxDigits = 3

type MultiplicationOperands = struct {
	left  int
	right int
}

func Solve() string {
	ch := make(chan string)
	go helpers.StreamInput("day_three/input.txt", ch)

	total := 0
	instructionsEnabled := true

	for line := range ch {
		var calls []MultiplicationOperands
		calls, instructionsEnabled = parseMultiplicationCalls(line, instructionsEnabled)

		for _, call := range calls {
			total += call.left * call.right
		}
	}

	return fmt.Sprintf("%d", total)
}

func parseMultiplicationCalls(input string, instructionsEnabled bool) ([]MultiplicationOperands, bool) {
	var left int
	var right int
	var err error

	seekMode := Mul
	operandSets := []MultiplicationOperands{}

	pos := 0

	for pos < len(input) {
		switch seekMode {
		case Mul:
			if instructionsEnabled && len(input) >= pos+4 && input[pos:pos+4] == "mul(" {
				seekMode = Left
				pos += 4
			} else if !instructionsEnabled && len(input) >= pos+4 && input[pos:pos+4] == "do()" {
				instructionsEnabled = true
				pos += 4
			} else if instructionsEnabled && len(input) >= pos+7 && input[pos:pos+7] == "don't()" {
				instructionsEnabled = false
				pos += 7
			} else {
				pos += 1
			}
		case Left:
			digits := 0

			for digits < MaxDigits {
				if unicode.IsDigit(rune(input[pos+digits])) {
					digits++
				} else {
					break
				}
			}

			if digits == 0 || (len(input) > pos+digits && input[pos+digits] != ',') {
				seekMode = Mul
				continue
			}

			seekMode = Right
			left, err = strconv.Atoi(input[pos : pos+digits])
			if err != nil {
				panic(fmt.Sprintf("Unable to parse digits from %s", input[pos:pos+digits]))
			}

			// + 1 for the comma
			pos += digits + 1
		case Right:
			digits := 0

			for digits < MaxDigits {
				if unicode.IsDigit(rune(input[pos+digits])) {
					digits++
				} else {
					break
				}
			}

			if digits > 0 && len(input) > pos+digits && input[pos+digits] == ')' {
				right, err = strconv.Atoi(input[pos : pos+digits])
				if err != nil {
					panic(fmt.Sprintf("Unable to parse digits from %s", input[pos:pos+digits]))
				}

				// + 1 for the (
				pos += digits + 1
				operandSets = append(operandSets, MultiplicationOperands{left: left, right: right})
			}

			seekMode = Mul
		}
	}

	return operandSets, instructionsEnabled
}
