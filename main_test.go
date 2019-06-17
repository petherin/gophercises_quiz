package main

import (
	"testing"
)

func TestMainFunc(t *testing.T) {
	main()
}

func BenchmarkVersion1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		version1()
	}
}

func BenchmarkVersion2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		version2()
	}
}

func BenchmarkVersion3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		version3()
	}
}

func BenchmarkVersion4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		version4()
	}
}

func BenchmarkVersion5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		version5()
	}
}

