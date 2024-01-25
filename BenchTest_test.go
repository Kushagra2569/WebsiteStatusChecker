package main

import (
	"fmt"
	"testing"
)

func BenchmarkStatusChecker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
  fmt.Println("benchmark done")
}
