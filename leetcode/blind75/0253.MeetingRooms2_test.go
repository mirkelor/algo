package blind75

import (
	"testing"

	"github.com/Mirkelor/algo/common"
	"github.com/stretchr/testify/assert"
)

func TestMeetingRooms2(t *testing.T) {
	tests := []struct {
		in       []*common.Interval
		expected int
	}{
		{[]*common.Interval{{Start: 0, End: 30}, {Start: 5, End: 10}, {Start: 15, End: 20}}, 2},
		{[]*common.Interval{{Start: 2, End: 7}}, 1},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, minMeetingRooms(tt.in))
	}
}
