package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	expected := 28
	actual := Years(4)

	if actual != expected {
		t.Error("Expected", expected, "Actual", actual)
	}
}

func TestYearsTwo(t *testing.T) {
	expected := 28
	actual := YearsTwo(4)

	if actual != expected {
		t.Error("Expected", expected, "Actual", actual)
	}
}

func ExampleYears() {
	fmt.Println(Years(3))
	// Output: 21
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(3))
	// Output: 21
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(i)
	}
}

func BenchmarkTwoYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(i)
	}
}