package main

import "testing"

func TestFibonacci(t *testing.T) {
	ans := fibonacci(6)
	if ans != 8 {
		t.Errorf("Fibonacci = %d; debe ser 8", ans)
	}
}
