package comma

import (
	"fmt"
	"testing"
)

type Test struct {
	input  string
	output string
}

var basicTests = []Test{
	{
		input:  "12345",
		output: "12,345",
	},
	{
		input:  "1234567",
		output: "1,234,567",
	},
	{
		input:  "12",
		output: "12",
	},
	{
		input:  "123",
		output: "123",
	},
}

func testComma(t *testing.T, fn func(string) string, tests []Test) {
	for _, test := range tests {
		name := fmt.Sprintf("intput:%s", test.input)
		t.Run(name, func(t *testing.T) {
			res := fn(test.input)
			if res != test.output {
				t.Errorf("got %s, expected: %s", res, test.output)
			}
		})
	}
}

func TestComma(t *testing.T) {
	testComma(t, Comma, basicTests)
}

func TestComma2(t *testing.T) {
	testComma(t, Comma2, basicTests)
}

func TestComma3(t *testing.T) {
	tests := append(basicTests, []Test{
		{
			input:  "-1",
			output: "-1",
		},
		{
			input:  "-123",
			output: "-123",
		},
		{
			input:  "-1.23",
			output: "-1.23",
		},
		{
			input:  "-123456.789",
			output: "-123,456.789",
		},
	}...)
	testComma(t, Comma3, tests)
}
