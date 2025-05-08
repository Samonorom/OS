// raid/raid1.go
package raid

import (
	"errors"
	"fmt"
	"hw7_raid_sim/disk"
)

type RAID1 struct {
	numDisks int
}

func NewRAID1(numDisks int) *RAID1 {
	if numDisks < 2 {
		panic("RAID1 requires at least 2 disks")
	}
	return &RAID1{numDisks: numDisks}
}

func (r *RAID1) Write(blockNum int, data []byte) error {
	fmt.Printf("RAID1 writing block %d to all %d disks\n", blockNum, r.numDisks)
	for i := 0; i < r.numDisks; i++ {
		err := disk.WriteBlock(i, blockNum, data)
		if err != nil {
			return fmt.Errorf("error writing to disk %d: %v", i, err)
		}
	}
	return nil
}

func (r *RAID1) Read(blockNum int) ([]byte, error) {
	for i := 0; i < r.numDisks; i++ {
		data, err := disk.ReadBlock(i, blockNum)
		if err == nil {
			return data, nil
		}
	}
	return nil, errors.New("all RAID1 disks failed")
}
