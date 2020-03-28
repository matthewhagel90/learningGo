package iteration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	assert.Equal(t, repeated, expected)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
