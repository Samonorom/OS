// raid/raid5.go
package raid

import (
	"fmt"
	"hw7_raid_sim/disk"
)

type RAID5 struct {
	numDisks int
}

func NewRAID5(numDisks int) *RAID5 {
	if numDisks < 3 {
		panic("RAID5 requires at least 3 disks")
	}
	return &RAID5{numDisks: numDisks}
}

func (r *RAID5) Write(blockNum int, data []byte) error {
	stripe := blockNum / (r.numDisks - 1)
	dataDisk := blockNum % (r.numDisks - 1)
	parityDisk := (r.numDisks - 1 - stripe%r.numDisks)

	actualDataDisk := dataDisk
	if dataDisk >= parityDisk {
		actualDataDisk++
	}

	fmt.Printf("RAID5 writing block %d: data to disk %d, stripe %d, parity to disk %d\n",
		blockNum, actualDataDisk, stripe, parityDisk)

	err := disk.WriteBlock(actualDataDisk, stripe, data)
	if err != nil {
		return fmt.Errorf("write error on data disk %d: %v", actualDataDisk, err)
	}

	parity := make([]byte, len(data))
	for i := 0; i < r.numDisks; i++ {
		if i == parityDisk || i == actualDataDisk {
			continue
		}
		b, err := disk.ReadBlock(i, stripe)
		if err == nil {
			parity = xor(parity, b)
		}
	}
	parity = xor(parity, data)

	return disk.WriteBlock(parityDisk, stripe, parity)
}

func (r *RAID5) Read(blockNum int) ([]byte, error) {
	stripe := blockNum / (r.numDisks - 1)
	dataDisk := blockNum % (r.numDisks - 1)
	parityDisk := (r.numDisks - 1 - stripe%r.numDisks)

	actualDataDisk := dataDisk
	if dataDisk >= parityDisk {
		actualDataDisk++
	}

	return disk.ReadBlock(actualDataDisk, stripe)
}
