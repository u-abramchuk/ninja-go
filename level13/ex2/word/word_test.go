package word

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCount(t *testing.T) {
	expected := 3
	actual := Count("Who am I?")

	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", actual, expected)
	}
}

func TestUseCount(t *testing.T) {
	text := "This assumes that the string representations of equivalent maps are the same"
	expected := map[string]int{
		"This":            1,
		"assumes":         1,
		"that":            1,
		"the":             2,
		"string":          1,
		"representations": 1,
		"of":              1,
		"equivalent":      1,
		"maps":            1,
		"are":             1,
		"same":            1,
	}
	actual := UseCount(text)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v,\nActual: %v", expected, actual)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("Who am I?")
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount("Who am I?")
	}
}

func ExampleCount() {
	fmt.Println(Count("I can see"))
	// Output: 3
}

func ExampleUseCount() {
	fmt.Println(UseCount("I can see"))
	// Output: map[I:1 can:1 see:1]
}
