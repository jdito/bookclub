package main

import (
	"strings"
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	args := strings.SplitAfter("Very long test string", "")
	for i := 0; i < b.N; i++ {
		echo1(args)
	}
}

func BenchmarkEcho2(b *testing.B) {
	args := strings.SplitAfter("Very long test string", "")
	for i := 0; i < b.N; i++ {
		echo2(args)
	}
}

func BenchmarkEcho3(b *testing.B) {
	args := strings.SplitAfter("Very long test string", "")
	for i := 0; i < b.N; i++ {
		echo3(args)
	}
}
