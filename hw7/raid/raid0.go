// raid/raid0.go
package raid

import (
	"fmt"
	"hw7_raid_sim/disk"
)

type RAID0 struct {
	numDisks int
}

func NewRAID0(numDisks int) *RAID0 {
	return &RAID0{numDisks: numDisks}
}

func (r *RAID0) Write(blockNum int, data []byte) error {
	diskNum := blockNum % r.numDisks
	fmt.Printf("RAID0 writing to disk %d: %s\n", diskNum, data)
	return disk.WriteBlock(diskNum, blockNum/r.numDisks, data)
}

func (r *RAID0) Read(blockNum int) ([]byte, error) {
	diskNum := blockNum % r.numDisks
	return disk.ReadBlock(diskNum, blockNum/r.numDisks)
}
