// raid/interface.go
package raid

type RAID interface {
	Write(blockNum int, data []byte) error
	Read(blockNum int) ([]byte, error)
}
