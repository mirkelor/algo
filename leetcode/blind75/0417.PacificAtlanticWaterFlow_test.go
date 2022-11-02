package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPacificAtlanticWaterFlow(t *testing.T) {
	tests := []struct {
		in       [][]int
		expected [][]int
	}{
		{[][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}},
			[][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}}},
		{[][]int{{1}}, [][]int{{0, 0}}},
		{[][]int{{1, 1}, {1, 1}, {1, 1}}, [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}, {2, 0}, {2, 1}}},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, pacificAtlantic(test.in))
	}
}
