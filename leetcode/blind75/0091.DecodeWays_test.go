package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeWays(t *testing.T) {
	tests := []struct {
		in       string
		expected int
	}{
		{"12", 2},
		{"226", 3},
		{"06", 0},
		{"306", 0},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, numDecodings(tt.in))
	}
}
