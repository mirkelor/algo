package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCourseSchedule(t *testing.T) {
	tests := []struct {
		in1      int
		in2      [][]int
		expected bool
	}{
		{2, [][]int{{1, 0}}, true},
		{2, [][]int{{1, 0}, {0, 1}}, false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, canFinish(tt.in1, tt.in2))
	}
}
