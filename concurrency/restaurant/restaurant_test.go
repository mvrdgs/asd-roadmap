package restaurant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestaurant(t *testing.T) {
	rest := NewRestaurant(2, true)
	rest.Open()

	for i := 0; i < 5; i++ {
		rest.AddNewOrder()
	}

	counter := rest.Close()
	assert.Equal(t, uint(5), counter)
}

func BenchmarkRestaurant(b *testing.B) {
	rest := NewRestaurant(2, false)
	rest.Open()

	for i := 0; i < b.N; i++ {
		rest.AddNewOrder()
	}

	rest.Close()
}
