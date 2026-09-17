package skus

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func init() {
	InitWithDataFileName("../../../../init/skus.json")
}

func TestList(t *testing.T) {
	t.Parallel()

	type args struct {
		skip  int64
		limit int32
	}

	var tests = []struct {
		name     string
		expected []int64
		given    args
	}{
		{
			name:     "0_0",
			expected: []int64{},
			given:    args{0, 0},
		},
		{
			name:     "0_5",
			expected: []int64{1076963, 1148162, 1625903, 2618151, 2956315},
			given:    args{0, 5},
		},
		{
			name:     "1148162_5",
			expected: []int64{1625903, 2618151, 2956315, 2958025, 3596599},
			given:    args{1148162, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := List(tt.given.skip, tt.given.limit)
			require.Equalf(t, tt.expected, actual,
				fmt.Sprintf("(%s): expected %v, actual %v", tt.name, tt.expected, actual),
			)
		})
	}
}
