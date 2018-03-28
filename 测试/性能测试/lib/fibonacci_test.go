package lib

import "testing"

func BenchmarkFibonacci5(b *testing.B) {
	for i:=0;i < b.N; i++ {
		Fibonacci(5)
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i:=0;i < b.N; i++ {
		Fibonacci(10)
	}
}


func BenchmarkFibonacci20(b *testing.B) {
	for i:=0;i < b.N; i++ {
		Fibonacci(20)
	}
}

