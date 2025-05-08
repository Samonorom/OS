// main.go
package main

import (
	"flag"
	"fmt"
	"hw7_raid_sim/raid"
	"hw7_raid_sim/utils"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func CleanupDiskFiles() {
	files, err := os.ReadDir("data")
	if err != nil {
		return // silently ignore if directory doesn't exist
	}
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".dat") {
			os.Remove(filepath.Join("data", f.Name()))
		}
	}
}

func main() {
	CleanupDiskFiles()

	level := flag.String("level", "raid0", "RAID level: raid0, raid1, raid4, raid5")
	disks := flag.Int("disks", 4, "Number of disks")
	sizeMB := flag.Int("size", 100, "Benchmark size in MB")
	flag.Parse()

	var r raid.RAID
	switch *level {
	case "raid0":
		r = raid.NewRAID0(*disks)
	case "raid1":
		r = raid.NewRAID1(*disks)
	case "raid4":
		r = raid.NewRAID4(*disks)
	case "raid5":
		r = raid.NewRAID5(*disks)
	default:
		panic("Unsupported RAID level")
	}

	//blockSize := 512
	blockSize := 4096
	numBlocks := (*sizeMB * 1024 * 1024) / blockSize

	// Generate data
	data := make([][]byte, numBlocks)
	for i := range data {
		b := make([]byte, blockSize)
		rand.Read(b)
		data[i] = b
	}

	// Write benchmark
	start := time.Now()
	for i := 0; i < numBlocks; i++ {
		r.Write(i, data[i])
	}
	writeTime := time.Since(start)

	// Read benchmark
	start = time.Now()
	for i := 0; i < numBlocks; i++ {
		r.Read(i)
	}
	readTime := time.Since(start)

	fmt.Printf("RAID Level: %s\n", *level)
	fmt.Printf("Disks: %d | Size: %d MB | Blocks: %d\n", *disks, *sizeMB, numBlocks)
	fmt.Printf("Total Write Time: %s\n", utils.FormatDuration(writeTime))
	fmt.Printf("Total Read Time:  %s\n", utils.FormatDuration(readTime))
}
