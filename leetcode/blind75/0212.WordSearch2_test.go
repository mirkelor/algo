package blind75

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordSearch2(t *testing.T) {
	tests := []struct {
		in1      [][]byte
		in2      []string
		expected []string
	}{
		{[][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, []string{"oath", "pea", "eat", "rain"}, []string{"oath", "eat"}},
		{[][]byte{{'a', 'b'}, {'c', 'd'}}, []string{"abcb"}, []string{}},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, findWords(tt.in1, tt.in2))
	}
}
