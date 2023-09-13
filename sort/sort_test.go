package sort

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCocktailSort(t *testing.T) {
	arr := []int{2, 3, 1, 5, 4, 100, 47, 42, 38, 76, 0, 9}
	CocktailSort(arr)

	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 9, 38, 42, 47, 76, 100}, arr)
}

func TestBubbleSort(t *testing.T) {
	arr := []int{2, 3, 1, 5, 4, 100, 47, 42, 38, 76, 0, 9}
	BubbleSort(arr)

	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 9, 38, 42, 47, 76, 100}, arr)
}

func CreateRandomArray(size int) []int {
	rand.Intn(size)
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(size)
	}
	return arr
}

func BenchmarkCocktailSort(b *testing.B) {
	arr := CreateRandomArray(10000)

	for i := 0; i < b.N; i++ {
		CocktailSort(arr)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	arr := CreateRandomArray(10000)

	for i := 0; i < b.N; i++ {
		BubbleSort(arr)
	}
}
