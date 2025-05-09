# Makefile for HW8 Safe Logger

FILES = benchmark.go log_entry.go naive_logger.go mutex_logger.go channel_logger.go

run:
	go run $(FILES)

build:
	go build -o hw8_logger $(FILES)

clean:
	rm -f hw8_logger naive.log mutex.log channel.log
