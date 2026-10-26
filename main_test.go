package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	got := generateRandomElements(10)
	assert.Len(t, got, 10)
	for _, value := range got {
		assert.Greater(t, value, 0, "число должно быть положительным")
	}
	got = generateRandomElements(0)
	assert.Empty(t, got)
	got = generateRandomElements(-1)
	assert.Empty(t, got)
}

func TestMaximum(t *testing.T) {
	data := []int{}
	got := maximum(data)
	assert.Equal(t, 0, got)

	data = []int{4}
	got = maximum(data)
	assert.Equal(t, 4, got)

	data = []int{2, 12, 9}
	got = maximum(data)
	assert.Equal(t, 12, got)

	data = []int{500, 1, 9}
	got = maximum(data)
	assert.Equal(t, 500, got)

	data = []int{2, 1, 6}
	got = maximum(data)
	assert.Equal(t, 6, got)

	data = []int{15, 8, 15}
	got = maximum(data)
	assert.Equal(t, 15, got)
}

func TestMaxChunks(t *testing.T) {
	data := []int{}
	got := maxChunks(data)
	assert.Equal(t, 0, got)

	data = []int{100, 2, 3}
	got = maxChunks(data)
	assert.Equal(t, 100, got)

	data = []int{1, 2, 3, 7, 7, 4, 5, 6}
	got = maxChunks(data)
	assert.Equal(t, 7, got)

	data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	got = maxChunks(data)
	assert.Equal(t, 9, got)
}

func TestMaxChunksAndMaximumEqual(t *testing.T) {
	data := generateRandomElements(100)
	expected := maximum(data)
	got := maxChunks(data)

	assert.Equal(t, expected, got)
}
