// raid/raid4.go
package raid

import (
	"fmt"
	"hw7_raid_sim/disk"
)

type RAID4 struct {
	numDisks  int
	parityIdx int
}

func NewRAID4(numDisks int) *RAID4 {
	if numDisks < 3 {
		panic("RAID4 requires at least 3 disks")
	}
	return &RAID4{numDisks: numDisks, parityIdx: numDisks - 1}
}

func (r *RAID4) Write(blockNum int, data []byte) error {
	dataDisk := blockNum % (r.numDisks - 1)
	stripe := blockNum / (r.numDisks - 1)

	fmt.Printf("RAID4 writing block %d: data to disk %d, stripe %d, updating parity on disk %d\n",
		blockNum, dataDisk, stripe, r.parityIdx)

	// Write data
	err := disk.WriteBlock(dataDisk, stripe, data)
	if err != nil {
		return fmt.Errorf("write error on data disk %d: %v", dataDisk, err)
	}

	// Recalculate parity
	parity := make([]byte, len(data))
	for i := 0; i < r.numDisks-1; i++ {
		if i == dataDisk {
			parity = xor(parity, data)
		} else {
			b, err := disk.ReadBlock(i, stripe)
			if err == nil {
				parity = xor(parity, b)
			}
		}
	}

	return disk.WriteBlock(r.parityIdx, stripe, parity)
}

func (r *RAID4) Read(blockNum int) ([]byte, error) {
	dataDisk := blockNum % (r.numDisks - 1)
	stripe := blockNum / (r.numDisks - 1)
	return disk.ReadBlock(dataDisk, stripe)
}

func xor(a, b []byte) []byte {
	length := len(a)
	if len(b) < length {
		length = len(b)
	}
	x := make([]byte, length)
	for i := 0; i < length; i++ {
		x[i] = a[i] ^ b[i]
	}
	return x
}
