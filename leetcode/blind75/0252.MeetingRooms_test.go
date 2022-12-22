package blind75

import (
	"testing"

	"github.com/Mirkelor/algo/common"
	"github.com/stretchr/testify/assert"
)

func TestMeetingRooms(t *testing.T) {
	tests := []struct {
		in       []*common.Interval
		expected bool
	}{
		{[]*common.Interval{{Start: 0, End: 30}, {Start: 5, End: 10}, {Start: 15, End: 20}}, false},
		{[]*common.Interval{{Start: 5, End: 8}, {Start: 9, End: 15}}, true},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, canAttendMeetings(tt.in))
	}
}
