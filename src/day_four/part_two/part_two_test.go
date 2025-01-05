package dayfourparttwo

import "testing"

func TestCountXMasOccurrences(t *testing.T) {
  tests := []struct {
    inputs   []string
    expected int
  }{
    {
      inputs: []string{ "M M", " A ", "S S" },
      expected: 1,
    },
    {
      inputs: []string{
        "MMMSXXMASM",
        "MSAMXMSMSA",
        "AMXSXMAAMM",
        "MSAMASMSMX",
        "XMASAMXAMM",
        "XXAMMXXAMA",
        "SMSMSASXSS",
        "SAXAMASAAA",
        "MAMMMXMMMM",
        "MXMXAXMASX",
      },
      expected: 9,
    },
  }


  for _, test := range tests {
    actual := countXMasOccurrences(test.inputs)

    if actual != test.expected {
      t.Errorf("countOccurrences(%v) = %d; want %d", test.inputs, actual, test.expected)
    }
  }
}
