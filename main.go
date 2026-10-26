package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	elements := make([]int, size)
	for i := range elements {
		elements[i] = rand.Int()
		if elements[i] == 0 {
			elements[i] = 1
		}
	}
	return elements
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for i := 1; i < len(data); i++ {
		if max < data[i] {
			max = data[i]
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	chunksCount := CHUNKS
	if len(data) < CHUNKS {
		chunksCount = len(data)
	}
	remainder := len(data) % chunksCount
	lenChunks := len(data) / chunksCount
	maxValues := make([]int, chunksCount)
	var wg sync.WaitGroup
	start := 0
	for i := 0; i < chunksCount; i++ {
		chunkSize := lenChunks
		if i < remainder {
			chunkSize = lenChunks + 1
		}
		end := start + chunkSize

		wg.Add(1)
		go func(start int, end int, i int) {
			defer wg.Done()
			maxValues[i] = maximum(data[start:end])
		}(start, end, i)
		start = end
	}
	wg.Wait()
	max := maxValues[0]
	for j := 1; j < chunksCount; j++ {
		if max < maxValues[j] {
			max = maxValues[j]
		}
	}
	return max
}

func main() {

	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)
	start := time.Now()
	fmt.Println("Ищем максимальное значение в один поток")

	max := maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)

	start = time.Now()
	chunksCount := CHUNKS
	if len(data) < CHUNKS {
		chunksCount = len(data)
	}
	fmt.Printf("Ищем максимальное значение в %d потоков\n", chunksCount)

	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
