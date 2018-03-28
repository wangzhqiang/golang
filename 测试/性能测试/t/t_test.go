package t

import "testing"

const N = 1000000

type T struct {
	data [10]int32
}

func newT() []*T {
	s := make([]*T, N)
	for i := 0; i < N; i++ {
		s[i] = new(T)
	}
	return s
}

func BenchmarkStruct(b *testing.B) {
	s := make([]T, N)
	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		for i := range s {
			_ = s[i].data[0]
		}
	}
}

func BenchmarkPointer(b *testing.B) {
	s := newT()
	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		for i := range s {
			_ = s[i].data[0]
		}
	}
}
