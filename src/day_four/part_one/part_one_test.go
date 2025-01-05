package dayfourpartone

import "testing"

type CountOccurrencesInputs struct {
	lines []string
	word  string
}

func TestCountOccurrences(t *testing.T) {
	tests := []struct {
		inputs   CountOccurrencesInputs
		expected int
	}{
		{
			inputs: CountOccurrencesInputs{
				lines: []string{""},
				word:  "XMAS",
			}, expected: 0,
		},
		{
			inputs: CountOccurrencesInputs{
				lines: []string{"XMA"},
				word:  "XMAS",
			}, expected: 0,
		},
		{
			inputs: CountOccurrencesInputs{
				lines: []string{"XMAS"},
				word:  "XMAS",
			}, expected: 1,
		},
		{
			inputs: CountOccurrencesInputs{
				lines: []string{"X", "M", "A", "S"},
				word:  "XMAS",
			}, expected: 1,
		},
		{
			inputs: CountOccurrencesInputs{
				lines: []string{"X   ", " M  ", "  A ", "   S"},
				word:  "XMAS",
			}, expected: 1,
		},
		{
			inputs: CountOccurrencesInputs{
				lines: []string{"X   ", "MM  ", "A A ", "S  S"},
				word:  "XMAS",
			}, expected: 2,
		},
		{
			inputs: CountOccurrencesInputs{
				lines: []string{"S", "A", "M", "X"},
				word:  "XMAS",
			}, expected: 1,
		},
		{
			inputs: CountOccurrencesInputs{
				lines: []string{"S   ", " A  ", "  M ", "   X"},
				word:  "XMAS",
			}, expected: 1,
		},
		{
			inputs: CountOccurrencesInputs{
				lines: []string{"S   ", "AA  ", "M M ", "X  X"},
				word:  "XMAS",
			}, expected: 2,
		},
		{
			inputs: CountOccurrencesInputs{
				lines: []string{
					" X  XX", "  M M ", " XMAS ", "  S S ",
				},
				word: "XMAS",
			}, expected: 3,
		},
	}

	for _, test := range tests {
		actual := countOccurrences(test.inputs.lines, test.inputs.word)

		if actual != test.expected {
			t.Errorf("countOccurrences(%v, %q) = %d; want %d", test.inputs.lines, test.inputs.word, actual, test.expected)
		}
	}
}
