package commands

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetId(t *testing.T) {

	testTable := []struct {
		numbers  string
		expected int64
	}{
		{
			numbers:  "542",
			expected: 542,
		},
		{
			numbers:  "-145",
			expected: -145,
		},
		{
			numbers:  "0",
			expected: 0,
		},
	}

	for _, testCase := range testTable {
		result := GetID(testCase.numbers)

		t.Logf("Calling GetId(%v), result %d\n", testCase.numbers, result)

		assert.Equal(t, testCase.expected, result, fmt.Sprintf("Incorrect result. Expect %d, got %d", testCase.expected, result))
	}
}
