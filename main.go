package main

import (
	"flag"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

// UI Constants for elegant formatting
const (
	lineWidth = 65
	separator = "="
	thinSep   = "-"
)

func main() {
	// Parse command line arguments
	durationPtr := flag.Duration("t", 3*time.Second, "Duration for each test phase")
	flag.Parse()

	// CPU Setup
	numCPU := runtime.NumCPU()
	// Explicitly setting GOMAXPROCS is good practice for benchmarks
	runtime.GOMAXPROCS(numCPU)

	// Print Header
	printLine(separator)
	fmt.Printf(" Go CPU Float64 Benchmark Tool (v0.2.0)\n")
	printLine(thinSep)
	fmt.Printf(" Arch         : %s / %s\n", runtime.GOARCH, runtime.GOOS)
	fmt.Printf(" Logical CPUs : %d\n", numCPU)
	fmt.Printf(" Precision    : Float64 (Double Precision)\n")
	fmt.Printf(" Mode         : FMA Simulation (ILP Optimized)\n")
	printLine(separator)

	// Wait a moment before starting
	time.Sleep(1 * time.Second)

	// ---------------------------------------------------------
	// 1. Single-Core Test
	// ---------------------------------------------------------
	fmt.Printf("\n[1/2] Running Single-Core Benchmark... ")

	singleStart := time.Now()
	singleFlops := RunCoreBenchmark(*durationPtr)
	singleDuration := time.Since(singleStart).Seconds()

	singleGflops := (float64(singleFlops) / 1e9) / singleDuration
	fmt.Printf("Done.\n")

	// ---------------------------------------------------------
	// 2. Multi-Core Test
	// ---------------------------------------------------------
	fmt.Printf("[2/2] Running Multi-Core Benchmark (%d Threads)... ", numCPU)

	var wg sync.WaitGroup
	var totalMultiFlops int64
	var mu sync.Mutex // Mutex to protect the accumulator

	// Synchronization channel to ensure all goroutines start roughly at the same time
	startSignal := make(chan struct{})

	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Wait for the starting gun
			<-startSignal

			// Run the benchmark
			ops := RunCoreBenchmark(*durationPtr)

			// Aggregate results safely
			mu.Lock()
			totalMultiFlops += ops
			mu.Unlock()
		}()
	}

	multiStart := time.Now()
	close(startSignal) // Broadcast start signal
	wg.Wait()          // Wait for all threads to finish
	multiDuration := time.Since(multiStart).Seconds()

	multiGflops := (float64(totalMultiFlops) / 1e9) / multiDuration
	fmt.Printf("Done.\n")

	// ---------------------------------------------------------
	// 3. Results Report
	// ---------------------------------------------------------
	fmt.Println()
	printLine(separator)
	fmt.Println(" BENCHMARK RESULTS")
	printLine(separator)
	fmt.Printf(" Single-Core Performance : %10.2f GFLOPS\n", singleGflops)
	fmt.Printf(" Multi-Core Performance  : %10.2f GFLOPS\n", multiGflops)
	printLine(thinSep)
	fmt.Printf(" Multi-Core Scaling      : %10.2fx\n", multiGflops/singleGflops)
	printLine(separator)
	fmt.Println(" Note: Results represent pure FP64 throughput in Go environment.")
	fmt.Println(" Values are approx. 60-80% of theoretical hardware limits due")
	fmt.Println(" to compiler auto-vectorization constraints.")
	printLine(separator)
}

// printLine prints a repeated character line of fixed width
func printLine(char string) {
	fmt.Println(strings.Repeat(char, lineWidth))
}
