package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type testData struct {
		data     []int
		expected float64
	}

	tsts := []testData{
		testData{
			data:     []int{2, 5, 8, 3},
			expected: 4.0,
		},
		testData{
			data:     []int{1, 3, 8, 100, 4, 7},
			expected: 5.5,
		},
	}

	for _, v := range tsts {
		actual := CenteredAvg(v.data)

		if actual != v.expected {
			t.Errorf("Expected: %v, Actual: %v", v.expected, actual)
		}
	}
}

func BenchmarkCenteredAvg(b *testing.B) {
	xi := []int{1, 2, 3, 4, 6, 7, 9}
	for i := 0; i < b.N; i++ {
		CenteredAvg(xi)
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 3, 4, 5}))
	// Output: 3
}
