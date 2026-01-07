# go-benchmarks

[![GitHub Actions Workflow Status](https://github.com/QChemX/go-benchmarks/actions/workflows/release.yml/badge.svg)](https://github.com/QChemX/go-benchmarks/blob/main/.github/workflows/release.yml)
[![GitHub last commit](https://img.shields.io/github/last-commit/QChemX/go-benchmarks)](https://github.com/QChemX/go-benchmarks/commits/main/)
[![GitHub License](https://img.shields.io/github/license/QChemX/go-benchmarks)](https://github.com/QChemX/go-benchmarks/blob/main/LICENSE)

Test the performance of CPU using a CLI tool built in Go.

## Usage

Check the usage of `benchmarks.exe`:

```bash
> ./benchmarks.exe --help
Usage of ./benchmarks.exe:
  -t duration
        Duration for each test phase (default 3s)
```

## Examples

Run `benchmarks.exe` on Windows:

```bash
> ./benchmarks.exe
=================================================================
 Go CPU Float64 Benchmark Tool (v0.2.0)
-----------------------------------------------------------------
 Arch         : amd64 / windows
 Logical CPUs : 20
 Precision    : Float64 (Double Precision)
 Mode         : FMA Simulation (ILP Optimized)
=================================================================

[1/2] Running Single-Core Benchmark... Done.
[2/2] Running Multi-Core Benchmark (20 Threads)... Done.

=================================================================
 BENCHMARK RESULTS
=================================================================
 Single-Core Performance :     193.19 GFLOPS
 Multi-Core Performance  :    1266.15 GFLOPS
-----------------------------------------------------------------
 Multi-Core Scaling      :       6.55x
=================================================================
 Note: Results represent pure FP64 throughput in Go environment.
 Values are approx. 60-80% of theoretical hardware limits due
 to compiler auto-vectorization constraints.
=================================================================
```
