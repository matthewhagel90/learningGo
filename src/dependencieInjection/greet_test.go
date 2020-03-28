package dependencieInjection

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "matt")

	got := buffer.String()
	want := "Hello, matt"

	assert.Equal(t, want, got)
}
