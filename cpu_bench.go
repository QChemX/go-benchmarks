package main

import (
	"runtime"
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

// Run single-core benchmark and return GFLOPS
func runSingleCore() float64 {
	runtime.GOMAXPROCS(1)
	n := 200
	flops := float64(2 * n * n * n) // total FLOPs for one matrix multiplication

	const iterations = 3
	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = heavyComputation(n)
	}
	elapsed := time.Since(start).Seconds()
	gflops := flops * float64(iterations) / elapsed / 1e9
	return gflops
}

// Run multi-core benchmark and return GFLOPS
func runMultiCore() float64 {
	runtime.GOMAXPROCS(runtime.NumCPU())
	n := 200
	flops := float64(2 * n * n * n)

	const iterations = 3
	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = heavyComputation(n)
	}
	elapsed := time.Since(start).Seconds()
	gflops := flops * float64(iterations) / elapsed / 1e9
	return gflops
}
