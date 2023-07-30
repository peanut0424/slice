package slice

import "testing"

func TestSlice_DelItem(t *testing.T) {
	testCases := []struct {
		name        string
		originCap   int
		enqueueLoop int
		expectCap   int
	}{
		{
			name:        "小于256",
			originCap:   128,
			enqueueLoop: 64,
			expectCap:   128,
		},
		{
			name:        "大于256, 不足1/4",
			originCap:   1000,
			enqueueLoop: 20,
			expectCap:   500,
		},
		{
			name:        "大于256, 超过1/4",
			originCap:   1000,
			enqueueLoop: 400,
			expectCap:   1000,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := make([]interface{}, 0, tc.originCap)

			for i := 0; i < tc.enqueueLoop; i++ {
				l = append(l, i)
			}
			l = shrink(l)
			t.Log(tc.expectCap, cap(l))
		})
	}

}
