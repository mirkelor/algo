package blind75

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianFromDataStream(t *testing.T) {
	mf := MedianFinderConstructor()
	mf.AddNum(1)
	mf.AddNum(2)
	assert.Equal(t, 1.5, mf.FindMedian())

	mf.AddNum(3)
	assert.Equal(t, 2.0, mf.FindMedian())
}
