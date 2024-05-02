package movements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveUp(t *testing.T) {
	casesToTest := []struct {
		description string
		values      [][]int
		expected    [][]int
	}{
		{
			description: "should merge equal values",
			values: [][]int{
				{2, 0, 2, 0},
				{2, 4, 2, 0},
				{0, 4, 0, 8},
				{0, 0, 0, 8},
			},
			expected: [][]int{
				{4, 8, 4, 16},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			description: "should merge equal values and skip zeros",
			values: [][]int{
				{0, 2, 0, 0},
				{4, 0, 0, 0},
				{0, 0, 8, 16},
				{4, 2, 8, 16},
			},
			expected: [][]int{
				{8, 4, 16, 32},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			description: "should merge pairs of equal values that are one on top of the other",
			values: [][]int{
				{0, 2, 2, 2},
				{4, 4, 2, 4},
				{8, 16, 8, 4},
				{8, 16, 8, 16},
			},
			expected: [][]int{
				{4, 2, 4, 2},
				{16, 4, 16, 8},
				{0, 32, 0, 16},
				{0, 0, 0, 0},
			},
		},
		{
			description: "should merge pairs of equal values that are one on top of the other",
			values: [][]int{
				{2, 4, 64, 32},
				{2, 4, 64, 32},
				{2, 16, 8, 64},
				{2, 16, 8, 64},
			},
			expected: [][]int{
				{4, 8, 128, 64},
				{4, 32, 16, 128},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
		{
			description: "more cases",
			values: [][]int{
				{2, 4, 8, 16},
				{4, 8, 16, 32},
				{8, 16, 32, 64},
				{16, 32, 64, 128},
			},
			expected: [][]int{
				{2, 4, 8, 16},
				{4, 8, 16, 32},
				{8, 16, 32, 64},
				{16, 32, 64, 128},
			},
		},
	}

	for _, ct := range casesToTest {
		t.Run(ct.description, func(t *testing.T) {
			moveUp(ct.values)

			assert.Equal(t, ct.expected, ct.values)
		})
	}
}
