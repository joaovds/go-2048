package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsGameOver(t *testing.T) {
	t.Run("should return false when has empty cell and is not game over", func(t *testing.T) {
		values := [][]int{
			{2, 2, 4, 8},
			{2, 2, 4, 8},
			{2, 2, 4, 8},
			{2, 2, 4, 0},
		}
		assert.False(t, IsGameOver(values))
	})

	t.Run("should return not game over", func(t *testing.T) {
		testCases := []struct {
			values [][]int
		}{
			{
				values: [][]int{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
					{9, 10, 11, 12},
					{13, 14, 15, 12},
				},
			},
			{
				values: [][]int{
					{1, 2, 3, 4},
					{5, 7, 7, 8},
					{9, 10, 11, 12},
					{13, 14, 15, 16},
				},
			},
			{
				values: [][]int{
					{2, 4, 4, 16},
					{2, 8, 16, 32},
					{8, 16, 32, 128},
					{16, 32, 64, 128},
				},
			},
		}

		for _, tc := range testCases {
			assert.False(t, IsGameOver(tc.values))
		}
	})

	t.Run("should return true when is game over", func(t *testing.T) {
		values := [][]int{
			{2, 4, 8, 16},
			{4, 8, 16, 32},
			{8, 16, 32, 64},
			{16, 32, 64, 128},
		}

		assert.True(t, IsGameOver(values))
	})
}

func TestIsEqualTop(t *testing.T) {
	t.Run("should return true when values are equal", func(t *testing.T) {
		values := [][]int{
			{2, 2, 4, 8},
			{2, 2, 4, 8},
			{2, 2, 4, 8},
			{2, 2, 4, 8},
		}
		assert.True(t, isEqualTop(values, 1, 0))
		assert.True(t, isEqualTop(values, 2, 0))
		assert.True(t, isEqualTop(values, 3, 0))
		assert.True(t, isEqualTop(values, 2, 1))
	})

	t.Run("should return false when values are not equal", func(t *testing.T) {
		values := [][]int{
			{2, 2, 4, 8},
			{2, 3, 4, 8},
			{2, 2, 4, 80},
			{2, 2, 4, 8},
		}
		assert.False(t, isEqualTop(values, 1, 1))
		assert.False(t, isEqualTop(values, 2, 1))
		assert.False(t, isEqualTop(values, 2, 3))
		assert.False(t, isEqualTop(values, 3, 3))
	})
}

func TestIsEqualRight(t *testing.T) {
	t.Run("should return true when values are equal", func(t *testing.T) {
		values := [][]int{
			{2, 2, 4, 8},
			{2, 2, 8, 8},
			{2, 4, 4, 8},
			{2, 8, 4, 8},
		}
		assert.True(t, isEqualRight(values, 0, 0))
		assert.True(t, isEqualRight(values, 1, 0))
		assert.True(t, isEqualRight(values, 2, 1))
	})

	t.Run("should return false when values are not equal", func(t *testing.T) {
		values := [][]int{
			{2, 2, 4, 8},
			{2, 2, 8, 8},
			{2, 4, 4, 8},
			{2, 8, 4, 8},
		}
		assert.False(t, isEqualRight(values, 1, 1))
		assert.False(t, isEqualRight(values, 3, 2))
		assert.False(t, isEqualRight(values, 2, 2))
	})
}

func TestHasEmptyCell(t *testing.T) {
	t.Run("should return true when has empty cell", func(t *testing.T) {
		values := [][]int{
			{2, 2, 4, 8},
			{2, 2, 4, 8},
			{2, 2, 4, 8},
			{2, 2, 4, 0},
		}
		assert.True(t, HasEmptyCell(values))
	})

	t.Run("should return false when has no empty cell", func(t *testing.T) {
		values := [][]int{
			{2, 2, 4, 8},
			{2, 2, 4, 8},
			{2, 2, 4, 8},
			{2, 2, 4, 8},
		}
		assert.False(t, HasEmptyCell(values))
	})
}
