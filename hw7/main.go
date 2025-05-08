// main.go
package main

import (
	"flag"
	"fmt"
	"hw7_raid_sim/raid"
	"hw7_raid_sim/utils"
	"os"
	"time"
)

func main() {
	fmt.Println("RAID Benchmarking Simulator")

	raidLevel := flag.String("level", "raid0", "RAID level to use: raid0, raid1, raid4, raid5")
	disks := flag.Int("disks", 4, "Number of disks")
	block := flag.String("block", "example block data", "Data to write")
	index := flag.Int("index", 0, "Block number")
	flag.Parse()

	var r raid.RAID
	switch *raidLevel {
	case "raid0":
		r = raid.NewRAID0(*disks)
	case "raid1":
		r = raid.NewRAID1(*disks)
	case "raid4":
		r = raid.NewRAID4(*disks)
	case "raid5":
		r = raid.NewRAID5(*disks)
	default:
		fmt.Fprintf(os.Stderr, "Unknown RAID level: %s\n", *raidLevel)
		os.Exit(1)
	}

	start := time.Now()
	r.Write(*index, []byte(*block))
	writeDuration := time.Since(start)

	start = time.Now()
	data, err := r.Read(*index)
	readDuration := time.Since(start)

	if err != nil {
		fmt.Println("Read error:", err)
	} else {
		fmt.Println("Read data:", string(data))
	}

	fmt.Printf("Write Time: %s | Read Time: %s\n", utils.FormatDuration(writeDuration), utils.FormatDuration(readDuration))
}
