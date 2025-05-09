// naive_logger.go
package main

import (
	"os"
)

type NaiveLogger struct {
	file *os.File
}

func NewNaiveLogger(filename string) (*NaiveLogger, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return &NaiveLogger{file: f}, nil
}

func (l *NaiveLogger) Log(entry LogEntry) error {
	_, err := l.file.WriteString(entry.Format())
	if err != nil {
		return err
	}
	return l.file.Sync() // fsync after every write
}

func (l *NaiveLogger) Close() {
	l.file.Close()
}
