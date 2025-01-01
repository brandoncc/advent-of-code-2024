package daythreepartone

import "testing"

func TestParseMultiplicationCalls(t *testing.T) {
	tests := []struct {
		input    string
		expected []MultiplicationOperands
	}{
		{"", make([]MultiplicationOperands, 0)},
		{"mul(1,2)", []MultiplicationOperands{{left: 1, right: 2}}},
		{"mul(12,2)", []MultiplicationOperands{{left: 12, right: 2}}},
		{"mul(123,2)", []MultiplicationOperands{{left: 123, right: 2}}},
		{"mul(1234,2)", make([]MultiplicationOperands, 0)},
		{"mul(1,23)", []MultiplicationOperands{{left: 1, right: 23}}},
		{"mul(1,234)", []MultiplicationOperands{{left: 1, right: 234}}},
		{"mul(1,2345)", make([]MultiplicationOperands, 0)},
		{"mul(,2)", make([]MultiplicationOperands, 0)},
		{"mul(1,)", make([]MultiplicationOperands, 0)},
		{"mul(,)", make([]MultiplicationOperands, 0)},
		{"mul(a1,2)", make([]MultiplicationOperands, 0)},
		{"mul(1,b2)", make([]MultiplicationOperands, 0)},
		{"mul( 1 , 2 )", make([]MultiplicationOperands, 0)},
		{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", []MultiplicationOperands{
			{left: 2, right: 4},
			{left: 5, right: 5},
			{left: 11, right: 8},
			{left: 8, right: 5},
		}},
		{"{!how()'&where()don't()select()@]how()}mul(884,758);-mul(971,475)who()~from()]~mul(358,23)}", []MultiplicationOperands{
			{left: 884, right: 758},
			{left: 971, right: 475},
			{left: 358, right: 23},
		}},
	}

	for _, test := range tests {
		result := parseMultiplicationCalls(test.input)

		if len(result) != len(test.expected) {
			t.Fatalf("Wrong number of multiplication calls parsed from %s. Expected %d, got %d", test.input, len(test.expected), len(result))
		}

		for i, expected := range test.expected {
			actual := result[i]

			if expected.left != actual.left || expected.right != actual.right {
				t.Errorf("The wrong operands were parsed for mul() number %d from %s. Expected: %d, %d, got: %d, %d", i+1, test.input, expected.left, expected.right, actual.left, actual.right)
			}
		}
	}
}
