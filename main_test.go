package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		size   int
		answer int
	}{
		{-2, 0},
		{0, 0},
		{1, 1},
		{11, 11},
		{3045, 3045},
		{1004004, 1004004},
	}
	for _, v := range tests {
		slice := generateRandomElements(v.size)
		require.Equal(t, v.answer, len(slice))
		if v.answer > 0 {
			require.Positive(t, slices.Min(slice))
		}
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		slice  []int
		answer int
	}{
		{nil, 0},
		{[]int{}, 0},
		{[]int{0}, 0},
		{[]int{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}, 4},
		{[]int{1, 2, 3, 4, 567895, 5, 6, 7, 8, 34, 9, 42, 10, 43567}, 567895},
	}
	for _, v := range tests {
		data := maximum(v.slice)
		require.Equal(t, v.answer, data)
	}
}
