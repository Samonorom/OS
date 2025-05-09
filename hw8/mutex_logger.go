// mutex_logger.go
package main

import (
	"os"
	"sync"
)

type MutexLogger struct {
	file  *os.File
	mutex sync.Mutex
	count int
}

func NewMutexLogger(filename string) (*MutexLogger, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return &MutexLogger{file: f}, nil
}

func (l *MutexLogger) Log(entry LogEntry) error {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	_, err := l.file.WriteString(entry.Format())
	if err != nil {
		return err
	}

	l.count++
	if l.count >= 10 {
		err = l.file.Sync()
		l.count = 0
	}
	return err
}

func (l *MutexLogger) Close() {
	l.file.Sync()
	l.file.Close()
}
