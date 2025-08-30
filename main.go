package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	n := flag.Int("n", 600, "matrix size (N x N)")
	iters := flag.Int("iters", 3, "number of iterations")
	flag.Parse()

	fmt.Println("CPU benchmarks with GFLOPS")
	fmt.Printf("matrix size = %d, iterations = %d\n", *n, *iters)

	single := runSingleCore(*n, *iters)
	fmt.Printf("single-core GFLOPS: %.2f\n", single)

	multi := runMultiCore(*n, *iters)
	fmt.Printf("multi-core GFLOPS: %.2f\n", multi)

	os.Exit(0)
}
