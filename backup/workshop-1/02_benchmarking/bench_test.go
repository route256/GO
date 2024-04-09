package main

import (
	"testing"
)

func Benchmark_v1(b *testing.B) {
	incorrectLine := "sdfsdfsdf"
	incorrectAmount := "Name:John, Amount:sdf"
	for i := 0; i < b.N; i++ {
		parseLineV1(incorrectLine)
		parseLineV1(incorrectAmount)
	}
}

func Benchmark_v2(b *testing.B) {
	incorrectLine := "sdfsdfsdf"
	incorrectAmount := "Name:John, Amount:sdf"
	for i := 0; i < b.N; i++ {
		parseLineV2(incorrectLine)
		parseLineV2(incorrectAmount)
	}
}
