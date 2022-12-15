package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeAndDecodeStrings(t *testing.T) {
	tests := []struct {
		in       []string
		expected string
	}{
		{[]string{"hello", "world"}, "5#hello5#world"},
		{[]string{"hi", "hasan"}, "2#hi5#hasan"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, encode(tt.in))
		assert.Equal(t, tt.in, decode(tt.expected))
	}
}
