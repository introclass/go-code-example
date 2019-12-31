package testcase

import (
	"testing"
)

func TestSumSlice(t *testing.T) {
	arrays := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	expect := 45
	sum := SumSlice(arrays[:])
	if sum != expect {
		t.Errorf("result is %d (should be %d)", sum, expect)
	}
}

func BenchmarkSumSlice(b *testing.B) {
	arrays := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := arrays[:]
	for i := 0; i < b.N; i++ {
		SumSlice(s)
	}
}

func BenchmarkSumArray(b *testing.B) {
	arrays := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		SumArray(arrays)
	}
}

func BenchmarkSumArrayP(b *testing.B) {
	arrays := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arraysp := &arrays
	for i := 0; i < b.N; i++ {
		SumArrayP(arraysp)
	}
}
