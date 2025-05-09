// channel_logger.go
package main

import (
	"os"
)

type ChannelLogger struct {
	logChan chan LogEntry
	file    *os.File
	done    chan struct{}
}

func NewChannelLogger(filename string) (*ChannelLogger, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	logger := &ChannelLogger{
		logChan: make(chan LogEntry, 100),
		file:    f,
		done:    make(chan struct{}),
	}

	go logger.listen()
	return logger, nil
}

func (l *ChannelLogger) listen() {
	buffer := 0
	for entry := range l.logChan {
		l.file.WriteString(entry.Format())
		buffer++
		if buffer >= 10 {
			l.file.Sync()
			buffer = 0
		}
	}
	// Final flush when done
	l.file.Sync()
	l.done <- struct{}{}
}

func (l *ChannelLogger) Log(entry LogEntry) error {
	l.logChan <- entry
	return nil
}

func (l *ChannelLogger) Close() {
	close(l.logChan)
	<-l.done
	l.file.Close()
}
