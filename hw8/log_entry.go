// log_entry.go
package main

import (
	"fmt"
	"time"
)

type LogEntry struct {
	Timestamp time.Time
	Level     string
	Context   string
	Message   string
}

func (e LogEntry) Format() string {
	return fmt.Sprintf("[%s] [%s] [%s] %s\n", e.Timestamp.Format("2006-01-02 15:04:05"), e.Level, e.Context, e.Message)
}
