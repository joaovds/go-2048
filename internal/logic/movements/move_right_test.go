package movements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveRight(t *testing.T) {
	casesToTest := []struct {
		description string
		values      [][]int
		expected    [][]int
	}{
		{
			description: "should merge equal values",
			values: [][]int{
				{0, 0, 2, 2},
				{0, 0, 4, 4},
				{0, 0, 8, 8},
				{0, 0, 16, 16},
			},
			expected: [][]int{
				{0, 0, 0, 4},
				{0, 0, 0, 8},
				{0, 0, 0, 16},
				{0, 0, 0, 32},
			},
		},
	}

	for _, ct := range casesToTest {
		t.Run(ct.description, func(t *testing.T) {
			moveRight(ct.values)

			assert.Equal(t, ct.expected, ct.values)
		})
	}
}
