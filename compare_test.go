package main

import (
	"testing"
)

func TestCompare(t *testing.T) {

	tests := []struct{
		A string
		B string
		Expected int
	}{
		{
			A: "log",
			B: "log",
			Expected: 0,
		},



		{
			A: "log/2015",
			B: "log/2015",
			Expected: 0,
		},



		{
			A: "log/2016",
			B: "log/2015",
			Expected: -1,
		},
		{
			A: "log/2015",
			B: "log/2016",
			Expected: 1,
		},



		{
			A: "log/2015/04",
			B: "log/2015/04",
			Expected: 0,
		},



		{
			A: "log/2015/04",
			B: "log/2015/01",
			Expected: -1,
		},
		{
			A: "log/2015/01",
			B: "log/2015/04",
			Expected: 1,
		},



		{
			A: "log/2015/04/11",
			B: "log/2015/04/11",
			Expected: 0,
		},



		{
			A: "log/2015/04/11",
			B: "log/2015/04/07",
			Expected: -1,
		},
		{
			A: "log/2015/04/07",
			B: "log/2015/04/11",
			Expected: 1,
		},



		{
			A: "log/2015/04/11/1234567",
			B: "log/2015/04/11/1234567",
			Expected: 0,
		},



		{
			A: "log/2015/04/11/9999999",
			B: "log/2015/04/11/1234567",
			Expected: -1,
		},
		{
			A: "log/2015/04/11/1234567",
			B: "log/2015/04/11/9999999",
			Expected: 1,
		},
	}

	for testNumber, test := range tests {

		actual := compare(test.A, test.B)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual result of the comparison is not what was expected.",  testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			t.Logf("A: %q", test.A)
			t.Logf("B: %q", test.B)
			continue
		}
	}
}
