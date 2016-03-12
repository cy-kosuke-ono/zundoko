package zundoko

import "testing"

func BenchmarkRun1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run()
	}
}

func BenchmarkRun2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run2()
	}
}
