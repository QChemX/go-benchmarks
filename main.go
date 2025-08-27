package main

import (
	"fmt"
	"os"
	"os/exec"
)

// the main function
func main() {
	fmt.Println("CPU benchmarks with GFLOPS")
	fmt.Println("run single-core and multi-core tests...")

	cmd := exec.Command("go", "test", "-bench=.", "-benchmem")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("performance tests failed:", err)
		os.Exit(1)
	}
}
