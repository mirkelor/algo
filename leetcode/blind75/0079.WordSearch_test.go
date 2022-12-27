// true Input: board = {["A","B","C","E"},["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// true Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
// false Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordSearch(t *testing.T) {
	tests := []struct {
		in1      [][]byte
		in2      string
		expected bool
	}{
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED", true},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE", true},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB", false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, exist(tt.in1, tt.in2))
	}
}
