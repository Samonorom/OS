// benchmark/benchmark.go
package benchmark

import (
	"fmt"
	"hw7_raid_sim/raid"
	"hw7_raid_sim/utils"
	"math/rand"
	"time"
)

func Run(r raid.RAID, sizeMB int) {
	blockSize := 512 // match disk.go
	numBlocks := (sizeMB * 1024 * 1024) / blockSize
	fmt.Printf("\n--- Benchmarking %T with %d blocks (%d MB) ---\n", r, numBlocks, sizeMB)

	// generate block data
	src := make([][]byte, numBlocks)
	for i := range src {
		b := make([]byte, blockSize)
		rand.Read(b)
		src[i] = b
	}

	start := time.Now()
	for i := range src {
		err := r.Write(i, src[i])
		if err != nil {
			fmt.Printf("Write error at block %d: %v\n", i, err)
		}
	}
	writeTime := time.Since(start)

	start = time.Now()
	for i := range src {
		_, err := r.Read(i)
		if err != nil {
			fmt.Printf("Read error at block %d: %v\n", i, err)
		}
	}
	readTime := time.Since(start)

	fmt.Printf("Write Time: %s\nRead Time: %s\n",
		utils.FormatDuration(writeTime), utils.FormatDuration(readTime))
}
