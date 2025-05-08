// disk/disk.go
package disk

import (
	"fmt"
	"os"
	"path/filepath"
)

const blockSize = 4096
const BlockSize = 4096 //fpr RAID0

//const 	blockSize = 512 // Block size is 512 bytes

// WriteBlock writes a block to diskN.dat at the given index.
func WriteBlock(diskID int, blockIndex int, data []byte) error {
	if len(data) > blockSize {
		return fmt.Errorf("block too large: %d > %d", len(data), blockSize)
	}

	path := filepath.Join("data", fmt.Sprintf("disk%d.dat", diskID))
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	offset := int64(blockIndex * blockSize)
	_, err = f.Seek(offset, 0)
	if err != nil {
		return err
	}

	padded := make([]byte, blockSize)
	copy(padded, data)
	_, err = f.Write(padded)
	if err != nil {
		return err
	}

	return f.Sync() // ✅ Flushes the write to disk
}

// ReadBlock reads a block from diskN.dat at the given index.
func ReadBlock(diskID int, blockIndex int) ([]byte, error) {
	path := filepath.Join("data", fmt.Sprintf("disk%d.dat", diskID))
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	offset := int64(blockIndex * blockSize)
	_, err = f.Seek(offset, 0)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, blockSize)
	_, err = f.Read(buf)
	return buf, err
}
