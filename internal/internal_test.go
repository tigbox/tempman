package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	result := Hello()
	expected := "world"
	assert.Equal(t, expected, result)
}
