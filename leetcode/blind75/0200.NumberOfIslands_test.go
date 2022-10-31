package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberOfIslands(t *testing.T) {
	tests := []struct {
		in       [][]byte
		expected int
	}{
		{[][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'}}, 1},
		{[][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'}}, 3},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, numIslands(test.in))
	}
}
