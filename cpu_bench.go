package main

import (
	"time"
)

// RunCoreBenchmark runs a floating-point intensive workload for a specific duration.
// It returns the total number of floating-point operations (FLOPs) performed.
func RunCoreBenchmark(duration time.Duration) int64 {
	// -------------------------------------------------------------------
	// Core Parameters
	// -------------------------------------------------------------------
	// The batch size determines how many unrolled iterations run before checking time.
	// We use 16 variables, each performing 1 multiplication and 1 addition (FMA).
	// Total ops per inner loop: 16 vars * 2 ops (mul+add) = 32 ops.
	// Total ops per batch: 32 ops * 1000 iterations = 32,000 FLOPs.
	const batchSize = 1000
	const flopsPerBatch = 16 * 2 * batchSize

	// Initialize 16 independent registers (Float64).
	// Using distinct initial values helps prevent compiler optimization from merging variables.
	var a0, a1, a2, a3, a4, a5, a6, a7 float64 = 1.0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7
	var b0, b1, b2, b3, b4, b5, b6, b7 float64 = 2.0, 2.1, 2.2, 2.3, 2.4, 2.5, 2.6, 2.7

	// Constants for the FMA operation
	const mul float64 = 1.000000001
	const add float64 = 0.000000001

	var totalFlops int64 = 0
	start := time.Now()

	// -------------------------------------------------------------------
	// Main Computational Loop
	// -------------------------------------------------------------------
	for {
		// Check the duration periodically.
		// To minimize the overhead of time.Since(), we only check every N batches.
		if totalFlops%10000 == 0 {
			if time.Since(start) >= duration {
				break
			}
		}

		// Aggressive Loop Unrolling.
		// The goal is to fill the CPU's execution ports with independent instructions,
		// allowing Superscalar execution (doing multiple math ops per clock cycle).
		for i := 0; i < batchSize; i++ {
			// Simulating Fused Multiply-Add (FMA): result = a * mul + add
			a0 = a0*mul + add
			b0 = b0*mul + add
			a1 = a1*mul + add
			b1 = b1*mul + add
			a2 = a2*mul + add
			b2 = b2*mul + add
			a3 = a3*mul + add
			b3 = b3*mul + add
			a4 = a4*mul + add
			b4 = b4*mul + add
			a5 = a5*mul + add
			b5 = b5*mul + add
			a6 = a6*mul + add
			b6 = b6*mul + add
			a7 = a7*mul + add
			b7 = b7*mul + add
		}

		totalFlops += flopsPerBatch
	}

	// -------------------------------------------------------------------
	// Prevent Dead Code Elimination (DCE)
	// -------------------------------------------------------------------
	// Go compilers are smart. If the result is never used, the compiler might remove the loop.
	// We use a condition that is theoretically impossible to reach but unknown to the compiler at build time.
	if a0 == 999999.999 {
		// This block will never execute, but it forces the compiler to calculate all variables.
		_ = a0 + a1 + a2 + a3 + a4 + a5 + a6 + a7 +
			b0 + b1 + b2 + b3 + b4 + b5 + b6 + b7
	}

	return totalFlops
}
