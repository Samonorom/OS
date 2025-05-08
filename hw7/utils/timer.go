// utils/timer.go
package utils

import "time"

func FormatDuration(d time.Duration) string {
	return d.Round(time.Microsecond).String()
}
