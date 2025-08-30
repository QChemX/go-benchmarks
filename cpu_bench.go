package main

import (
	"runtime"
	"sync"
	"time"
)

type Mat struct {
	data []float64
	n    int
}

func newMatrix(n int) *Mat {
	return &Mat{
		data: make([]float64, n*n),
		n:    n,
	}
}

func (m *Mat) at(i, j int) float64 {
	return m.data[i*m.n+j]
}

func (m *Mat) set(i, j int, v float64) {
	m.data[i*m.n+j] = v
}

// block matrix multiplication: C = A * B
func matMulBlock(A, B, C *Mat, block int, parallel bool) {
	n := A.n

	// zero result
	for i := range C.data {
		C.data[i] = 0
	}

	worker := func(iStart, iEnd int) {
		for ii := iStart; ii < iEnd; ii += block {
			iMax := min(ii+block, n)
			for jj := 0; jj < n; jj += block {
				jMax := min(jj+block, n)
				for kk := 0; kk < n; kk += block {
					kMax := min(kk+block, n)
					// block multiply
					for i := ii; i < iMax; i++ {
						for k := kk; k < kMax; k++ {
							aik := A.at(i, k)
							for j := jj; j < jMax; j++ {
								C.data[i*n+j] += aik * B.at(k, j)
							}
						}
					}
				}
			}
		}
	}

	if parallel {
		numCPU := runtime.NumCPU()
		var wg sync.WaitGroup
		chunk := (n + numCPU - 1) / numCPU
		for c := 0; c < numCPU; c++ {
			start := c * chunk
			end := min(start+chunk, n)
			if start >= n {
				break
			}
			wg.Add(1)
			go func(si, ei int) {
				defer wg.Done()
				worker(si, ei)
			}(start, end)
		}
		wg.Wait()
	} else {
		worker(0, n)
	}
}

func heavyComputationBlock(A, B, C *Mat, block int, parallel bool) float64 {
	matMulBlock(A, B, C, block, parallel)
	return C.at(A.n-1, B.n-1)
}

func runSingleCore(n, iterations int) float64 {
	runtime.GOMAXPROCS(1)
	A := newMatrix(n)
	B := newMatrix(n)
	C := newMatrix(n)

	// init
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			A.set(i, j, float64(i+j))
			B.set(i, j, float64(i-j))
		}
	}

	block := 32
	flops := float64(2*n*n*n) * float64(iterations)

	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = heavyComputationBlock(A, B, C, block, false)
	}
	elapsed := time.Since(start).Seconds()

	return flops / elapsed / 1e9
}

func runMultiCore(n, iterations int) float64 {
	runtime.GOMAXPROCS(runtime.NumCPU())
	A := newMatrix(n)
	B := newMatrix(n)
	C := newMatrix(n)

	// init
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			A.set(i, j, float64(i+j))
			B.set(i, j, float64(i-j))
		}
	}

	block := 32
	flops := float64(2*n*n*n) * float64(iterations)

	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = heavyComputationBlock(A, B, C, block, true)
	}
	elapsed := time.Since(start).Seconds()

	return flops / elapsed / 1e9
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
