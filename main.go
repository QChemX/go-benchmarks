package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("CPU benchmarks with GFLOPS")
	fmt.Println("run single-core and multi-core tests...")

	single := runSingleCore()
	fmt.Printf("single-core GFLOPS: %.2f\n", single)

	multi := runMultiCore()
	fmt.Printf("multi-core GFLOPS: %.2f\n", multi)

	os.Exit(0)
}
