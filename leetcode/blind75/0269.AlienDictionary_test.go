package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlienDictionary(t *testing.T) {
	tests := []struct {
		in       []string
		expected string
	}{
		{[]string{"wrt", "wrf", "er", "ett", "rftt"}, "wertf"},
		{[]string{"z", "x"}, "zx"},
		{[]string{"zxy", "zx"}, ""},
		{[]string{"zy", "zx", "zy"}, ""},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, alienOrder(tt.in))
	}
}
