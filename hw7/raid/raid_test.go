// raid/raid_test.go
package raid

import (
	"testing"
)

func TestRAID0(t *testing.T) {
	r := NewRAID0(2)
	data := []byte("RAID0 block")
	r.Write(0, data)
	read, _ := r.Read(0)
	if string(read) != string(data) {
		t.Errorf("Expected %s, got %s", data, read)
	}
}

/*
func TestWriteAndRead(t *testing.T) {
	r := NewRaid(4)

	input := []byte("test block")
	r.Write(0, input)

	output := r.Read(0)
	if string(output) != string(input) {
		t.Errorf("Expected '%s', got '%s'", input, output)
	}
}
*/
