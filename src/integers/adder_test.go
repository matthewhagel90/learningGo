package integers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	assert.Equal(t, sum, expected)
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
