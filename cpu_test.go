package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// matrix multiplication simulates high load
func heavyComputation(n int) float64 {
	matA := make([][]float64, n)
	matB := make([][]float64, n)
	for i := 0; i < n; i++ {
		matA[i] = make([]float64, n)
		matB[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			matA[i][j] = float64(i + j)
			matB[i][j] = float64(i - j)
		}
	}

	result := make([][]float64, n)
	for i := range result {
		result[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := 0.0
			for k := 0; k < n; k++ {
				sum += matA[i][k] * matB[k][j]
			}
			result[i][j] = sum
		}
	}

	return result[n-1][n-1]
}

// test single-core performance and output GFLOPS
func BenchmarkSingleCoreGFLOPS(b *testing.B) {
	runtime.GOMAXPROCS(1)
	n := 200
	flops := float64(2 * n * n * n) // the number of floating-point operations

	b.ResetTimer()
	start := time.Now()
	for i := 0; i < b.N; i++ {
		_ = heavyComputation(n)
	}
	elapsed := time.Since(start).Seconds()
	gflops := flops * float64(b.N) / elapsed / 1e9
	fmt.Printf("single-core GFLOPS: %.2f\n", gflops)
}

// test multi-core performance and output GFLOPS
func BenchmarkMultiCoreGFLOPS(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	n := 200
	flops := float64(2 * n * n * n)

	b.ResetTimer()
	start := time.Now()
	for i := 0; i < b.N; i++ {
		_ = heavyComputation(n)
	}
	elapsed := time.Since(start).Seconds()
	gflops := flops * float64(b.N) / elapsed / 1e9
	fmt.Printf("multi-core GFLOPS: %.2f\n", gflops)
}
