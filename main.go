package main

import (
	"fmt"
	"log"
	"math/rand"
	"slices"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size < 1 {
		log.Println("Error: size less than 1")
		return []int{}
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]int, size)
	for i := 0; i < size; i++ {

		result[i] = r.Intn(size) + 1
		// +1 потому, что в задании написано "максимальное значение в слайсе целых положительных чисел"
		//без +1 Intn() генерирует в значения 0
	}

	return result
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		log.Println("Error: Array must be greater than 0")
		return 0
	}

	return slices.Max(data)
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		log.Println("Error: Array must be greater than 0")
		return 0
	}

	size := len(data) / CHUNKS
	maxes := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		start := i * size
		end := start + size

		// Обработка последнего chunk, если длина не делится нацело
		if i == CHUNKS-1 {
			end = len(data)
		}

		// Пропуск пустых chunk
		if start >= end {
			continue
		}

		chunk := data[start:end]

		wg.Add(1)
		go func(chunk []int, index int) {
			defer wg.Done()
			maxes[index] = maximum(chunk)
		}(chunk, i)
	}

	wg.Wait()

	return maximum(maxes)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)
	// ваш код здесь

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
