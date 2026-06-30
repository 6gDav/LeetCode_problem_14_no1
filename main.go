package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	inputs := []struct {
		nameofbenchmark string
		strings         []string
		target          string
	}{
		{
			"[Small Input Test]",
			[]string{"flower", "flow", "flight"},
			"fl",
		},
		{
			"[Stress Test - 10 000]",
			func() []string {
				res := make([]string, 0, 10_000)
				for i := 0; i < 10_000; i++ {
					res = append(res, "fl")
				}
				return res
			}(),
			"fl",
		},
	}

	fmt.Println("=== Benchmarking ===")

	for _, test := range inputs {
		var m1, m2 runtime.MemStats
		runtime.GC()
		runtime.ReadMemStats(&m1)
		start := time.Now()

		_ = longestCommonPrefix(test.strings)

		elapsed := time.Since(start)
		runtime.ReadMemStats(&m2)

		fmt.Printf("\n%s\n", test.nameofbenchmark)

		fmt.Printf("Execution Time: %v\n", elapsed)
		fmt.Printf("Memory Delta: %v bytes\n", m2.Alloc-m1.Alloc)
		fmt.Printf("Current Memory: %v bytes\n", m2.Sys)
	}
}
