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
	offset := blockNum / r.numDisks
	fmt.Printf("RAID0 writing block %d to disk %d\n", blockNum, diskNum)

	padded := make([]byte, disk.BlockSize)
	copy(padded, data)
	return disk.WriteBlock(diskNum, offset, padded)
}

func (r *RAID0) Read(blockNum int) ([]byte, error) {
	diskNum := blockNum % r.numDisks
	offset := blockNum / r.numDisks
	return disk.ReadBlock(diskNum, offset)
}
