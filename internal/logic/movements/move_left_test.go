package movements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveLeft(t *testing.T) {
	casesToTest := []struct {
		description string
		values      [][]int
		expected    [][]int
	}{
		{
			description: "should merge equal values",
			values: [][]int{
				{2, 2, 0, 0},
				{4, 4, 0, 0},
				{8, 8, 0, 0},
				{16, 16, 0, 0},
			},
			expected: [][]int{
				{4, 0, 0, 0},
				{8, 0, 0, 0},
				{16, 0, 0, 0},
				{32, 0, 0, 0},
			},
		},
		{
			description: "should merge equal values and skip zeros",
			values: [][]int{
				{0, 2, 0, 2},
				{4, 0, 0, 4},
				{8, 0, 8, 0},
				{0, 16, 16, 0},
			},
			expected: [][]int{
				{4, 0, 0, 0},
				{8, 0, 0, 0},
				{16, 0, 0, 0},
				{32, 0, 0, 0},
			},
		},
		{
			description: "should merge pairs of equal values that are not side by side",
			values: [][]int{
				{0, 2, 2, 2},
				{4, 4, 0, 4},
				{8, 8, 8, 0},
				{8, 16, 8, 16},
			},
			expected: [][]int{
				{4, 2, 0, 0},
				{8, 4, 0, 0},
				{16, 8, 0, 0},
				{8, 16, 8, 16},
			},
		},
		{
			description: "should merge pairs of equal values that are not side by side",
			values: [][]int{
				{2, 2, 2, 2},
				{4, 4, 8, 4},
				{8, 16, 8, 8},
				{8, 16, 16, 16},
			},
			expected: [][]int{
				{4, 4, 0, 0},
				{8, 8, 4, 0},
				{8, 16, 16, 0},
				{8, 32, 16, 0},
			},
		},
		{
			description: "more cases",
			values: [][]int{
				{2, 2, 0, 2},
				{8, 0, 0, 0},
				{8, 4, 8, 8},
				{2, 16, 16, 128},
			},
			expected: [][]int{
				{4, 2, 0, 0},
				{8, 0, 0, 0},
				{8, 4, 16, 0},
				{2, 32, 128, 0},
			},
		},
	}

	for _, ct := range casesToTest {
		t.Run(ct.description, func(t *testing.T) {
			moveLeft(ct.values)

			assert.Equal(t, ct.expected, ct.values)
		})
	}
}
